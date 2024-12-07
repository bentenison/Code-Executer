package executorapi

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	pb "github.com/bentenison/microservice/api/domain/executor-api/grpc/proto"
	"github.com/bentenison/microservice/app/domain/executorapp"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/otel"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"

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
	_, span := otel.AddSpan(stream.Context(), "api.handleexecution.start", attribute.KeyValue{Key: "executor", Value: attribute.Value{}})
	defer span.End()
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
	var questionId, userId, Lang, Ext string
	filePath := filepath.Join("static", fmt.Sprintf("code_%s%s", stream.Context().Value("tracectx"), Ext))
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
		qid := chunk.GetQid()
		uid := chunk.GetUid()
		lang := chunk.GetLang()
		ext := chunk.GetFileExt()
		if uid != "" && qid != "" {
			questionId = qid
			userId = uid
		}
		if lang != "" && ext != "" {
			Lang = lang
			Ext = ext
		}
		s.log.Errorc(stream.Context(), "questionID", map[string]interface{}{
			"quid": qid,
			"uid":  uid,
		})

		if err == io.EOF {
			res, err := s.app.HandleExecution(stream.Context(), filePath, Lang, questionId, userId, Ext)
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
			// }
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

func (s *api) handleSubmission(c *gin.Context) {

}
