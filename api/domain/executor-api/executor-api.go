package executorapi

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	pb "github.com/bentenison/microservice/api/domain/executor-api/grpc/proto"
	"github.com/bentenison/microservice/app/domain/executorapp"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/version"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

type api struct {
	app *executorapp.App
	log *logger.CustomLogger
	pb.UnimplementedExecutorServiceServer
}

// var grpcApi *api

func newAPI(app *executorapp.App, log *logger.CustomLogger) *api {
	api := &api{
		app: app,
		log: log,
	}
	// grpcApi = api
	return api
}

// func NewExecutorServer(log *logger.CustomLogger) *api {
// 	return &api{
// 		log: log,
// 	}
// }

func (s *api) HandleExecution(stream grpc.ClientStreamingServer[pb.ExecutionRequest, pb.ExecutionResponse]) error {
	// log.Println("context", stream.Context())
	_, err := os.Stat("./static")
	if err != nil {
		err := os.MkdirAll("./static", 0755)
		if err != nil {
			s.log.Errorc(stream.Context(), "error while creating dir", map[string]interface{}{
				"error": err.Error(),
			})
			return err
		}
	}
	filePath := filepath.Join("static", fmt.Sprintf("code_%s.py", stream.Context().Value("tracectx")))
	file, err := os.Create(filePath)
	if err != nil {
		s.log.Errorc(stream.Context(), "error while creating file", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	defer file.Close()
	var questionId, userId string
	for {
		chunk, err := stream.Recv()
		qid := chunk.GetQid()
		uid := chunk.GetUid()
		if uid != "" && qid != "" {
			questionId = qid
			userId = uid
		}
		s.log.Errorc(stream.Context(), "questionID", map[string]interface{}{
			"quid": qid,
			"uid":  uid,
		})
		if err == io.EOF {
			if version.BuildInfo.LanguageName == "" {
				res, err := s.app.HandleExecution(stream.Context(), filePath, "python", questionId, userId)
				if err != nil {
					s.log.Errorc(stream.Context(), "error while executing the code", map[string]interface{}{
						"error": err.Error(),
					})
					return err
				}
				return stream.SendAndClose(res)
				// return stream.SendAndClose(&pb.ExecutionResponse{
				// 	Message: "File uploaded successfully",
				// 	Success: true,
			}
		}
		// if chunk.Qid != "" && chunk.Uid != "" {
		// }
		// log.Println(chunk.Qid)
		// log.Println(chunk.Uid)
		if err != nil {
			return err
		}
		_, err = file.Write(chunk.Content)
		if err != nil {
			return err
		}
	}
}

func (s *api) handleSubmission(c *gin.Context) {

}
