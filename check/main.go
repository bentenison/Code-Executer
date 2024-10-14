package main

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	pb "check/proto"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Stats struct {
	MemoryStats struct {
		Usage uint64 `json:"usage"`
		Limit uint64 `json:"limit"`
	} `json:"memory_stats"`
	CPUStats struct {
		CPUUsage struct {
			TotalUsage uint64 `json:"total_usage"`
		} `json:"cpu_usage"`
	} `json:"cpu_stats"`
}

const containerName = "dazzling_einstein" // Change this to your container name

func runPythonCode(code string) (string, error) {

	// suffix := rand.Int()
	// suff := strconv.Itoa(suffix)
	// // Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}
	stats, err := cli.ContainerStatsOneShot(context.TODO(), "55804d50954612d7851db357bd720aa8bfe35f3dfa79cd0836fe88557295c5df")
	if err != nil {
		return "", err
	}
	// json.Unmarshal(stats.Body)
	// var data []byte
	// // // data := bytes.Buffer{}
	// d, err := stats.Body.Read(data)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(d)
	decoder := json.NewDecoder(stats.Body)
	for {
		var s Stats
		if err := decoder.Decode(&s); err == io.EOF {
			break // End of the stream
		} else if err != nil {
			return "", err
		}

		// Extract memory usage
		fmt.Printf("Usage: %d mb, Limit: %d mb\n", s.MemoryStats.Usage/(1024*1024), s.MemoryStats.Limit/(1024*1024))

		// Extract CPU usage
		fmt.Printf("Total CPU Usage: %d\n", s.CPUStats.CPUUsage.TotalUsage)

		// Optionally break after one read to get stats at a single point in time
		break
	}
	// defer stats.Body.Close()
	// log.Println("the stats are", stats.Body.Read(data))
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		log.Fatalf("Errorc listing containers: %v", err)
	}
	// // var containerName string
	// Print container names and IDs
	for _, container := range containers {
		for _, name := range container.Names {
			fmt.Printf("Container ID: %s, Name: %s\n", container.ID, name)
		}
	}
	// Create a temporary file for the Python code
	// tmpFile, err := os.Create(fmt.Sprintf("%s.py", suff))
	// if err != nil {
	// 	return "", err
	// }
	// defer os.Remove(tmpFile.Name())

	// // Write the Python code to the file
	// if err := os.WriteFile(tmpFile.Name(), []byte(code), 0644); err != nil {
	// 	return "", err
	// }
	// buf, err := tarFile(tmpFile.Name())
	// if err != nil {
	// 	log.Println("error in tarfile", err)
	// }
	// // Copy the file to the existing container
	// err = cli.CopyToContainer(context.Background(), containerName, "app/", &buf, container.CopyToContainerOptions{})
	// if err != nil {
	// 	return "", err
	// }

	// // Execute the Python script in the existing container
	// execConfig := container.ExecOptions{
	// 	Cmd:          []string{"python", filepath.Join("/app", filepath.Base(tmpFile.Name()))},
	// 	AttachStdout: true,
	// 	AttachStderr: true,
	// }

	// execID, err := cli.ContainerExecCreate(context.Background(), containerName, execConfig)
	// if err != nil {
	// 	return "", err
	// }

	// // err = cli.ContainerExecStart(context.Background(), execID.ID, container.ExecStartOptions{})
	// // if err != nil {
	// // 	return "", err
	// // }
	// res, err := cli.ContainerExecAttach(context.TODO(), execID.ID, container.ExecAttachOptions{})
	// if err != nil {
	// 	return "", err
	// }
	// defer res.Close()
	// var logBuf bytes.Buffer
	// if _, err := logBuf.ReadFrom(res.Conn); err != nil {
	// 	return "", err
	// }
	// return logBuf.String(), nil
	return "", nil
}

func main() {
	// r := gin.Default()
	// r.GET("/execute", ExecutionHandler)

	// // test := 100
	// // start := time.Now()
	// // for i := 0; i < test; i++ {
	// // }
	// if err := r.Run(":8000"); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("time required for 100 executions is", time.Since(start))
	conn, err := grpc.NewClient(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cli := pb.NewExecutorServiceClient(conn)
	uploadFile(cli, "./sample.py")
	// client:= proto
}
func uploadFile(client pb.ExecutorServiceClient, filePath string) {
	stream, err := client.HandleExecution(context.Background())
	if err != nil {
		log.Fatalf("Error opening stream: %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	buf := make([]byte, 1024) // 1 KB chunks
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		err = stream.Send(&pb.ExecutionRequest{Content: buf[:n], Uid: "abc123", Qid: "pqr123"})
		if err != nil {
			log.Fatalf("Error sending file chunk: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}
	log.Println(res)
	// fmt.Printf("Upload status: %v, message: %s\n", res.Success, res.Message)
}
func ExecutionHandler(c *gin.Context) {
	pythonCode := `
def main(n):
    # User's main logic starts here
    if n <= 1:
        return False  # Numbers less than or equal to 1 are not prime
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False  # Found a factor, so n is not prime
    return True
    # User's main logic ends here

if __name__ == '__main__':
    all_passed = True
    test_cases = [
        (2,),
        (11,),
        (15,),
        (1,)
    ]
    expected_outputs = [
        True,
        True,
        False,
        False
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = main(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)
`
	output, err := runPythonCode(pythonCode)
	if err != nil {
		log.Fatalf("Errorc running python code: %v", err)
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
