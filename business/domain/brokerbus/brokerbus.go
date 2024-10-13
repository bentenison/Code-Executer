package brokerbus

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"text/template"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
)

type Storer interface {
	GetQuestionTemplate(ctx context.Context, id string) (Question, error)
}

type Business struct {
	log      *logger.CustomLogger
	db       mux.DataSource
	delegate *delegate.Delegate
	storer   Storer
}

func NewBusiness(logger *logger.CustomLogger, ds mux.DataSource, delegate *delegate.Delegate, storer Storer) *Business {
	return &Business{
		log:      logger,
		db:       ds,
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
	b.createCodeTemplate(ctx, question)
	return question, err
}

func (b *Business) createCodeTemplate(ctx context.Context, question Question) {
	tmplt, err := template.New("code").Parse(question.TemplateCode)
	if err != nil {
		b.log.Errorc(ctx, "error creating template from string", map[string]interface{}{
			"error": err,
		})
	}
	f, err := os.OpenFile(fmt.Sprintf("./static/code_%s", question.QuestionId), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		b.log.Errorc(ctx, "error while creating file", map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer f.Close()
	tmplt.Execute(f, question)
}
func decodeSnippet(snippet string) (string, error) {
	snipByte, err := base64.StdEncoding.DecodeString(snippet)
	if err != nil {
		return "", err
	}
	return string(snipByte), nil
}
