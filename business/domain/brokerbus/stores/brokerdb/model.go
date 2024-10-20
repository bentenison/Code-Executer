package brokerdb

import (
	"database/sql"
	"time"

	"github.com/bentenison/microservice/business/domain/brokerbus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type TestCase struct {
// 	Input          interface{} `bson:"input" json:"input"`                     // Input can be of any type
// 	ExpectedOutput interface{} `bson:"expected_output" json:"expected_output"` // Expected output can be of any type
// }

//	type Question struct {
//		QuestionId   primitive.ObjectID `bson:"_id" json:"questionId"`
//		Title        string             `bson:"title" json:"title"`                 // Title of the problem
//		Description  string             `bson:"description" json:"description"`     // Problem description
//		TemplateCode string             `bson:"template_code" json:"template_code"` // Code template for user logic
//		Language     string             `bson:"language" json:"language"`           // Programming language (e.g., Python)
//		LanguageCode string             `bson:"language_code" json:"language_code"` // Language code (e.g., "py")
//		TestCases    []TestCase         `bson:"test_cases" json:"test_cases"`       // List of test cases with dynamic types
//		Difficulty   string             `bson:"difficulty" json:"difficulty"`       // Difficulty level of the problem
//		Tags         []string           `bson:"tags" json:"tags"`                   // Tags related to the problem
//	}

type Question struct {
	QuestionId        primitive.ObjectID `json:"_id" bson:"_id"`
	Title             string             `json:"title" bson:"title"`
	Description       string             `json:"description" bson:"description"`
	Input             Input              `json:"input" bson:"input"`
	Output            Output             `json:"output" bson:"output"`
	TemplateCode      string             `json:"template_code" bson:"template_code"`
	Language          string             `json:"language" bson:"language"`
	LanguageCode      string             `json:"language_code" bson:"language_code"`
	Difficulty        string             `json:"difficulty" bson:"difficulty"`
	Tags              []string           `json:"tags" bson:"tags"`
	UserLogicTemplate UserLogicTemplate  `json:"user_logic_template" bson:"user_logic_template"`
	TestcaseTemplate  TestcaseTemplate   `json:"testcase_template" bson:"testcase_template"`
	Testcases         []Testcase         `json:"testcases" bson:"testcases"`
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
	Input          int   `json:"input" bson:"input"`
	ExpectedOutput []int `json:"expectedOutput" bson:"expectedOutput"`
}
type Answer struct {
	ID        string     `json:"id"`
	Logic     string     `json:"logic"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	TestCases []Testcase `json:"testcases"`
}

// Submission struct for the 'submissions' table
type SubmissionDB struct {
	ID              sql.NullString `db:"id"`
	UserID          sql.NullString `db:"user_id"`
	LanguageID      sql.NullString `db:"language_id"`
	CodeSnippet     sql.NullString `db:"code_snippet"`
	SubmissionTime  sql.NullTime   `db:"submission_time"`
	ExecutionStatus sql.NullString `db:"execution_status"`
	ResultID        sql.NullString `db:"result_id"`
	IsPublic        sql.NullBool   `db:"is_public"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	UpdatedAt       sql.NullTime   `db:"updated_at"`
}

// PerformanceMetrics struct for the 'performance_metrics' table
type PerformanceMetricsDB struct {
	ID            sql.NullString `db:"id"`
	SubmissionID  sql.NullString `db:"submission_id"`
	ExecutionTime sql.NullTime   `db:"execution_time"`
	MemoryUsage   sql.NullInt64  `db:"memory_usage"`
	Status        sql.NullString `db:"status"`
	CreatedAt     sql.NullTime   `db:"created_at"`
	UpdatedAt     sql.NullTime   `db:"updated_at"`
}

// CodeExecutionStats struct for the 'code_execution_stats' table
type CodeExecutionStatsDB struct {
	ID            sql.NullString `db:"id"`
	UserID        sql.NullString `db:"user_id"`
	LanguageID    sql.NullString `db:"language_id"`
	ExecutionTime sql.NullTime   `db:"execution_time"`
	MemoryUsage   sql.NullInt64  `db:"memory_usage"`
	Status        sql.NullString `db:"status"`
	ErrorMessage  sql.NullString `db:"error_message"`
	CodeSnippet   sql.NullString `db:"code_snippet"`
	ContainerID   sql.NullString `db:"container_id"`
	CreatedAt     sql.NullTime   `db:"created_at"`
	UpdatedAt     sql.NullTime   `db:"updated_at"`
}

// User struct for the 'users' table
type UserDB struct {
	ID           sql.NullString `db:"id"`
	Username     sql.NullString `db:"username"`
	Email        sql.NullString `db:"email"`
	PasswordHash sql.NullString `db:"password_hash"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Role         sql.NullString `db:"role"`
	CreatedAt    sql.NullTime   `db:"created_at"`
	UpdatedAt    sql.NullTime   `db:"updated_at"`
}

// Language struct for the 'languages' table
type LanguageDB struct {
	ID               sql.NullString `db:"id"`
	Code             sql.NullString `db:"code"`
	Name             sql.NullString `db:"name"`
	ContainerID      sql.NullString `db:"container_id"`
	ContainerName    sql.NullString `db:"container_name"`
	Version          sql.NullString `db:"version"`
	DocumentationURL sql.NullString `db:"documentation_url"`
	IsActive         sql.NullBool   `db:"is_active"`
	CreatedAt        sql.NullTime   `db:"created_at"`
	UpdatedAt        sql.NullTime   `db:"updated_at"`
}

func toBusQuestion(q Question) brokerbus.Question {
	busQuestion := brokerbus.Question{}
	busQuestion.QuestionId = q.QuestionId.Hex()
	busQuestion.Title = q.Title
	busQuestion.Description = q.Description
	busQuestion.Input = brokerbus.Input(q.Input)
	busQuestion.Output = brokerbus.Output(q.Output)
	busQuestion.UserLogicTemplate = brokerbus.UserLogicTemplate(q.UserLogicTemplate)
	busQuestion.TestcaseTemplate = brokerbus.TestcaseTemplate(q.TestcaseTemplate)
	busQuestion.Difficulty = q.Difficulty
	busQuestion.TemplateCode = q.TemplateCode
	busQuestion.Language = q.Language
	busQuestion.LanguageCode = q.LanguageCode
	busQuestion.Tags = q.Tags
	busQuestion.Testcases = addTestCases(q.Testcases)
	return busQuestion
}
func addTestCases(cases []Testcase) []brokerbus.Testcase {
	out := []brokerbus.Testcase{}
	for _, v := range cases {
		out = append(out, brokerbus.Testcase(v))
	}
	return out
}
func toBusLanguage(lang *LanguageDB) *brokerbus.Language {
	var lg brokerbus.Language
	lg.ID = lang.ID.String
	lg.Code = lang.Code.String
	lg.Name = lang.Name.String
	lg.ContainerID = lang.ContainerID.String
	lg.ContainerName = lang.ContainerName.String
	lg.Version = lang.Version.String
	lg.DocumentationURL = lang.DocumentationURL.String
	lg.IsActive = lang.IsActive.Bool
	lg.UpdatedAt = lang.UpdatedAt.Time
	lg.CreatedAt = lang.UpdatedAt.Time
	return &lg
}
func toBusLanguages(lang []LanguageDB) []*brokerbus.Language {
	var langs []*brokerbus.Language
	for _, v := range lang {
		lg := toBusLanguage(&v)
		langs = append(langs, lg)
	}
	return langs
}
