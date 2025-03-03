package brokerbus

import (
	"errors"
	"strconv"
	"time"

	"github.com/bentenison/microservice/api/domain/broker-api/grpc/executorclient/proto/execClient"
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
	ID               string    `json:"id,omitempty" db:"id" bson:"id"`
	UserID           string    `json:"user_id,omitempty" db:"user_id" bson:"user_id"`
	LanguageID       string    `json:"language_id,omitempty" db:"language_id" bson:"language_id"`
	CodeSnippet      string    `json:"code_snippet,omitempty" db:"code_snippet" bson:"code_snippet"`
	ExecutionStatus  string    `json:"execution_status,omitempty" db:"execution_status" bson:"execution_status"`
	ResultID         string    `json:"result_id,omitempty" db:"result_id" bson:"result_id"`
	QuestionId       string    `json:"question_id,omitempty" db:"question_id" bson:"question_id"`
	FileExtension    string    `json:"file_extension,omitempty" db:"file_extension" bson:"file_extension"`
	ChallengeID      string    `json:"challenge_id,omitempty" db:"challenge_id" bson:"challenge_id"`
	RunCount         int       `json:"run_count,omitempty" db:"run_count" bson:"run_count"`
	SubmissionTime   time.Time `json:"submission_time,omitempty" db:"submission_time" bson:"submission_time"`
	CreatedAt        time.Time `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
	UpdatedAt        time.Time `json:"updated_at,omitempty" db:"updated_at" bson:"updated_at"`
	IsPublic         bool      `json:"is_public,omitempty" db:"is_public" bson:"is_public"`
	IsChallenge      bool      `json:"is_challenge,omitempty" db:"is_challenge" bson:"is_challenge"`
	HintUsed         bool      `json:"hint_used,omitempty" db:"hint_used" bson:"hint_used"`
	IsQuestionChange bool      `json:"is_question_change,omitempty" db:"is_question_change" bson:"is_question_change"`
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
	Description      string    `json:"description,omitempty" db:"description" bson:"description"` // Optional field
	Tags             []string  `json:"tags,omitempty" db:"tags" bson:"tags"`                      // Optional field (array of strings)
	LogoURL          string    `json:"logo_url,omitempty" db:"logo_url" bson:"logo_url"`
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
type CodeSnippet struct {
	SnippetID string    `bson:"snippet_id,omitempty" json:"snippet_id,omitempty" db:"id"`
	Code      string    `bson:"code" json:"code,omitempty" db:"code"`
	Language  string    `bson:"language" json:"language,omitempty" db:"language"`
	CreatedBy string    `bson:"created_by" json:"created_by,omitempty" db:"created_by"`
	CreatedAt time.Time `bson:"createdAt" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updated_at,omitempty" db:"updated_at"`
}
type FormatterRequest struct {
	Lang string `json:"language"`
	Code string `json:"code"`
}

// Response structure for the formatted code
type FormatterResponse struct {
	FormattedCode string `json:"formatted_code"`
}
type ExampleData struct {
	ActorID    int       `json:"actor_id" bson:"actor_id"`
	FirstName  string    `json:"first_name" bson:"first_name"`
	LastName   string    `json:"last_name" bson:"last_name"`
	LastUpdate time.Time `json:"last_update" bson:"last_update"`
}

type Table struct {
	TableName        string                   `json:"table_name" bson:"table_name"`
	CreateTableQuery string                   `json:"create_table_query" bson:"create_table_query"`
	Columns          []string                 `json:"columns" bson:"columns"`
	ExampleData      []map[string]interface{} `json:"example_data" bson:"example_data"`
	RestoreQuery     string                   `json:"restore_query" bson:"restore_query"`
}

type DBQuestion struct {
	QueryType         string                   `json:"query_type" bson:"query_type"`
	QueryModifiesData bool                     `json:"query_modifies_data" bson:"query_modifies_data"`
	QuestionText      string                   `json:"question_text" bson:"question_text"`
	ExpectedResult    []map[string]interface{} `json:"expected_result" bson:"expected_result"`
	Hints             []string                 `json:"hints" bson:"hints"`
	ExpectedQuery     string                   `json:"expected_query" bson:"expected_query"`
}

