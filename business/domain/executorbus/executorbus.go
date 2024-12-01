package executorbus

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	pb "github.com/bentenison/microservice/api/domain/executor-api/grpc/proto"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"golang.org/x/exp/rand"
)

const maxRetries = 5

type Storer interface {
	GetLanguages(ctx context.Context) ([]*Language, error)
	GetAllLangSpecs() ([]LanguageSpecification, error)
	GetLanguageSpecsByID(id int) (LanguageSpecification, error)
	Set(ctx context.Context, key string, val any, ttl time.Duration) (string, error)
	Get(ctx context.Context, key string, res any) error
}
type Business struct {
	log           *logger.CustomLogger
	delegate      *delegate.Delegate
	storer        Storer
	cli           *client.Client
	languages     []*Language
	containerSpec map[string]ContainerSpec
	lb            map[string]*LoadBalancer
	lbMutex       sync.RWMutex
}
type LoadBalancer struct {
	containers []string
	currentIdx int
	mu         sync.Mutex
}

func NewContainerLB(allowedLanguages []*Language, cli *client.Client, log *logger.CustomLogger) (map[string]*LoadBalancer, error) {
	lb := make(map[string]*LoadBalancer)
	for _, v := range allowedLanguages {
		filters := filters.NewArgs()
		filters.Add("label", fmt.Sprintf("app=%s-executor", strings.ToLower(v.Name)))
		filters.Add("label", fmt.Sprintf("language=%s", strings.ToLower(v.Name)))
		containers, err := cli.ContainerList(context.Background(), container.ListOptions{
			Filters: filters,
		})
		if err != nil {
			log.Errorc(context.TODO(), "Errorc listing containers:", map[string]interface{}{
				"error": err.Error(),
			})
			return lb, err
		}
		if _, ok := lb[strings.ToLower(v.Name)]; !ok {
			lb[strings.ToLower(v.Name)] = &LoadBalancer{}
		}
		// lb[strings.ToLower(v.Name)].mu = sync.Mutex{}
		for _, container := range containers {
			lb[strings.ToLower(v.Name)].containers = append(lb[strings.ToLower(v.Name)].containers, container.ID)
		}
	}
	return lb, nil
}
func UpdateLB(allowedLanguages []*Language, b *Business, t time.Time) {
	// lb := make(map[string]*LoadBalancer)
	for _, v := range allowedLanguages {
		filters := filters.NewArgs()
		filters.Add("label", fmt.Sprintf("app=%s-executor", strings.ToLower(v.Name)))
		filters.Add("label", fmt.Sprintf("language=%s", strings.ToLower(v.Name)))
		containers, err := b.cli.ContainerList(context.Background(), container.ListOptions{
			Filters: filters,
		})
		if err != nil {
			b.log.Errorc(context.TODO(), "Errorc listing containers:", map[string]interface{}{
				"error": err.Error(),
			})
			// return lb, err
		}
		b.lbMutex.Lock()
		defer b.lbMutex.Unlock()
		if _, ok := b.lb[strings.ToLower(v.Name)]; !ok {
			b.lb[strings.ToLower(v.Name)] = &LoadBalancer{}
		}
		// lb[strings.ToLower(v.Name)].mu = sync.Mutex{}
		var uniqueContainers []string
		for _, container := range containers {
			uniqueContainers = insertUnique(b.lb[strings.ToLower(v.Name)].containers, container.ID)
		}
		b.lb[strings.ToLower(v.Name)].containers = uniqueContainers
	}
	// b.log.Infoc(context.TODO(), "the time this func called was", map[string]interface{}{
	// 	"time": t,
	// })
	// return
	// return lb, nil
}
func insertUnique(slice []string, value string) []string {
	seen := make(map[string]struct{})
	for _, v := range slice {
		seen[v] = struct{}{} // Mark elements already in the slice
	}
	if _, exists := seen[value]; exists {
		return slice
	}
	return append(slice, value)
}

