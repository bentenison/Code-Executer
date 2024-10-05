package main

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const containerName = "nervous_lewin" // Change this to your container name

func runPythonCode(code string) (string, error) {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
	}

	// Print container names and IDs
	for _, container := range containers {
		for _, name := range container.Names {
			fmt.Printf("Container ID: %s, Name: %s\n", container.ID, name)
		}
	}
	// Create a temporary file for the Python code
	tmpFile, err := os.Create("temp.py")
	if err != nil {
		return "", err
	}
	// defer os.Remove(tmpFile.Name())

	// Write the Python code to the file
	if err := os.WriteFile(tmpFile.Name(), []byte(code), 0644); err != nil {
		return "", err
	}
	buf, err := tarFile(tmpFile.Name())
	if err != nil {
		log.Fatal("error in tarfile", err)
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
	pythonCode := `
print("Hello from Python in an existing container!"
`

	output, err := runPythonCode(pythonCode)
	if err != nil {
		log.Fatalf("Error running Python code: %v", err)
	}

	fmt.Printf("%s", output)
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
