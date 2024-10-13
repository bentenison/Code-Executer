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
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type Storer interface{}
type Business struct {
	log      *logger.CustomLogger
	delegate *delegate.Delegate
	db       mux.DataSource
	storer   Storer
	cli      *client.Client
}

func NewBusiness(log *logger.CustomLogger, delegate *delegate.Delegate, db mux.DataSource, storer Storer, cli *client.Client) *Business {
	return &Business{
		log:      log,
		delegate: delegate,
		db:       db,
		storer:   storer,
		cli:      cli,
	}
}
func (b *Business) ExecuteCode(ctx context.Context, path, language string) (*pb.ExecutionResponse, error) {
	var execResponse *pb.ExecutionResponse
	// get container spec
	specs, err := b.getContainerSpec(language)
	if err != nil {
		b.log.Errorc(ctx, "error in getting specs", map[string]interface{}{
			"error": err.Error(),
		})
		return execResponse, err
	}
	b.log.Infoc(ctx, "container sppec values", map[string]interface{}{
		"containerSpec": specs,
	})
	// create a tarfile from the temp file
	buf, err := b.readTempFile(path)
	if err != nil {
		return execResponse, err
	}
	// Copy the file to the existing container
	err = b.cli.CopyToContainer(context.Background(), specs.ID, "app/", &buf, container.CopyToContainerOptions{})
	if err != nil {
		return execResponse, err
	}

	// // Execute the Python script in the existing container
	execConfig := container.ExecOptions{
		Cmd:          []string{"python", filepath.Join("/app", filepath.Base(tmpFile.Name()))},
		AttachStdout: true,
		AttachStderr: true,
	}

	execID, err := b.cli.ContainerExecCreate(context.Background(), specs.ID, execConfig)
	if err != nil {
		return execResponse, err
	}

	// err = cli.ContainerExecStart(context.Background(), execID.ID, container.ExecStartOptions{})
	// if err != nil {
	// 	return "", err
	// }
	startTime := time.Now()
	res, err := b.cli.ContainerExecAttach(context.TODO(), execID.ID, container.ExecAttachOptions{})
	if err != nil {
		return execResponse, err
	}
	endTime := time.Since(startTime)
	_ = res
	stats, err := b.cli.ContainerStatsOneShot(ctx, specs.ID)
	if err != nil {
		return execResponse, err
	}
	decoder := json.NewDecoder(stats.Body)
	var s Stats
	for {
		if err := decoder.Decode(&s); err == io.EOF {
			break // End of the stream
		} else if err != nil {
			return execResponse, err
		}

		// Extract memory usage
		fmt.Printf("Usage: %d mb, Limit: %d mb\n", s.MemoryStats.Usage/(1024*1024), s.MemoryStats.Limit/(1024*1024))

		// Extract CPU usage
		fmt.Printf("Total CPU Usage: %d\n", s.CPUStats.CPUUsage.TotalUsage)

		// Optionally break after one read to get stats at a single point in time
		break
	}
	execResponse.CpuStats = fmt.Sprintf("%.2f", s.CPUStats.CPUUsage.TotalUsage)
	execResponse.RamUsed = fmt.Sprintf("%.2f", s.MemoryStats.Usage)
	// defer res.Close()
	// var logBuf bytes.Buffer
	// if _, err := logBuf.ReadFrom(res.Conn); err != nil {
	// 	return "", err
	// }
	return execResponse, nil
}
func (b *Business) getContainerSpec(language string) (ContainerSpec, error) {
	var containerSpec ContainerSpec

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
		container.Ports = container.Ports
	}
	return containerSpec, nil
}

func (b *Business) readTempFile(path string) (bytes.Buffer, error) {
	// suff := fmt.Sprintf("%d", rand.Int())
	// tmpFile, err := os.Create(fmt.Sprintf("%s.py", suff))
	// if err != nil {
	// 	return nil, err
	// }
	// defer os.Remove(tmpFile.Name())
	// data, err := os.ReadFile(path)
	// if err != nil {
	// 	return bytes.Buffer{}, err
	// }
	// Write the Python code to the file
	// if err := os.WriteFile(tmpFile.Name(), []byte(code), 0644); err != nil {
	// 	return err
	// }
	buf, err := tarFile(path)
	if err != nil {
		log.Println("error in tarfile", err)
		return buf, err
	}
	return buf, nil
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