//	func (b *Business) doEvery(d time.Duration, f func(allowedLanguages []*Language, b *Business, t time.Time)) {
//		for x := range time.Tick(d) {
//			f(b.languages, b, x)
//		}
//	}
func (ds *Business) Run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			go func() {
				UpdateLB(ds.languages, ds, time.Now())
			}()
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// periodic task
func (ds *Business) task() {

	// fmt.Printf("Ticker task, %v\n", time.Now())
}
func NewBusiness(log *logger.CustomLogger, delegate *delegate.Delegate, storer Storer, cli *client.Client) *Business {
	business := &Business{}
	langages, err := storer.GetLanguages(context.TODO())
	if err != nil {
		log.Errorc(context.TODO(), "error in getting allowed labuages", map[string]interface{}{
			"error": err.Error(),
		})
	}
	lb, err := NewContainerLB(langages, cli, log)
	if err != nil {
		log.Errorc(context.TODO(), "error in creating LB", map[string]interface{}{
			"error": err.Error(),
		})
	}
	business.cli = cli
	business.lb = lb
	business.log = log
	business.storer = storer
	business.delegate = delegate
	business.languages = langages
	business.containerSpec = make(map[string]ContainerSpec)
	// go business.doEvery(5*time.Second, UpdateLB)
	go business.Run(context.TODO())
	return business
	// return &Business{
	// 	log:           log,
	// 	delegate:      delegate,
	// 	storer:        storer,
	// 	cli:           cli,
	// 	containerSpec: make(map[string]ContainerSpec),
	// 	lb:            lb,
	// }
}
func (b *Business) ExecuteCode(ctx context.Context, path, language, uid, qid, ext string) (*pb.ExecutionResponse, error) {
	var execResponse pb.ExecutionResponse
	// get container spec
	specs, err := b.getContainerSpec(language)
	if err != nil {
		b.log.Errorc(ctx, "error in getting specs", map[string]interface{}{
			"error": err.Error(),
		})
	}
	// 	return &execResponse, err
	// }
	b.log.Infoc(ctx, "container spec values", map[string]interface{}{
		"containerSpec": specs,
	})
	// // create a tarfile from the temp file
	// buf, err := b.readTempFile(path)
	// if err != nil {
	// 	return &execResponse, err
	// }
	// // Copy the file to the existing container
	// err = b.cli.CopyToContainer(context.Background(), specs.ID, "app/", buf, container.CopyToContainerOptions{})
	// if err != nil {
	// 	return &execResponse, err
	// }

	// // // Execute the Python script in the existing container
	// execConfig := container.ExecOptions{
	// 	Cmd:          []string{"python", filepath.Join("/app", filepath.Base(path))},
	// 	AttachStdout: true,
	// 	AttachStderr: true,
	// 	Tty:          true,
	// }

	// execID, err := b.cli.ContainerExecCreate(context.Background(), specs.ID, execConfig)
	// if err != nil {
	// 	return &execResponse, err
	// }

	// err = cli.ContainerExecStart(context.Background(), execID.ID, container.ExecStartOptions{})
	// if err != nil {
	// 	return "", err
	// }
	startTime := time.Now()
	// res, err := b.cli.ContainerExecAttach(context.TODO(), execID.ID, container.ExecAttachOptions{
	// 	Tty: true,
	// })
	// if err != nil {
	// 	return &execResponse, err
	// }
	res, containerID, err := b.executeWithRetry(ctx, path, language, ext, 5)
	if err != nil {
		return &execResponse, err
	}
	endTime := time.Since(startTime)
	// _ = res
	stats, err := b.cli.ContainerStatsOneShot(ctx, containerID)
	if err != nil {
		return &execResponse, err
	}
	defer stats.Body.Close()
	// defer res.Close()
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
	go b.performCleanup(filepath.Base(path), containerID)
	execResponse.ContainerID = containerID
	execResponse.Output = res
	// fmt.Println("the log is", string(logBuf.Bytes()))
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
func (b *Business) ActualCodeExecution(containerID, path, cmd string) (bytes.Buffer, error) {
	var execResponse bytes.Buffer
	buf, err := b.readTempFile(path)
	if err != nil {
		return execResponse, err
	}
	// Copy the file to the existing container
	err = b.cli.CopyToContainer(context.Background(), containerID, "app/", buf, container.CopyToContainerOptions{})
	if err != nil {
		return execResponse, err
	}
	command, err := prepareCommand(cmd, filepath.Base(path))
	if err != nil {
		return execResponse, err
	}
	// // Execute the Python script in the existing container
	execConfig := container.ExecOptions{
		Cmd:          []string{"sh", "-c", command},
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
	}

	execID, err := b.cli.ContainerExecCreate(context.Background(), containerID, execConfig)
	if err != nil {
		return execResponse, err
	}

	res, err := b.cli.ContainerExecAttach(context.TODO(), execID.ID, container.ExecAttachOptions{
		Tty: true,
	})
	if err != nil {
		return execResponse, err
	}
	if _, err := execResponse.ReadFrom(res.Reader); err != nil {
		return execResponse, err
	}
	return execResponse, nil
}
func prepareCommand(cmd, filename string) (string, error) {
	var finalResult string
	type CommandData struct {
		FileName string
		Program  string
	}
	if strings.Contains(cmd, ",") {
		programs := strings.Split(filename, ".")
		filename = fmt.Sprintf("/app/%s", filename)
		if len(programs) > 0 {
			program := programs[0]
			data := CommandData{
				FileName: filename,
				Program:  program,
			}
			tmpl, err := template.New("cmd").Parse(cmd)
			if err != nil {
				return "", err
			}
			var result bytes.Buffer

			// Execute the template with the data to generate the final command
			err = tmpl.Execute(&result, data)
			if err != nil {
				return "", err
			}

			finalResult = strings.Replace(result.String(), ",", " ", -1)
		}
		return finalResult, nil
	}
	filename = fmt.Sprintf("/app/%s", filename)
	return fmt.Sprintf("%s %s", cmd, filename), nil
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

// executeWithRetry executes a task with retry logic and exponential backoff
func (b *Business) executeWithRetry(ctx context.Context, codeFilePath, language, ext string, retries int) (string, string, error) {
	var lastError error
	// specs, err := b.getContainerSpec(language)
	// if err != nil {
	// 	b.log.Errorc(ctx, "error in getting specs", map[string]interface{}{
	// 		"error": err.Error(),
	// 	})
	// 	return "", err
	// }
	if !strings.Contains(codeFilePath, ext) {
		err := os.Rename(codeFilePath, fmt.Sprintf("%s%s", codeFilePath, ext))
		if err != nil {
			return "", "", fmt.Errorf("err renaming file: %v", err)
		}
		codeFilePath = fmt.Sprintf("%s%s", codeFilePath, ext)
	}
	for attempt := 1; attempt <= retries; attempt++ {
		containerID, err := b.lb[language].getNextContainer()
		if err != nil {
			return "", "", fmt.Errorf("no containers available: %v", err)
		}
		specs := []LanguageSpecification{}
		err = b.storer.Get(ctx, "langSpecs", &specs)
		if err != nil {
			specs, err = b.storer.GetAllLangSpecs()
			if err != nil {
				return "", containerID, err
			}
			_, err = b.storer.Set(ctx, "langSpecs", specs, 0)
			if err != nil {
				return "", containerID, err
			}
		}
		var cmd string
		for _, v := range specs {
			if strings.EqualFold(strings.ToLower(v.LanguageName), strings.ToLower(language)) {
				cmd = v.Command
			}
		}
		//TODO:get command of the specific language from here and pass it down
		logBuf, err := b.ActualCodeExecution(containerID, codeFilePath, cmd)
		if err == nil {
			out := convertOutput(logBuf)
			return out, containerID, nil // Successful execution
		}
		lastError = err
		b.log.Errorc(ctx, "error in code execution", map[string]interface{}{
			"attempt":     attempt,
			"containerID": containerID,
			"error":       err.Error(),
		})
		b.log.Errorc(ctx, fmt.Sprintf("Attempt %d failed for container %s: %v", attempt, containerID, err), map[string]interface{}{})

		// Exponential backoff
		backoffDuration := time.Duration(rand.Intn(int(math.Pow(2, float64(attempt))))) * time.Second
		b.log.Errorc(ctx, fmt.Sprintf("Retrying in %v...", backoffDuration), map[string]interface{}{})
		time.Sleep(backoffDuration)
	}

	return "", "", fmt.Errorf("execution failed after %d attempts: %v", retries, lastError)
}

func (lb *LoadBalancer) getNextContainer() (string, error) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.containers) == 0 {
		return "", fmt.Errorf("no containers available")
	}

	container := lb.containers[lb.currentIdx]
	lb.currentIdx = (lb.currentIdx + 1) % len(lb.containers)
	return container, nil
}

