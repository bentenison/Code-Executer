package brokerapi

import (
	"time"

	brokerapp "github.com/bentenison/microservice/app/domain/broker-app"
	"github.com/google/uuid"
)

type SubmissionPayload struct {
	LanguageCode   string `json:"language_code,omitempty"`
	Language       string `json:"language,omitempty"`
	CodeSnippet    string `json:"code_snippet,omitempty"`
	UserId         string `json:"user_id,omitempty"`
	QuestionId     string `json:"question_id,omitempty"`
	SubmissionTime string `json:"submission_time,omitempty"`
}

func toAppSubmission(payload SubmissionPayload) brokerapp.Submission {
	submission := brokerapp.Submission{
		LanguageID:  payload.LanguageCode,
		UserID:      payload.UserId,
		CodeSnippet: payload.CodeSnippet,
		QuestionId:  payload.QuestionId,
	}
	submission.ID = uuid.NewString()
	submission.SubmissionTime = time.Now()
	submission.CreatedAt = time.Now()
	submission.UpdatedAt = time.Now()
	return submission
}
