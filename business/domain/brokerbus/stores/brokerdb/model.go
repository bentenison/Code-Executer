package brokerdb

import (
	"database/sql"

	"github.com/bentenison/microservice/business/domain/brokerbus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestCase struct {
	Input          interface{} `bson:"input" json:"input"`                     // Input can be of any type
	ExpectedOutput interface{} `bson:"expected_output" json:"expected_output"` // Expected output can be of any type
}

type Question struct {
	QuestionId   primitive.ObjectID `bson:"_id" json:"questionId"`
	Title        string             `bson:"title" json:"title"`                 // Title of the problem
	Description  string             `bson:"description" json:"description"`     // Problem description
	TemplateCode string             `bson:"template_code" json:"template_code"` // Code template for user logic
	Language     string             `bson:"language" json:"language"`           // Programming language (e.g., Python)
	LanguageCode string             `bson:"language_code" json:"language_code"` // Language code (e.g., "py")
	TestCases    []TestCase         `bson:"test_cases" json:"test_cases"`       // List of test cases with dynamic types
	Difficulty   string             `bson:"difficulty" json:"difficulty"`       // Difficulty level of the problem
	Tags         []string           `bson:"tags" json:"tags"`                   // Tags related to the problem
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
	busQuestion.Difficulty = q.Difficulty
	busQuestion.TemplateCode = q.TemplateCode
	busQuestion.Language = q.Language
	busQuestion.LanguageCode = q.LanguageCode
	busQuestion.Tags = q.Tags
	busQuestion.TestCases = addTestCases(q.TestCases)
	return busQuestion
}
func addTestCases(cases []TestCase) []brokerbus.TestCase {
	out := []brokerbus.TestCase{}
	for _, v := range cases {
		out = append(out, brokerbus.TestCase(v))
	}
	return out
}