// Add container to the list
func (lb *LoadBalancer) addContainer(containerID string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.containers = append(lb.containers, containerID)
}

// Remove container from the list
func (lb *LoadBalancer) removeContainer(containerID string) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for i, c := range lb.containers {
		if c == containerID {
			lb.containers = append(lb.containers[:i], lb.containers[i+1:]...)
			break
		}
	}
}

// func (b *Business) updateContainerHealth(language string) {
// 	lb := b.lb[language]
// 	for _, containerID := range lb.containers {
// 		if !b.isContainerHealthy(containerID) {
// 			log.Printf("Container %s is unhealthy. Removing from load balancer.", containerID)
// 			lb.removeContainer(containerID) // lb is the load balancer
// 		} else {
// 			if !lb.contains(containerID) {
// 				log.Printf("Adding healthy container %s to load balancer.", containerID)
// 				lb.addContainer(containerID)
// 			}
// 		}
// 	}
// }

// Add health check function
func (b *Business) isContainerHealthy(containerID string) bool {
	containerInfo, err := b.cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		log.Printf("Error inspecting container %s: %v", containerID, err)
		return false
	}
	return containerInfo.State.Health != nil && containerInfo.State.Health.Status == "healthy"
}

// // Scale up/down based on the load
// func (b *Business) scaleContainers(targetCount int) error {
// 	currentContainers, err := getRunningPythonContainers()
// 	if err != nil {
// 		return err
// 	}

