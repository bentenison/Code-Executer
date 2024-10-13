package executorapi

import (
	"io"
	"os"
	"path/filepath"

	pb "github.com/bentenison/microservice/api/domain/executor-api/grpc/proto"
	"github.com/bentenison/microservice/app/domain/executorapp"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type api struct {
	app *executorapp.App
	log *logger.CustomLogger
	pb.UnimplementedExecutorServiceServer
}

var grpcApi *api

func newAPI(app *executorapp.App, log *logger.CustomLogger) *api {
	api := &api{
		app: app,
		log: log,
	}
	grpcApi = api
	return api
}
func NewExecutorServer(log *logger.CustomLogger) *api {
	return &api{
		log: log,
	}
}

func (s *api) HandleExecution(stream grpc.ClientStreamingServer[pb.ExecutionRequest, pb.ExecutionResponse]) error {
	// log.Println("context", stream.Context())
	_, err := os.Stat("./static")
	if err != nil {
		err := os.MkdirAll("./static", 0664)
		if err != nil {
			s.log.Errorc(stream.Context(), "error while creating dir", map[string]interface{}{
				"error": err.Error(),
			})
			return err
		}
	}
	filePath := filepath.Join("static", "code.py")
	file, err := os.Create(filePath)
	if err != nil {
		s.log.Errorc(stream.Context(), "error while creating file", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	defer file.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			// return stream.SendAndClose(&pb.ExecutionResponse{
			// 	Message: "File uploaded successfully",
			// 	Success: true,
			// })
		}
		if err != nil {
			return err
		}

		_, err = file.Write(chunk.Content)
		if err != nil {
			return err
		}
	}
}

func (api *api) handleSubmission(c *gin.Context) {

}