type Validation struct {
	StrictOrdering   bool `json:"strict_ordering" bson:"strict_ordering"`
	IgnoreCase       bool `json:"ignore_case" bson:"ignore_case"`
	IgnoreWhitespace bool `json:"ignore_whitespace" bson:"ignore_whitespace"`
}

type SQLQuestion struct {
	ID          string     `json:"id" bson:"id"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Database    string     `json:"database" bson:"database"`
	Difficulty  string     `json:"difficulty" bson:"difficulty"`
	Tags        []string   `json:"tags" bson:"tags"`
	Tables      []Table    `json:"tables" bson:"tables"`
	DBQuestion  DBQuestion `json:"question" bson:"question"`
	Validation  Validation `json:"validation" bson:"validation"`
}

func createCodeExecutionStats(pb *execClient.ExecutionResponse, id, uid, codesnippet, langId string) *CodeExecutionStats {
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

type OverAllUser struct {
	UserID             string    `json:"user_id,omitempty" db:"user_id" bson:"user_id"`
	Username           string    `json:"username,omitempty" db:"username" bson:"username"`
	Rank               int       `json:"rank,omitempty" db:"rank" bson:"rank"`
	CreatedAt          time.Time `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
	AttemptedQuestions []string  `json:"attempted_questions,omitempty" db:"attempted_questions" bson:"attempted_questions"` // List of question IDs user has faced
	SelectedLanguage   string    `json:"selected_language,omitempty" db:"selected_language" bson:"selected_language"`       // User's selected programming language
	NoAttempted        int64     `json:"no_attempted,omitempty" db:"no_attempted" bson:"no_attempted"`
	TotalCorrect       int64     `json:"total_correct,omitempty" db:"total_correct" bson:"total_correct"`
	TotalWrong         int64     `json:"total_wrong,omitempty" db:"total_wrong" bson:"total_wrong"`
	TotalSubmissions   int64     `json:"total_submissions,omitempty" db:"total_submissions" bson:"total_submissions"`
}

// UserMetrics represents metrics for a user per language
type UserMetrics struct {
	UserID        string    `json:"user_id,omitempty" bson:"user_id" db:"user_id"`
	Username      string    `json:"username,omitempty" bson:"username" db:"username"`
	Level         int       `json:"level,omitempty" bson:"level" db:"level"` // 1 = Easy, 2 = Medium, 3 = Hard
	TotalScore    int       `json:"total_score,omitempty" bson:"total_score" db:"total_score"`
	Accuracy      float64   `json:"accuracy,omitempty" bson:"accuracy" db:"accuracy"`    // Percentage of correct answers
	SpeedAvg      float64   `json:"speed_avg,omitempty" bson:"speed_avg" db:"speed_avg"` // Average time (in seconds)
	PenaltyPoints int       `json:"penalty_points,omitempty" bson:"penalty_points" db:"penalty_points"`
	Rank          int       `json:"rank,omitempty" bson:"rank" db:"rank"`
	Language      string    `json:"language,omitempty" db:"language" bson:"language"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at" db:"created_at"`

	// New fields
	CorrectAnswers    int       `json:"correct_answers,omitempty" bson:"correct_answers" db:"correct_answers"`             // Total correct answers
	TotalQuestions    int       `json:"total_questions,omitempty" bson:"total_questions" db:"total_questions"`             // Total questions attempted
	TotalTime         float64   `json:"total_time,omitempty" bson:"total_time" db:"total_time"`                            // Total time taken in seconds
	TotalSubmissions  int       `json:"total_submissions,omitempty" bson:"total_submissions" db:"total_submissions"`       // Number of submissions
	CodeQualityScores []float64 `json:"code_quality_scores,omitempty" bson:"code_quality_scores" db:"code_quality_scores"` // Code quality scores for each submission
}