// 	if len(currentContainers) < targetCount {
// 		for i := len(currentContainers); i < targetCount; i++ {
// 			err := b.startNewContainer() // Logic to start a new Python container
// 			if err != nil {
// 				return fmt.Errorf("failed to scale up: %v", err)
// 			}
// 		}
// 	} else if len(currentContainers) > targetCount {
// 		for i := len(currentContainers) - 1; i >= targetCount; i-- {
// 			err := b.stopContainer(currentContainers[i]) // Logic to stop and remove a container
// 			if err != nil {
// 				return fmt.Errorf("failed to scale down: %v", err)
// 			}
// 		}
// 	}

// 	return nil
// }

// startNewContainer starts a new Docker container from a given image
func (b *Business) startNewContainer(imageName string) (string, error) {
	// Pull the image (if not already available)
	_, err := b.cli.ImagePull(context.Background(), imageName, image.PullOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to pull image: %v", err)
	}
	containerConfig := &container.Config{
		Image: imageName,
		Cmd:   []string{"tail", "-f", "/dev/null"}, // Keep the container running, modify as needed
	}

	hostConfig := &container.HostConfig{
		AutoRemove: true, // Automatically remove the container when it's stopped
	}

	resp, err := b.cli.ContainerCreate(context.Background(), containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("failed to create container: %v", err)
	}

	// Start the container
	err = b.cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to start container: %v", err)
	}
	return resp.ID, nil
}

// stopContainer stops and removes the container by its ID
func (b *Business) stopContainer(containerID string) error {
	// Stop the container
	err := b.cli.ContainerStop(context.Background(), containerID, container.StopOptions{})
	if err != nil {
		return fmt.Errorf("failed to stop container %s: %v", containerID, err)
	}

	// Remove the container
	err = b.cli.ContainerRemove(context.Background(), containerID, container.RemoveOptions{
		Force: true, // Force removal if the container is running
	})
	if err != nil {
		return fmt.Errorf("failed to remove container %s: %v", containerID, err)
	}
	return nil
}
