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

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"password_hash"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserPayload struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Role         string `json:"role"`
}
type token struct {
	Token string `json:"token,omitempty"`
}
