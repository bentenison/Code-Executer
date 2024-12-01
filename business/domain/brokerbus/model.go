package brokerbus

import (
	"errors"
	"strconv"
	"time"

	execpb "github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto"
	"github.com/golang-jwt/jwt/v5"
)

//	type Question struct {
//		QuestionId   string     `bson:"_id" json:"questionId"`
//		Title        string     `bson:"title" json:"title"`                 // Title of the problem
//		Description  string     `bson:"description" json:"description"`     // Problem description
//		TemplateCode string     `bson:"template_code" json:"template_code"` // Code template for user logic
//		Language     string     `bson:"language" json:"language"`           // Programming language (e.g., Python)
//		LanguageCode string     `bson:"language_code" json:"language_code"` // Language code (e.g., "py")
//		TestCases    []TestCase `bson:"test_cases" json:"test_cases"`       // List of test cases with dynamic types
//		Difficulty   string     `bson:"difficulty" json:"difficulty"`       // Difficulty level of the problem
//		Tags         []string   `bson:"tags" json:"tags"`
//		Logic        string     // Tags related to the problem
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
	TestCases         string            `json:"tstcsc,omitempty" bson:"tstcsc" db:"tstcsc"`
	Answer            Answer            `json:"answer,omitempty" bson:"answer" db:"answer"`
	IsQC              bool              `json:"is_qc,omitempty" bson:"is_qc" db:"is_qc"`
	FileExtension     string            `json:"file_extension,omitempty" db:"file_extension"`
	ClassName         string            `json:"clsnm,omitempty" bson:"clsnm" db:"clsnm"`
}

type Input struct {
	Description string `json:"description" bson:"description"`
	Expected    string `json:"expected" bson:"expected"`
}

type Output struct {
	Description string `json:"description" bson:"description"`
}

type UserLogicTemplate struct {
	Description     string `json:"description" bson:"description"`
	Code            string `json:"code" bson:"code"`
	CodeRunTemplate string `json:"code_run_template,omitempty" bson:"code_run_template"`
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

// Submission struct for returning a submission result
type Submission struct {
	ID              string    `json:"id,omitempty" db:"id"`
	UserID          string    `json:"user_id,omitempty" db:"user_id"`
	LanguageID      string    `json:"language_id,omitempty" db:"language_id"`
	CodeSnippet     string    `json:"code_snippet,omitempty" db:"code_snippet"`
	SubmissionTime  time.Time `json:"submission_time,omitempty" db:"submission_time"`
	ExecutionStatus string    `json:"execution_status,omitempty" db:"execution_status"`
	ResultID        string    `json:"result_id,omitempty" db:"result_id"`
	IsPublic        bool      `json:"is_public,omitempty" db:"is_public"`
	CreatedAt       time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" db:"updated_at"`
	QuestionId      string    `json:"question_id,omitempty" db:"question_id"`
	FileExtension   string    `json:"file_extension,omitempty" db:"file_extension"`
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
	ExecutionTime    float64   `json:"execution_time,omitempty" db:"execution_time"`
	MemoryUsage      int64     `json:"memory_usage,omitempty" db:"memory_usage"`
	TotalMemory      int64     `json:"total_memory,omitempty" db:"total_memory"`
	CPUUsage         int64     `json:"cpu_usage,omitempty" db:"cpu_usage"`
	MemoryPercentage float64   `json:"memory_percentage,omitempty" db:"memory_percentage"`
	CreatedAt        time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at,omitempty" db:"updated_at"`
	ID               string    `json:"id,omitempty" db:"id"`
	UserID           string    `json:"user_id,omitempty" db:"user_id"`
	LanguageID       string    `json:"language_id,omitempty" db:"language_id"`
	Status           string    `json:"status,omitempty" db:"status"`
	ErrorMessage     string    `json:"error_message,omitempty" db:"error_message"`
	CodeSnippet      string    `json:"code_snippet,omitempty" db:"code_snippet"`
	ContainerID      string    `json:"container_id,omitempty" db:"container_id"`
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
	FileExtension    string    `json:"file_extension"`
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
type Claims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func createCodeExecutionStats(pb *execpb.ExecutionResponse, id, uid, codesnippet, langId string) *CodeExecutionStats {
	var codeExecutionStats CodeExecutionStats
	cpuUsage, _ := strconv.Atoi(pb.CpuStats)
	execTime, _ := ConvertToMilliseconds(pb.ExecTime)
	ramUsed, _ := strconv.Atoi(pb.RamUsed)
	totalRAM, _ := strconv.Atoi(pb.TotalRAM)
	percent, _ := strconv.Atoi(pb.PercetRAMUsage)
	codeExecutionStats.CPUUsage = int64(cpuUsage)
	codeExecutionStats.ExecutionTime = execTime
	codeExecutionStats.MemoryPercentage = float64(percent)
	codeExecutionStats.MemoryUsage = int64(ramUsed)
	codeExecutionStats.TotalMemory = int64(totalRAM / 1024)
	codeExecutionStats.ErrorMessage = pb.Output
	codeExecutionStats.ID = id
	codeExecutionStats.UserID = uid
	codeExecutionStats.Status = "EXECUTED"
	codeExecutionStats.CodeSnippet = codesnippet
	codeExecutionStats.LanguageID = langId
	codeExecutionStats.CreatedAt = time.Now()
	codeExecutionStats.UpdatedAt = time.Now()
	return &codeExecutionStats
}

func ConvertToMilliseconds(input string) (float64, error) {
	var numberStr string
	var unit byte

	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' || input[i] == '.' {
			numberStr += string(input[i])
		} else {
			unit = input[i]
			break
		}
	}

	// Convert the numeric part to a float64
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, err
	}

	var factor float64
	switch unit {
	case 'n':
		factor = 1.0 / 1_000_000
	case 'Â':
		factor = 1.0 / 1_000
	case 'm':
		factor = 1.0
	default:
		return 0, errors.New("unknown unit")
	}

	return number * factor, nil
}
