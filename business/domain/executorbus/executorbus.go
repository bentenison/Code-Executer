package executorbus

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "github.com/bentenison/microservice/api/domain/executor-api/grpc/proto"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Storer interface{}
type Business struct {
	log           *logger.CustomLogger
	delegate      *delegate.Delegate
	storer        Storer
	cli           *client.Client
	containerSpec map[string]ContainerSpec
}

func NewBusiness(log *logger.CustomLogger, delegate *delegate.Delegate, storer Storer, cli *client.Client) *Business {
	return &Business{
		log:           log,
		delegate:      delegate,
		storer:        storer,
		cli:           cli,
		containerSpec: make(map[string]ContainerSpec),
	}
}
func (b *Business) ExecuteCode(ctx context.Context, path, language, uid, qid string) (*pb.ExecutionResponse, error) {
	var execResponse pb.ExecutionResponse
	// b.log.Errorc(ctx, "questionID", map[string]interface{}{
	// 	"qid": qid,
	// })
	// b.log.Errorc(ctx, "userID", map[string]interface{}{
	// 	"uid": uid,
	// })
	// get container spec
	specs, err := b.getContainerSpec(language)
	if err != nil {
		b.log.Errorc(ctx, "error in getting specs", map[string]interface{}{
			"error": err.Error(),
		})
		return &execResponse, err
	}
	b.log.Infoc(ctx, "container sppec values", map[string]interface{}{
		"containerSpec": specs,
	})
	// create a tarfile from the temp file
	buf, err := b.readTempFile(path)
	if err != nil {
		return &execResponse, err
	}
	// Copy the file to the existing container
	err = b.cli.CopyToContainer(context.Background(), specs.ID, "app/", buf, container.CopyToContainerOptions{})
	if err != nil {
		return &execResponse, err
	}

	// // Execute the Python script in the existing container
	execConfig := container.ExecOptions{
		Cmd:          []string{"python", filepath.Join("/app", filepath.Base(path))},
		AttachStdout: true,
		AttachStderr: true,
	}

	execID, err := b.cli.ContainerExecCreate(context.Background(), specs.ID, execConfig)
	if err != nil {
		return &execResponse, err
	}

	// err = cli.ContainerExecStart(context.Background(), execID.ID, container.ExecStartOptions{})
	// if err != nil {
	// 	return "", err
	// }
	startTime := time.Now()
	res, err := b.cli.ContainerExecAttach(context.TODO(), execID.ID, container.ExecAttachOptions{})
	if err != nil {
		return &execResponse, err
	}
	endTime := time.Since(startTime)
	// _ = res
	stats, err := b.cli.ContainerStatsOneShot(ctx, specs.ID)
	if err != nil {
		return &execResponse, err
	}
	defer stats.Body.Close()
	defer res.Close()
	decoder := json.NewDecoder(stats.Body)
	var s Stats
	for {
		if err := decoder.Decode(&s); err == io.EOF {
			break // End of the stream
		} else if err != nil {
			return &execResponse, err
		}
		// Optionally break after one read to get stats at a single point in time
		break
	}
	execResponse.CpuStats = fmt.Sprintf("%d", s.CPUStats.CPUUsage.TotalUsage)
	execResponse.RamUsed = fmt.Sprintf("%d", s.MemoryStats.Usage)
	execResponse.TotalRAM = fmt.Sprintf("%d", s.MemoryStats.Limit)
	execResponse.PercetRAMUsage = fmt.Sprintf("%d", (s.MemoryStats.Usage/s.MemoryStats.Limit)*100)
	execResponse.ExecTime = endTime.String()
	var logBuf bytes.Buffer
	if _, err := logBuf.ReadFrom(res.Conn); err != nil {
		return &execResponse, err
	}
	go b.performCleanup(filepath.Base(path), specs.ID)
	execResponse.Output = convertOutput(logBuf)
	//TODO: ADD the code execution stats here
	return &execResponse, nil
}
func convertOutput(logBuf bytes.Buffer) string {
	result := []byte{}
	for _, b := range logBuf.Bytes() {
		if b >= 32 && b <= 126 {
			result = append(result, b)
		}
	}
	return string(result)
}
func (b *Business) getContainerSpec(language string) (ContainerSpec, error) {
	var containerSpec ContainerSpec
	if spec, ok := b.containerSpec[language]; ok {
		return spec, nil
	}
	filters := filters.NewArgs()
	filters.Add("label", fmt.Sprintf("app=%s-executor", language))
	filters.Add("label", fmt.Sprintf("language=%s", language))
	containers, err := b.cli.ContainerList(context.Background(), container.ListOptions{
		Filters: filters,
	})
	if err != nil {
		b.log.Errorc(context.TODO(), "Errorc listing containers:", map[string]interface{}{
			"error": err.Error(),
		})
		return containerSpec, err
	}
	// var containerName string
	// Print container names and IDs
	for _, container := range containers {
		containerSpec.ID = container.ID
		containerSpec.Image = container.Image
		containerSpec.Names = container.Names
		containerSpec.ImageID = container.ImageID
		containerSpec.Command = container.Command
		containerSpec.Created = container.Created
		containerSpec.Status = container.Status
		// containerSpec.Ports = []Port(container.Ports)
	}
	b.containerSpec[language] = containerSpec
	return containerSpec, nil
}
func (b *Business) performCleanup(path, id string) error {
	cmd := []string{"rm", filepath.Join("/app", path)}
	// Create a new exec instance
	execConfig := container.ExecOptions{
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
		Cmd:          cmd,
	}

	execIDResp, err := b.cli.ContainerExecCreate(context.Background(), id, execConfig)
	if err != nil {
		return err
	}
	// Start the exec instance
	execStartCheck := container.ExecStartOptions{Detach: false, Tty: false}
	err = b.cli.ContainerExecStart(context.Background(), execIDResp.ID, execStartCheck)
	if err != nil {
		log.Fatalf("Error starting exec instance: %v", err)
	}
	return nil
}
func (b *Business) readTempFile(path string) (*bytes.Buffer, error) {
	buf, err := tarFile(path)
	if err != nil {
		log.Println("error in tarfile", err)
		return buf, err
	}
	return buf, nil
}
func tarFile(filePath string) (*bytes.Buffer, error) {
	// Create a buffer to hold the tar content (if needed, you can write directly to file or response)
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	// Open the file (opening the file here, to avoid doing it twice)
	fileReader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	// Get file information
	fi, err := fileReader.Stat()
	if err != nil {
		return nil, err
	}

	// Create a tar header based on file info
	header, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		return nil, err
	}

	// Set the header name (preserving the file name)
	header.Name = fi.Name()

	// Write the header to the tar writer
	if err := tw.WriteHeader(header); err != nil {
		return nil, err
	}

	// If it's a regular file, copy the content directly into the tar archive
	if fi.Mode().IsRegular() {
		// Stream the file contents into the tar, minimizing memory footprint
		if _, err := io.Copy(tw, fileReader); err != nil {
			return nil, err
		}
	}

	// Close the tar writer to flush all remaining data
	if err := tw.Close(); err != nil {
		return nil, err
	}
	os.Remove(filePath)
	// Return the buffer (can be written to disk or sent over network)
	return &buf, nil
}

// func tarFile(filePath string) (*bytes.Buffer, error) {
// 	var buf bytes.Buffer
// 	tw := tar.NewWriter(&buf)

// 	// Get file information
// 	fi, err := os.Stat(filePath)
// 	if err != nil {
// 		return buf, err
// 	}
// 	// Create a tar header
// 	header, err := tar.FileInfoHeader(fi, "")
// 	if err != nil {
// 		return buf, err
// 	}

// 	// Set the name for the header
// 	header.Name = fi.Name() // Use the file name

// 	// Write the header to the tar writer
// 	if err := tw.WriteHeader(header); err != nil {
// 		return buf, err
// 	}

// 	// If the file is regular, open and write its contents
// 	if fi.Mode().IsRegular() {
// 		fileReader, err := os.Open(filePath)
// 		if err != nil {
// 			return buf, err
// 		}
// 		defer fileReader.Close()

// 		// Copy the file contents to the tar writer
// 		if _, err := io.Copy(tw, fileReader); err != nil {
// 			return buf, err
// 		}
// 	}

// 	// Close the tar writer
// 	if err := tw.Close(); err != nil {
// 		return buf, err
// 	}

// 	return buf, nil
// }
