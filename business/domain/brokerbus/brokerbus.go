package brokerbus

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/template"
	"time"

	authpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/authclient/proto"
	execpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto"
	"github.com/bentenison/microservice/api/sdk/http/client"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
)

type Storer interface {
	GetQuestionTemplate(ctx context.Context, id string) (Question, error)
	AddSubmission(ctx context.Context, submission *Submission) (string, error)
	AddExecutionStats(ctx context.Context, newStat *CodeExecutionStats) (string, error)
	GetLanguages(ctx context.Context) ([]*Language, error)
	GetAllQuestionsDAO(ctx context.Context) ([]Question, error)
	GetAllAnswersDAO(ctx context.Context) ([]Answer, error)
	GetAnswerById(ctx context.Context, id string) (Answer, error)
	Get(ctx context.Context, key string, res any) error
	Set(ctx context.Context, key string, val any, ttl time.Duration) (string, error)
}

type Business struct {
	log      *logger.CustomLogger
	delegate *delegate.Delegate
	storer   Storer
}

func NewBusiness(logger *logger.CustomLogger, delegate *delegate.Delegate, storer Storer) *Business {
	return &Business{
		log:      logger,
		delegate: delegate,
		storer:   storer,
	}
}

func (b *Business) HandleSubmissonService(ctx context.Context, submission Submission, authcli authpb.AuthServiceClient, execcli execpb.ExecutorServiceClient) (*execpb.ExecutionResponse, error) {
	//check if question exists in redis first
	question := Question{}
	err := b.storer.Get(ctx, submission.QuestionId, &question)
	if err != nil {
		b.log.Errorc(ctx, "error while getting data from redis .. going for DB now.", map[string]interface{}{
			"error": err.Error(),
		})
		question, err = b.storer.GetQuestionTemplate(ctx, submission.QuestionId)
		if err != nil {
			b.log.Errorc(ctx, "error while getting template", map[string]interface{}{
				"error": err,
			})
			return nil, err
		}
		res, err := b.storer.Set(ctx, submission.QuestionId, &question, 0)
		if err != nil {
			b.log.Errorc(ctx, "error while setting template in redis", map[string]interface{}{
				"error": err,
				"res":   res,
			})
			return nil, err
		}
	}
	// fmt.Println(data)
	decodedSnippet, err := decodeSnippet(submission.CodeSnippet)
	if err != nil {
		b.log.Errorc(ctx, "error while decoding base64 snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	question.UserLogic = decodedSnippet
	question.TestCases = question.TestcaseTemplate.Code
	err = b.createCodeTemplate(ctx, question, submission.UserID)
	if err != nil {
		b.log.Errorc(ctx, "error while ceating template for the question", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	//TODO: create submission and add to DB
	submission.QuestionId = question.QuestionId
	submission.ExecutionStatus = "EXECUTED"

	id, err := b.storer.AddSubmission(ctx, &submission)
	if err != nil {
		b.log.Errorc(ctx, "error in adding submission", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
		// return "", err
	}
	_ = id
	path := fmt.Sprintf("./static/code_%s_%s.py", question.QuestionId, submission.UserID)
	//TODO: call the executor client to exec code
	res, err := startExecution(execcli, path)
	if err != nil {
		b.log.Errorc(ctx, "error in adding submission", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	stats := createCodeExecutionStats(res, id, submission.UserID, submission.CodeSnippet, submission.LanguageID)
	_, err = b.storer.AddExecutionStats(ctx, stats)
	if err != nil {
		b.log.Errorc(ctx, "error in adding code exec stats", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	//TODO: After result from executor client add the perfromance metrics and code_execution stats to DB
	// b.storer.AddExecutionStats(ctx,)
	return res, err
}
func (b *Business) HandleCodeRun(ctx context.Context, submission Submission, authcli authpb.AuthServiceClient, execcli execpb.ExecutorServiceClient) (*execpb.ExecutionResponse, error) {
	question, err := b.storer.GetQuestionTemplate(ctx, submission.QuestionId)
	if err != nil {
		b.log.Errorc(ctx, "error while getting template", map[string]interface{}{
			"error": err,
		})
		return nil, err
	}
	decodedSnippet, err := decodeSnippet(submission.CodeSnippet)
	if err != nil {
		b.log.Errorc(ctx, "error while decoding base64 snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	question.UserLogic = decodedSnippet
	// question.TestCases = question.TestcaseTemplate.Code
	err = b.createCodeTemplate(ctx, question, submission.UserID)
	if err != nil {
		b.log.Errorc(ctx, "error while ceating template for the question", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	//TODO: create submission and add to DB
	submission.QuestionId = question.QuestionId
	submission.ExecutionStatus = "EXECUTED"

	id, err := b.storer.AddSubmission(ctx, &submission)
	if err != nil {
		b.log.Errorc(ctx, "error in adding submission", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
		// return "", err
	}
	_ = id
	path := fmt.Sprintf("./static/code_%s_%s.py", question.QuestionId, submission.UserID)
	//TODO: call the executor client to exec code
	res, err := startExecution(execcli, path)
	if err != nil {
		b.log.Errorc(ctx, "error in adding submission", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	stats := createCodeExecutionStats(res, id, submission.UserID, submission.CodeSnippet, submission.LanguageID)
	_, err = b.storer.AddExecutionStats(ctx, stats)
	if err != nil {
		b.log.Errorc(ctx, "error in adding code exec stats", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	//TODO: After result from executor client add the perfromance metrics and code_execution stats to DB
	// b.storer.AddExecutionStats(ctx,)
	return res, err
}
func (b *Business) HandleAuthentication(ctx context.Context, username, password string) (string, error) {
	var cred struct {
		UserName string `json:"username,omitempty" bson:"userName,omitempty"`
		Password string `json:"password,omitempty" bson:"password,omitempty"`
	}
	cred.UserName = username
	cred.Password = password
	res, err := client.DoRequest("http://localhost:8001/auth/authenticate", "POST", cred)
	if err != nil {
		b.log.Errorc(ctx, "error while connecting auth-service", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	var token string
	err = json.Unmarshal(res, &token)
	if err != nil {
		b.log.Errorc(ctx, "error in unmarshaling res from auth-service", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	return token, nil
}
func (b *Business) HandleAuthorization(ctx context.Context, token string) (Claims, error) {
	var claims Claims
	res, err := client.DoRequest("http://localhost:8001/auth/authorize", "POST", token)
	if err != nil {
		b.log.Errorc(ctx, "error while connecting auth-service", map[string]interface{}{
			"error": err.Error(),
		})
		return claims, err
	}
	err = json.Unmarshal(res, &claims)
	if err != nil {
		b.log.Errorc(ctx, "error in unmarshaling res from auth-service", map[string]interface{}{
			"error": err.Error(),
		})
		return claims, err
	}
	return claims, nil
}
func (b *Business) HandleCreation(ctx context.Context, user UserPayload) (string, error) {
	var uid string
	res, err := client.DoRequest("http://localhost:8001/auth/create", "POST", user)
	if err != nil {
		b.log.Errorc(ctx, "error while connecting auth-service", map[string]interface{}{
			"error": err.Error(),
		})
		return uid, err
	}
	err = json.Unmarshal(res, &uid)
	if err != nil {
		b.log.Errorc(ctx, "error in unmarshaling res from auth-service", map[string]interface{}{
			"error": err.Error(),
		})
		return uid, err
	}
	return uid, nil
}
func (b *Business) createCodeTemplate(ctx context.Context, question Question, userId string) error {
	tmplt, err := template.New("code").Parse(question.ExecTemplate)
	if err != nil {
		b.log.Errorc(ctx, "error creating template from string", map[string]interface{}{
			"error": err,
		})
		return err
	}
	f, err := os.OpenFile(fmt.Sprintf("./static/code_%s_%s.py", question.QuestionId, userId), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		b.log.Errorc(ctx, "error while creating file", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	defer f.Close()
	err = tmplt.Execute(f, question)
	if err != nil {
		return err
	}
	return nil
}
func decodeSnippet(snippet string) (string, error) {
	snipByte, err := base64.StdEncoding.DecodeString(snippet)
	if err != nil {
		return "", err
	}
	return string(snipByte), nil
}
func startExecution(exec execpb.ExecutorServiceClient, path string) (*execpb.ExecutionResponse, error) {
	stream, err := exec.HandleExecution(context.Background())
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, 1024) // 1 KB chunks
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		err = stream.Send(&execpb.ExecutionRequest{Content: buf[:n], Uid: "abc123", Qid: "pqr123"})
		if err != nil {
			return nil, err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *Business) GetSingleQuestion(ctx context.Context, id string) (Question, error) {
	return b.storer.GetQuestionTemplate(ctx, id)
}
func (b *Business) GetAllQuestions(ctx context.Context) ([]Question, error) {
	return b.storer.GetAllQuestionsDAO(ctx)
}
func (b *Business) GetAllAnswers(ctx context.Context) ([]Answer, error) {
	return b.storer.GetAllAnswersDAO(ctx)
}
func (b *Business) GetAnswerByQuestion(ctx context.Context, id string) (Answer, error) {
	return b.storer.GetAnswerById(ctx, id)
}
