package brokerbus

import (
	"time"
)

type TestCase struct {
	Input          interface{} `bson:"input" json:"input"`                     // Input can be of any type
	ExpectedOutput interface{} `bson:"expected_output" json:"expected_output"` // Expected output can be of any type
}

type Question struct {
	QuestionId   string     `bson:"_id" json:"questionId"`
	Title        string     `bson:"title" json:"title"`                 // Title of the problem
	Description  string     `bson:"description" json:"description"`     // Problem description
	TemplateCode string     `bson:"template_code" json:"template_code"` // Code template for user logic
	Language     string     `bson:"language" json:"language"`           // Programming language (e.g., Python)
	LanguageCode string     `bson:"language_code" json:"language_code"` // Language code (e.g., "py")
	TestCases    []TestCase `bson:"test_cases" json:"test_cases"`       // List of test cases with dynamic types
	Difficulty   string     `bson:"difficulty" json:"difficulty"`       // Difficulty level of the problem
	Tags         []string   `bson:"tags" json:"tags"`
	Logic        string     // Tags related to the problem
}

// Submission struct for returning a submission result
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
}

// PerformanceMetrics struct for returning performance metrics
type PerformanceMetrics struct {
	ID            string    `json:"id"`
	SubmissionID  string    `json:"submission_id"`
	ExecutionTime time.Time `json:"execution_time"`
	MemoryUsage   int64     `json:"memory_usage"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CodeExecutionStats struct for returning code execution statistics
type CodeExecutionStats struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	LanguageID    string    `json:"language_id"`
	ExecutionTime time.Time `json:"execution_time"`
	MemoryUsage   int64     `json:"memory_usage"`
	Status        string    `json:"status"`
	ErrorMessage  string    `json:"error_message,omitempty"`
	CodeSnippet   string    `json:"code_snippet"`
	ContainerID   string    `json:"container_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// User struct for returning user information
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Language struct for returning programming language details
type Language struct {
	ID               string    `json:"id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	ContainerID      string    `json:"container_id"`
	ContainerName    string    `json:"container_name"`
	Version          string    `json:"version"`
	DocumentationURL string    `json:"documentation_url,omitempty"`
	IsActive         bool      `json:"is_active"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
