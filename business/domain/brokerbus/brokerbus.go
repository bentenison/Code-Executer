package brokerbus

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/bentenison/microservice/api/domain/broker-api/grpc/adminclient/proto/admClient"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/authclient/proto/authCli"
	"github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto/execClient"
	"github.com/bentenison/microservice/api/sdk/http/client"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/async/rabbit/rabbitproducer"
	"github.com/bentenison/microservice/foundation/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// Precompile regex patterns at the package level
var (
	ifPattern      = regexp.MustCompile(`if\s*\(`)
	forPattern     = regexp.MustCompile(`for\s*\(`)
	switchPattern  = regexp.MustCompile(`switch\s*\(`)
	varNamePattern = regexp.MustCompile(`[a-z]+[A-Z]`)
)

type Storer interface {
	GetQuestionTemplate(ctx context.Context, id string) (Question, error)
	AddSubmission(ctx context.Context, submission *Submission) (string, error)
	AddExecutionStats(ctx context.Context, newStat *CodeExecutionStats) (string, error)
	GetLanguages(ctx context.Context) ([]*Language, error)
	GetAllQuestionsDAO(ctx context.Context) ([]Question, error)
	GetAllAnswersDAO(ctx context.Context) ([]Answer, error)
	GetAnswerById(ctx context.Context, id string) (Answer, error)
	UpdateQCQuestion(ctx context.Context, id string) (*mongo.UpdateResult, error)
	GetQuestionTemplates(ctx context.Context) ([]Question, error)
	Get(ctx context.Context, key string, res any) error
	Set(ctx context.Context, key string, val any, ttl time.Duration) (string, error)
	CreateSnippet(ctx context.Context, snippet *CodeSnippet) (*mongo.InsertOneResult, error)
	GetSnippetById(ctx context.Context, id string) (*CodeSnippet, error)
	GetAllByUser(ctx context.Context, userId string) ([]CodeSnippet, error)
	GetDBQuestByID(ctx context.Context, questionID string) (*SQLQuestion, error)
}

type Business struct {
	log            *logger.CustomLogger
	delegate       *delegate.Delegate
	storer         Storer
	rabbitProducer *rabbitproducer.Producer
	admClient      admClient.AdminServiceClient
}

func NewBusiness(logger *logger.CustomLogger, delegate *delegate.Delegate, storer Storer, rabbitProducer *rabbitproducer.Producer, admclient admClient.AdminServiceClient) *Business {
	return &Business{
		log:            logger,
		delegate:       delegate,
		storer:         storer,
		rabbitProducer: rabbitProducer,
		admClient:      admclient,
	}
}

