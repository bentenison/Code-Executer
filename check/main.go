package main

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

const containerName = "hopeful_liskov" // Change this to your container name

func runPythonCode(code string) (string, error) {
	suffix := rand.IntN(4000)
	suff := strconv.Itoa(suffix)
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	// containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	// if err != nil {
	// 	log.Fatalf("Error listing containers: %v", err)
	// }
	// // var containerName string
	// // Print container names and IDs
	// for _, container := range containers {
	// 	for _, name := range container.Names {
	// 		fmt.Printf("Container ID: %s, Name: %s\n", container.ID, name)
	// 	}
	// }
	// Create a temporary file for the Python code
	tmpFile, err := os.Create(fmt.Sprintf("%s.py", suff))
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	// Write the Python code to the file
	if err := os.WriteFile(tmpFile.Name(), []byte(code), 0644); err != nil {
		return "", err
	}
	buf, err := tarFile(tmpFile.Name())
	if err != nil {
		log.Println("error in tarfile", err)
	}
	// Copy the file to the existing container
	err = cli.CopyToContainer(context.Background(), containerName, "app/", &buf, container.CopyToContainerOptions{})
	if err != nil {
		return "", err
	}

	// Execute the Python script in the existing container
	execConfig := container.ExecOptions{
		Cmd:          []string{"python", filepath.Join("/app", filepath.Base(tmpFile.Name()))},
		AttachStdout: true,
		AttachStderr: true,
	}

	execID, err := cli.ContainerExecCreate(context.Background(), containerName, execConfig)
	if err != nil {
		return "", err
	}

	// err = cli.ContainerExecStart(context.Background(), execID.ID, container.ExecStartOptions{})
	// if err != nil {
	// 	return "", err
	// }
	res, err := cli.ContainerExecAttach(context.TODO(), execID.ID, container.ExecAttachOptions{})
	if err != nil {
		return "", err
	}
	defer res.Close()
	var logBuf bytes.Buffer
	if _, err := logBuf.ReadFrom(res.Conn); err != nil {
		return "", err
	}

	return logBuf.String(), nil
}

func main() {
	r := gin.Default()
	r.GET("/execute", ExecutionHandler)

	// test := 100
	// start := time.Now()
	// for i := 0; i < test; i++ {
	// }
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("time required for 100 executions is", time.Since(start))

}
func ExecutionHandler(c *gin.Context) {

	pythonCode := `
def main(input):
		# User's main logic starts here
		cleaned_input = input.lower().replace(" ", "")
		return cleaned_input == cleaned_input[::-1]
		# User's main logic ends here
	
if __name__ == "__main__":
		test_input = "A man a plan a canal Panama"  # User input for the test case
		expected_output = "True"  # Expected output
		
		result = str(main(test_input))
		print("Pass" if result == expected_output else "Fail")`
	output, err := runPythonCode(pythonCode)
	if err != nil {
		log.Fatalf("Error running python code: %v", err)
	}

	fmt.Printf("%s", output)
	c.JSON(http.StatusOK, gin.H{
		"message": "execution success!",
	})
}
func tarFile(filePath string) (bytes.Buffer, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	// Get file information
	fi, err := os.Stat(filePath)
	if err != nil {
		return buf, err
	}
	// Create a tar header
	header, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		return buf, err
	}

	// Set the name for the header
	header.Name = fi.Name() // Use the file name

	// Write the header to the tar writer
	if err := tw.WriteHeader(header); err != nil {
		return buf, err
	}

	// If the file is regular, open and write its contents
	if fi.Mode().IsRegular() {
		fileReader, err := os.Open(filePath)
		if err != nil {
			return buf, err
		}
		defer fileReader.Close()

		// Copy the file contents to the tar writer
		if _, err := io.Copy(tw, fileReader); err != nil {
			return buf, err
		}
	}

	// Close the tar writer
	if err := tw.Close(); err != nil {
		return buf, err
	}

	return buf, nil
}