// GlobalUserPerformance represents overall metrics for a user
type GlobalUserPerformance struct {
	UserID        string    `json:"user_id,omitempty" bson:"user_id" db:"user_id"`
	Username      string    `json:"username,omitempty" bson:"username" db:"username"`
	Level         int       `json:"level,omitempty" bson:"level" db:"level"` // 1 = Easy, 2 = Medium, 3 = Hard
	TotalScore    int       `json:"total_score,omitempty" bson:"total_score" db:"total_score"`
	Accuracy      float64   `json:"accuracy,omitempty" bson:"accuracy" db:"accuracy"`    // Percentage of correct answers
	SpeedAvg      float64   `json:"speed_avg,omitempty" bson:"speed_avg" db:"speed_avg"` // Average time (in seconds)
	PenaltyPoints int       `json:"penalty_points,omitempty" bson:"penalty_points" db:"penalty_points"`
	Rank          int       `json:"rank,omitempty" bson:"rank" db:"rank"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at" db:"created_at"`

	// New fields
	CorrectAnswers    int       `json:"correct_answers,omitempty" bson:"correct_answers" db:"correct_answers"`
	TotalQuestions    int       `json:"total_questions,omitempty" bson:"total_questions" db:"total_questions"`
	TotalTime         float64   `json:"total_time,omitempty" bson:"total_time" db:"total_time"`
	TotalSubmissions  int       `json:"total_submissions,omitempty" bson:"total_submissions" db:"total_submissions"`
	CodeQualityScores []float64 `json:"code_quality_scores,omitempty" bson:"code_quality_scores" db:"code_quality_scores"`
}

// CalculateSpeedAvg calculates the average speed based on total time and submissions
func (um *UserMetrics) CalculateSpeedAvg() float64 {
	if um.TotalSubmissions == 0 {
		return 0
	}
	return um.TotalTime / float64(um.TotalSubmissions)
}

// CalculateAccuracy calculates the accuracy based on correct answers and total questions
func (um *UserMetrics) CalculateAccuracy() float64 {
	if um.TotalQuestions == 0 {
		return 0
	}
	return (float64(um.CorrectAnswers) / float64(um.TotalQuestions)) * 100
}

// CalculateCodeQuality calculates the average code quality score
func (um *UserMetrics) CalculateCodeQuality() float64 {
	if len(um.CodeQualityScores) == 0 {
		return 0
	}
	total := 0.0
	for _, score := range um.CodeQualityScores {
		total += score
	}
	return total / float64(len(um.CodeQualityScores))
}

type UserProgrammingAnalytics struct {
	UserID          string   `json:"user_id"`          // Unique identifier for the user
	QuestionID      string   `json:"question_id"`      // Unique identifier for the question
	QuestionText    string   `json:"question_text"`    // The text of the question
	Language        string   `json:"language"`         // Programming language used
	DifficultyLevel string   `json:"difficulty_level"` // Difficulty level of the question
	Attempts        int      `json:"attempts"`         // Number of attempts made for the question
	EventType       string   `json:"event_type"`       // Type of event (e.g., "submission", "attempt")
	Correct         bool     `json:"correct"`          // Whether the answer was correct or not
	SubmissionTime  string   `json:"submission_time"`  // Time of submission
	TimeTaken       string   `json:"time_taken"`       // Time taken to solve the question (in seconds)
	Score           float64  `json:"score"`            // Score for the submission
	CodeQuality     float64  `json:"code_quality"`     // Score for code quality
	Tags            []string `json:"tags"`             // Tags associated with the question
	SessionID       string   `json:"session_id"`       // Unique identifier for the user session
	CreatedAt       string   `json:"created_at"`       // Creation timestamp
	UpdatedAt       string   `json:"updated_at"`       // Last updated timestamp
	Timestamp       string   `json:"timestamp"`        // Timestamp of the event
}