func (b *Business) HandleSubmissonService(ctx context.Context, submission Submission, authcli authCli.AuthServiceClient, execcli execClient.ExecutorServiceClient) (*execClient.ExecutionResponse, error) {
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
	question.FileExtension = submission.FileExtension
	// fmt.Println(data)
	decodedSnippet, err := decodeSnippet(submission.CodeSnippet)
	if err != nil {
		b.log.Errorc(ctx, "error while decoding base64 snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	switch question.Language {
	case "python":
		question.UserLogic = decodedSnippet
		question.TestCases = question.TestcaseTemplate.Code
	case "java":
		question.UserLogic = decodedSnippet
		question.TestCases = question.TestcaseTemplate.Code
		question.ClassName = "Main"
	default:
		question.UserLogic = decodedSnippet
		question.TestCases = question.TestcaseTemplate.Code
	}
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
	path := fmt.Sprintf("./static/code_%s_%s%s", question.QuestionId, submission.UserID, submission.FileExtension)
	//TODO: call the executor client to exec code
	res, err := startExecution(execcli, path, question.Language, submission.FileExtension)
	if err != nil {
		b.log.Errorc(ctx, "error in adding submission", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	var analytics UserProgrammingAnalytics
	if strings.Contains(strings.ToLower(res.Output), "true") {
		testCases, err := json.Marshal(question.Testcases)
		if err != nil {
			b.log.Errorc(ctx, "error in marshaling testcases", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}
		res.Output = string(testCases)
		codeQuality, err := evaluateCodeQuality(decodedSnippet)
		if err != nil {
			b.log.Errorc(ctx, "error in evaluating code quality", map[string]interface{}{
				"error": err.Error(),
			}) // return nil, err
		}
		isCorrect := true
		res.IsCorrect = isCorrect
		res.IsRunOnly = false
		// execTime, err := strconv.ParseFloat(res.ExecTime, 64)
		// if err != nil {
		// 	b.log.Errorc(ctx, "error in converting execTime", map[string]interface{}{
		// 		"error": err.Error(),
		// 	})
		// }
		analytics = createAnalyticsObj(codeQuality, res.ExecTime, isCorrect, question.Difficulty, "submission", question.Language, question.QuestionId, submission.UserID, question.Title, question.Tags)
		b.log.Errorc(ctx, fmt.Sprintf("the code quality is %f and answer is %v", codeQuality, isCorrect), map[string]interface{}{})
		result, err := b.admClient.CompleteQuestion(context.TODO(), &admClient.CompleteQuestionRequest{
			ChallengeId: submission.ChallengeID,
			QuestionId:  submission.QuestionId,
			IsCorrect:   isCorrect,
			Language:    question.Language,
			IsChallenge: submission.IsChallenge,
			CodeQuality: codeQuality,
			UserId:      submission.UserID,
		})
		if err != nil {
			b.log.Errorc(ctx, "error in CompleteQuestion RPC:", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}
		_ = result
		_, err = b.admClient.UpdateUserMetrics(context.TODO(), &admClient.UpdateUserMetricsRequest{
			ChallengeId: submission.ChallengeID,
			QuestionId:  submission.QuestionId,
			IsCorrect:   isCorrect,
			Language:    question.Language,
			IsChallenge: submission.IsChallenge,
			CodeQuality: codeQuality,
			UserId:      submission.UserID,
		})
		if err != nil {
			b.log.Errorc(ctx, "error in UpdateUserMetrics RPC:", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}
		_, err = b.admClient.UpdateUserPerformance(context.TODO(), &admClient.UpdateUserPerformanceRequest{
			ChallengeId: submission.ChallengeID,
			QuestionId:  submission.QuestionId,
			IsCorrect:   isCorrect,
			Language:    question.Language,
			IsChallenge: submission.IsChallenge,
			CodeQuality: codeQuality,
			UserId:      submission.UserID,
		})
		if err != nil {
			b.log.Errorc(ctx, "error in UpdateUserPerformance RPC:", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}

	} else {
		// if strings.ToLower(res.Output) != "" {
		// if strings.Contains(res.GetOutput(), "error") {
		res.Output = "false"
		codeQuality, err := evaluateCodeQuality(decodedSnippet)
		if err != nil {
			b.log.Errorc(ctx, "error in evaluating code quality", map[string]interface{}{
				"error": err.Error(),
			}) // return nil, err
		}
		isCorrect := false
		res.IsCorrect = isCorrect
		res.IsRunOnly = false
		// execTime, err := strconv.ParseFloat(res.ExecTime, 64)
		// if err != nil {
		// 	b.log.Errorc(ctx, "error in converting execTime", map[string]interface{}{
		// 		"error": err.Error(),
		// 	})
		// }
		analytics = createAnalyticsObj(codeQuality, res.ExecTime, isCorrect, question.Difficulty, "submission", question.Language, question.QuestionId, submission.UserID, question.Title, question.Tags)

		b.log.Errorc(ctx, fmt.Sprintf("the code quality is %f and answer is %v", codeQuality, isCorrect), map[string]interface{}{})
		result, err := b.admClient.CompleteQuestion(context.TODO(), &admClient.CompleteQuestionRequest{
			ChallengeId: submission.ChallengeID,
			QuestionId:  submission.QuestionId,
			IsCorrect:   isCorrect,
			Language:    question.Language,
			IsChallenge: submission.IsChallenge,
			CodeQuality: codeQuality,
			UserId:      submission.UserID,
		})
		if err != nil {
			b.log.Errorc(ctx, "error in CompleteQuestion RPC:", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}
		_ = result
		_, err = b.admClient.UpdateUserMetrics(context.TODO(), &admClient.UpdateUserMetricsRequest{
			ChallengeId: submission.ChallengeID,
			QuestionId:  submission.QuestionId,
			IsCorrect:   isCorrect,
			Language:    question.Language,
			IsChallenge: submission.IsChallenge,
			CodeQuality: codeQuality,
			UserId:      submission.UserID,
		})
		if err != nil {
			b.log.Errorc(ctx, "error in UpdateUserMetrics RPC:", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}
		_, err = b.admClient.UpdateUserPerformance(context.TODO(), &admClient.UpdateUserPerformanceRequest{
			ChallengeId: submission.ChallengeID,
			QuestionId:  submission.QuestionId,
			IsCorrect:   isCorrect,
			Language:    question.Language,
			IsChallenge: submission.IsChallenge,
			CodeQuality: codeQuality,
			UserId:      submission.UserID,
		})
		if err != nil {
			b.log.Errorc(ctx, "error in UpdateUserPerformance RPC:", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}
	}
	// Output         string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	// ExecTime       string `protobuf:"bytes,2,opt,name=execTime,proto3" json:"execTime,omitempty"`
	// RamUsed        string `protobuf:"bytes,3,opt,name=ramUsed,proto3" json:"ramUsed,omitempty"`
	// CpuStats       string `protobuf:"bytes,4,opt,name=cpuStats,proto3" json:"cpuStats,omitempty"`
	// TotalRAM       string `protobuf:"bytes,5,opt,name=totalRAM,proto3" json:"totalRAM,omitempty"`
	// PercetRAMUsage string `protobuf:"bytes,6,opt,name=percetRAMUsage,proto3" json:"percetRAMUsage,omitempty"`
	// ContainerID    string `protobuf:"bytes,7,opt,name=containerID,proto3" json:"containerID,omitempty"`

	stats := createCodeExecutionStats(res, id, submission.UserID, submission.CodeSnippet, submission.LanguageID)
	//TODO: After result from executor client add the perfromance metrics and code_execution stats to DB
	err = b.rabbitProducer.ProduceMessage("code_execution_stats", *stats)
	if err != nil {
		b.log.Errorc(ctx, "error in adding code exec stats to elasticSearch", map[string]interface{}{
			"error": err.Error(),
		})
		// return nil, err
	}
	//TODO: After result from executor client add the perfromance metrics and code_execution stats to DB
	err = b.rabbitProducer.ProduceMessage("programming_questions", analytics)
	if err != nil {
		b.log.Errorc(ctx, "error in adding programming_questions stats to elasticSearch", map[string]interface{}{
			"error": err.Error(),
		})
		// return nil, err
	}
	_, err = b.storer.AddExecutionStats(ctx, stats)
	if err != nil {
		b.log.Errorc(ctx, "error in adding code exec stats", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	// b.storer.AddExecutionStats(ctx,)
	return res, err
}
func (b *Business) HandleQcService(ctx context.Context, question Question, authcli authCli.AuthServiceClient, execcli execClient.ExecutorServiceClient) (*execClient.ExecutionResponse, error) {
	switch question.Language {
	case "python":
		question.UserLogic = question.Answer.Logic
		question.TestCases = question.TestcaseTemplate.Code
	case "java":
		question.UserLogic = question.Answer.Logic
		question.TestCases = question.TestcaseTemplate.Code
		question.ClassName = "Main"
	default:
		question.UserLogic = question.Answer.Logic
		question.TestCases = question.TestcaseTemplate.Code
	}
	err := b.createCodeTemplateForQC(ctx, question)
	if err != nil {
		b.log.Errorc(ctx, "error while ceating template for the question", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	//TODO: create submission and add to DB
	// submission.QuestionId = question.QuestionId
	// submission.ExecutionStatus = "EXECUTED"

	path := fmt.Sprintf("./static/code_%s_%s%s", question.QuestionId, question.QuestionId, question.FileExtension)
	//TODO: call the executor client to exec code
	res, err := startExecution(execcli, path, question.Language, question.FileExtension)
	if err != nil {
		b.log.Errorc(ctx, "error in adding submission", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	if strings.ToLower(res.Output) != "" {
		if strings.Contains(res.GetOutput(), "error") {
			res.Output = "false"
		}
	}
	if strings.Contains(strings.ToLower(res.Output), "true") {
		res, err := b.storer.UpdateQCQuestion(ctx, question.QuestionId)
		if err != nil {
			b.log.Errorc(ctx, "error in adding code exec stats", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}
		_ = res
	}
	stats := createCodeExecutionStats(res, question.QuestionId, "Q-CREATOR", question.Answer.Logic, question.Language)
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
func (b *Business) HandleCodeRun(ctx context.Context, submission Submission, authcli authCli.AuthServiceClient, execcli execClient.ExecutorServiceClient) (*execClient.ExecutionResponse, error) {
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
	decodedSnippet, err := decodeSnippet(submission.CodeSnippet)
	if err != nil {
		b.log.Errorc(ctx, "error while decoding base64 snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	question.FileExtension = submission.FileExtension
	question.UserLogic = decodedSnippet
	question.FileExtension = submission.FileExtension
	// question.TestCases = question.TestcaseTemplate.Code
	switch question.Language {
	case "python":
		question.UserLogic = decodedSnippet
		question.TestCases = question.TestcaseTemplate.Code
	case "java":
		question.UserLogic = decodedSnippet
		question.TestCases = question.TestcaseTemplate.Code
		question.ClassName = "Main"
	default:
		question.UserLogic = decodedSnippet
		// question.TestCases = question.TestcaseTemplate.Code
	}
	err = b.createCodeTemplatetoRun(ctx, question, submission.UserID)
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
	//TODO: how we will handle diffrent languages and their respective extensions
	path := fmt.Sprintf("./static/code_%s_%s%s", question.QuestionId, submission.UserID, submission.FileExtension)
	//TODO: call the executor client to exec code
	res, err := startExecution(execcli, path, question.Language, submission.FileExtension)
	if err != nil {
		b.log.Errorc(ctx, "error in adding submission", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}

	// res.Output = "false"
	var isCorrect bool
	codeQuality, err := evaluateCodeQuality(decodedSnippet)
	if err != nil {
		b.log.Errorc(ctx, "error in evaluating code quality", map[string]interface{}{
			"error": err.Error(),
		}) // return nil, err
	}
	// isCorrect := false
	// execTime, err := strconv.ParseFloat(res.ExecTime, 64)
	// if err != nil {
	// 	b.log.Errorc(ctx, "error in converting execTime", map[string]interface{}{
	// 		"error": err.Error(),
	// 	})
	// }
	analytics := createAnalyticsObj(codeQuality, res.ExecTime, isCorrect, question.Difficulty, "run", question.Language, question.QuestionId, submission.UserID, question.Title, question.Tags)
	if res.Output == "" {
		isCorrect = true
		res.IsCorrect = isCorrect
		res.IsRunOnly = true
		res.Output = "The sample test case has passed successfully. If you believe the code is correct, please proceed with the submission."
	}
	stats := createCodeExecutionStats(res, id, submission.UserID, submission.CodeSnippet, submission.LanguageID)
	err = b.rabbitProducer.ProduceMessage("code_execution_stats", *stats)
	if err != nil {
		b.log.Errorc(ctx, "error in adding code exec stats to elasticSearch", map[string]interface{}{
			"error": err.Error(),
		})
		// return nil, err
	}
	//TODO: After result from executor client add the perfromance metrics and code_execution stats to DB
	err = b.rabbitProducer.ProduceMessage("programming_questions", analytics)
	if err != nil {
		b.log.Errorc(ctx, "error in adding programming_questions stats to elasticSearch", map[string]interface{}{
			"error": err.Error(),
		})
		// return nil, err
	}
	_, err = b.storer.AddExecutionStats(ctx, stats)
	if err != nil {
		b.log.Errorc(ctx, "error in adding code exec stats", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}

	// result, err := b.admClient.CompleteQuestion(context.TODO(), &admClient.CompleteQuestionRequest{})
	// if err != nil {
	// 	log.Println(err)
	// }
	// _ = result
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
func (b *Business) createCodeTemplateForQC(ctx context.Context, question Question) error {
	tmplt, err := template.New("code").Parse(question.ExecTemplate)
	if err != nil {
		b.log.Errorc(ctx, "error creating template from string", map[string]interface{}{
			"error": err,
		})
		return err
	}
	f, err := os.OpenFile(fmt.Sprintf("./static/code_%s_%s%s", question.QuestionId, question.QuestionId, question.FileExtension), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
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
func (b *Business) createCodeTemplatetoRun(ctx context.Context, question Question, userId string) error {
	tmplt, err := template.New("code").Parse(question.UserLogicTemplate.CodeRunTemplate)
	if err != nil {
		b.log.Errorc(ctx, "error creating template from string", map[string]interface{}{
			"error": err,
		})
		return err
	}
	f, err := os.OpenFile(fmt.Sprintf("./static/code_%s_%s%s", question.QuestionId, userId, question.FileExtension), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
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
func (b *Business) createCodeTemplate(ctx context.Context, question Question, userId string) error {
	tmplt, err := template.New("code").Parse(question.ExecTemplate)
	if err != nil {
		b.log.Errorc(ctx, "error creating template from string", map[string]interface{}{
			"error": err,
		})
		return err
	}
	f, err := os.OpenFile(fmt.Sprintf("./static/code_%s_%s%s", question.QuestionId, userId, question.FileExtension), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
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
func startExecution(exec execClient.ExecutorServiceClient, path, lang, ext string) (*execClient.ExecutionResponse, error) {
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
		err = stream.Send(&execClient.ExecutionRequest{Content: buf[:n], Uid: "abc123", Qid: "pqr123", Lang: lang, FileExt: ext})
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
func (b *Business) GetAllAllowedLanguages(ctx context.Context) ([]*Language, error) {
	return b.storer.GetLanguages(ctx)
}
func (b *Business) GetAllQuestTemplates(ctx context.Context) ([]Question, error) {
	return b.storer.GetQuestionTemplates(ctx)
}
func (b *Business) CreateCodeSnippet(ctx context.Context, snippet *CodeSnippet) (*mongo.InsertOneResult, error) {
	snippet.CreatedAt = time.Now()
	snippet.UpdatedAt = time.Now()
	return b.storer.CreateSnippet(ctx, snippet)
}
func (b *Business) GetSnippetById(ctx context.Context, id string) (*CodeSnippet, error) {
	return b.storer.GetSnippetById(ctx, id)
}
func (b *Business) GetAllSnippetsByUser(ctx context.Context, userId string) ([]CodeSnippet, error) {
	return b.storer.GetAllByUser(ctx, userId)
}
func (b *Business) GetDBQuestById(ctx context.Context, questId string) (*SQLQuestion, error) {
	return b.storer.GetDBQuestByID(ctx, questId)
}
func (b *Business) FormatCode(ctx context.Context, req FormatterRequest) (*FormatterResponse, error) {
	return b.CallFormatterService(ctx, req.Lang, req.Code)
}

func (b *Business) CallFormatterService(ctx context.Context, lang, code string) (*FormatterResponse, error) {
	decodedSnippet, err := decodeSnippet(code)
	if err != nil {
		b.log.Errorc(ctx, "error while decoding base64 snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	// Create the request body
	requestBody := FormatterRequest{
		Lang: lang,
		Code: decodedSnippet,
	}

	// Marshal the request body to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Send a POST request to the formatter service
	resp, err := http.Post("http://localhost:8010/format", "application/json", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to send request to formatter service: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("formatter service returned an error: %s", resp.Status)
	}

	// Decode the response
	var formatterResponse FormatterResponse
	if err := json.NewDecoder(resp.Body).Decode(&formatterResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response from formatter service: %v", err)
	}

	return &formatterResponse, nil
}

// evaluateCodeQuality evaluates the quality of the given code
func evaluateCodeQuality(code string) (float64, error) {
	if code == "" {
		return 0, fmt.Errorf("code cannot be empty")
	}

	complexityScore, err := calculateComplexity(code)
	if err != nil {
		return 0, err
	}

	lineCountScore, err := calculateLineCount(code)
	if err != nil {
		return 0, err
	}

	namingScore, err := checkVariableNaming(code)
	if err != nil {
		return 0, err
	}

	// Average the scores for a final quality score
	qualityScore := (complexityScore + lineCountScore + namingScore) / 3.0
	return qualityScore, nil
}

// calculateComplexity estimates cyclomatic complexity
func calculateComplexity(code string) (float64, error) {
	decisionPoints := 0
	decisionPoints += len(findMatches(code, ifPattern))
	decisionPoints += len(findMatches(code, forPattern))
	decisionPoints += len(findMatches(code, switchPattern))

	// Simplified cyclomatic complexity calculation
	complexity := float64(decisionPoints + 1)
	return 100 - complexity*10, nil // Inverted score
}

// calculateLineCount scores based on the number of lines in the code
func calculateLineCount(code string) (float64, error) {
	lines := strings.Split(code, "\n")
	totalLines := len(lines)

	// Penalize for functions longer than 20 lines
	if totalLines > 20 {
		return 100 - float64(totalLines-20)*5, nil // 5 points penalty for each line over 20
	}
	return 100, nil // Perfect score for short functions
}

// checkVariableNaming checks variable naming conventions
func checkVariableNaming(code string) (float64, error) {
	if varNamePattern.FindString(code) != "" {
		return 50, nil // Penalize for improper naming
	}
	return 100, nil // Perfect score for good naming
}

// findMatches finds regex matches in the code
func findMatches(code string, pattern *regexp.Regexp) []string {
	return pattern.FindAllString(code, -1)
}
func createAnalyticsObj(codequality float64, timeTaken string, isCorrect bool, diff string, ev, language, Qid, Uid, qt string, tags []string) UserProgrammingAnalytics {
	var ua UserProgrammingAnalytics
	timestamp := time.Now().Format(time.RFC3339) // ISO 8601 format

	ua.CodeQuality = codequality
	ua.Correct = isCorrect
	ua.DifficultyLevel = diff
	ua.EventType = ev
	ua.Language = language
	ua.QuestionID = Qid
	ua.UserID = Uid
	ua.TimeTaken = timeTaken
	ua.Tags = tags
	ua.QuestionText = qt
	ua.Timestamp = timestamp
	ua.CreatedAt = timestamp
	ua.UpdatedAt = timestamp
	ua.SubmissionTime = timestamp
	return ua
}
