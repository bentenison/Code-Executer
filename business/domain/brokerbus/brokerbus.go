package brokerbus

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/bentenison/microservice/api/sdk/http/client"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
)

type Storer interface {
	GetQuestionTemplate(ctx context.Context, id string) (Question, error)
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

func (b *Business) HandleSubmissonService(ctx context.Context, submission Submission) (Question, error) {
	question, err := b.storer.GetQuestionTemplate(ctx, submission.QuestionId)
	if err != nil {
		b.log.Errorc(ctx, "error while getting template", map[string]interface{}{
			"error": err,
		})
		return Question{}, err
	}
	decodedSnippet, err := decodeSnippet(submission.CodeSnippet)
	if err != nil {
		b.log.Errorc(ctx, "error while decoding base64 snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return Question{}, err
	}
	question.Logic = decodedSnippet
	err = b.createCodeTemplate(ctx, question, submission.UserID)
	if err != nil {
		b.log.Errorc(ctx, "error while ceating template for the question", map[string]interface{}{
			"error": err.Error(),
		})
	}
	//TODO: create submission and add to DB
	//TODO: call the executor client to exec code
	//TODO: After result from executor client add the perfromance metrics and code_execution stats to DB
	return question, err
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
	tmplt, err := template.New("code").Parse(question.TemplateCode)
	if err != nil {
		b.log.Errorc(ctx, "error creating template from string", map[string]interface{}{
			"error": err,
		})
		return err
	}
	f, err := os.OpenFile(fmt.Sprintf("./static/code_%s_%s", question.QuestionId, userId), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
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
