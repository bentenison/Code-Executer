package brokerapp

import (
	"time"

	"github.com/bentenison/microservice/business/domain/brokerbus"
)

type Submission struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	LanguageID      string    `json:"language_id"`
	CodeSnippet     string    `json:"code_snippet"`
	SubmissionTime  time.Time `json:"submission_time"`
	ExecutionStatus string    `json:"execution_status"`
	ResultID        string    `json:"result_id,omitempty"`
	IsPublic        bool      `json:"is_public"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	QuestionId      string    `json:"question_id"`
	FileExtension   string    `json:"file_extension,omitempty"`
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

//	func toBusUser(up *UserPayload) *authbus.User {
//		var u brokerbus.User
//		u.ID = uuid.NewString()
//		u.Username = up.Username
//		u.Email = up.Email
//		u.Password = up.Password
//		u.PasswordHash = up.PasswordHash
//		u.FirstName = up.FirstName
//		u.LastName = up.LastName
//		u.Role = up.Role
//		u.CreatedAt = time.Now()
//		u.UpdatedAt = time.Now()
//		return &u
//	}
type Question struct {
	QuestionId        string            `json:"id" bson:"id"`
	Title             string            `json:"title" bson:"title"`
	Description       string            `json:"description" bson:"description"`
	UserLogic         string            `json:"logic" bson:"logic"`
	Input             Input             `json:"input" bson:"input"`
	Output            Output            `json:"output" bson:"output"`
	TemplateCode      string            `json:"template_code" bson:"template_code"`
	Language          string            `json:"language" bson:"language"`
	LanguageCode      string            `json:"language_code" bson:"language_code"`
	Difficulty        string            `json:"difficulty" bson:"difficulty"`
	Tags              []string          `json:"tags" bson:"tags"`
	UserLogicTemplate UserLogicTemplate `json:"user_logic_template" bson:"user_logic_template"`
	TestcaseTemplate  TestcaseTemplate  `json:"testcase_template" bson:"testcase_template"`
	Testcases         []Testcase        `json:"testcases" bson:"testcases"`
	ExecTemplate      string            `json:"exec_template" bson:"exec_template"`
	TestCases         string
	Answer            Answer `json:"answer,omitempty" bson:"answer" db:"answer"`
	IsQC              bool   `json:"is_qc,omitempty" bson:"is_qc" db:"is_qc"`
	FileExtension     string `json:"file_extension,omitempty" bson:"file_extension" db:"file_extension"`
}

type Input struct {
	Description string `json:"description" bson:"description"`
	Expected    string `json:"expected" bson:"expected"`
}

type Output struct {
	Description string `json:"description" bson:"description"`
}

type UserLogicTemplate struct {
	Description string `json:"description" bson:"description"`
	Code        string `json:"code" bson:"code"`
}

type TestcaseTemplate struct {
	Description string `json:"description" bson:"description"`
	Code        string `json:"code" bson:"code"`
}

type Testcase struct {
	Input          interface{} `json:"input" bson:"input"`
	ExpectedOutput interface{} `json:"expectedOutput" bson:"expectedOutput"`
}

type Answer struct {
	ID        string     `json:"id"`
	Logic     string     `json:"logic"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	TestCases []Testcase `json:"testcases"`
}

func toBusQuestion(q Question) brokerbus.Question {
	var busQuest brokerbus.Question
	busQuest.Answer = toBusAnswer(q.Answer)
	busQuest.FileExtension = q.FileExtension
	busQuest.Title = q.Title
	busQuest.Language = q.Language
	busQuest.QuestionId = q.QuestionId
	busQuest.TemplateCode = q.TemplateCode
	busQuest.ExecTemplate = q.ExecTemplate
	busQuest.TestcaseTemplate = brokerbus.TestcaseTemplate(q.TestcaseTemplate)
	return busQuest
}
func toBusAnswer(a Answer) brokerbus.Answer {
	var busAns brokerbus.Answer
	busAns.ID = a.ID
	busAns.Logic = a.Logic
	return busAns
}
