package adminbus

import "time"

const (
	BRONZE int64 = iota
	SILVER
	GOLD
	PLATINUM
)

type User struct {
	UserID             string    `json:"user_id,omitempty" db:"user_id" bson:"user_id"`
	Username           string    `json:"username,omitempty" db:"username" bson:"username"`
	Rank               int       `json:"rank,omitempty" db:"rank" bson:"rank"`
	CreatedAt          time.Time `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
	SelectedLanguage   string    `json:"selected_language,omitempty" db:"selected_language" bson:"selected_language"`       // User's selected programming language
	AttemptedQuestions []string  `json:"attempted_questions,omitempty" db:"attempted_questions" bson:"attempted_questions"` // List of question IDs user has faced
	NoAttempted        int64     `json:"no_attempted,omitempty" db:"no_attempted" bson:"no_attempted"`
	TotalCorrect       int64     `json:"total_correct,omitempty" db:"total_correct" bson:"total_correct"`
	TotalWrong         int64     `json:"total_wrong,omitempty" db:"total_wrong" bson:"total_wrong"`
	TotalSubmissions   int64     `json:"total_submissions,omitempty" db:"total_submissions" bson:"total_submissions"`
}

// User [performance per question] for ES
type UserPerformance struct {
	UserID        string  `json:"user_id,omitempty" db:"user_id" bson:"user_id"`
	Accuracy      float64 `json:"accuracy,omitempty" db:"accuracy" bson:"accuracy"`
	SpeedAvg      float64 `json:"speed_avg,omitempty" db:"speed_avg" bson:"speed_avg"`
	PenaltyPoints int     `json:"penalty_points,omitempty" db:"penalty_points" bson:"penalty_points"`
	QuestionId    string  `json:"question_id,omitempty" db:"question_id" bson:"question_id"`
	Language      string  `json:"language,omitempty" db:"language" bson:"language"`
	Rank          int     `json:"rank,omitempty" db:"rank" bson:"rank"`
	CreatedAt     string  `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
}

// challenge performance to store in ES
type ChallengeData struct {
	ChallengeID   string  `json:"challenge_id,omitempty" db:"challenge_id" bson:"challenge_id"`
	Difficulty    int     `json:"difficulty,omitempty" db:"difficulty" bson:"difficulty"`
	Score         int     `json:"score,omitempty" db:"score" bson:"score"`
	Accuracy      float64 `json:"accuracy,omitempty" bson:"accuracy" db:"accuracy"`    // Percentage of correct answers
	SpeedAvg      float64 `json:"speed_avg,omitempty" bson:"speed_avg" db:"speed_avg"` // Average time (in seconds)
	Language      string  `json:"language,omitempty" db:"language" bson:"language"`
	PenaltyPoints int     `json:"penalty_points,omitempty" bson:"penalty_points" db:"penalty_points"`
	ExecutionTime int     `json:"execution_time,omitempty" db:"execution_time" bson:"execution_time"`
}

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

// CalculateSpeedAvg calculates the average speed based on total time and submissions
func (um *GlobalUserPerformance) CalculateSpeedAvg() float64 {
	if um.TotalSubmissions == 0 {
		return 0
	}
	return um.TotalTime / float64(um.TotalSubmissions)
}

// CalculateAccuracy calculates the accuracy based on correct answers and total questions
func (um *GlobalUserPerformance) CalculateAccuracy() float64 {
	if um.TotalQuestions == 0 {
		return 0
	}
	return (float64(um.CorrectAnswers) / float64(um.TotalQuestions)) * 100
}

// CalculateCodeQuality calculates the average code quality score
func (um *GlobalUserPerformance) CalculateCodeQuality() float64 {
	if len(um.CodeQualityScores) == 0 {
		return 0
	}
	total := 0.0
	for _, score := range um.CodeQualityScores {
		total += score
	}
	return total / float64(len(um.CodeQualityScores))
}

type Question struct {
	QuestionID       string   `json:"question_id,omitempty" db:"question_id" bson:"question_id"`
	Title            string   `json:"title,omitempty" db:"title" bson:"title"`
	Description      string   `json:"description,omitempty" db:"description" bson:"description"`
	Logic            string   `json:"logic,omitempty" db:"logic" bson:"logic"`
	Difficulty       string   `json:"difficulty,omitempty" db:"difficulty" bson:"difficulty"`
	Tags             []string `json:"tags,omitempty" db:"tags" bson:"tags"`
	Language         string   `json:"language,omitempty" db:"language" bson:"language"` // Language the question is written in
	IsCompleted      bool     `json:"is_completed,omitempty" db:"is_completed" bson:"is_completed"`
	StartedAt        int64    `json:"started_at,omitempty" db:"started_at" bson:"started_at"`
	EndedAt          int64    `json:"ended_at,omitempty" db:"ended_at" bson:"ended_at"`
	TotalSubmissions int64    `json:"total_submissions,omitempty" db:"total_submissions" bson:"total_submissions"`
	Score            int64    `json:"score,omitempty" db:"score" bson:"score"`
	MaxScore         float64  `json:"max_score,omitempty" db:"max_score" bson:"max_score"`
}

type Challenge struct {
	ChallengeID string `json:"challenge_id,omitempty" db:"challenge_id" bson:"challenge_id"`
	UserID      string `json:"user_id,omitempty" bson:"user_id" db:"user_id"`
	// Title          string     `json:"title,omitempty" db:"title" bson:"title"`
	// Description    string     `json:"description,omitempty" db:"description" bson:"description"`
	// Logic          string     `json:"logic,omitempty" db:"logic" bson:"logic"`
	Tags           []string   `json:"tags,omitempty" db:"tags" bson:"tags"`
	Difficulty     string     `json:"difficulty,omitempty" db:"difficulty" bson:"difficulty"`
	UserRank       int        `json:"user_rank,omitempty" db:"user_rank" bson:"user_rank"` // rank assigned to the user for this challenge
	Score          int        `json:"score,omitempty" bson:"score" db:"score"`
	MaxScore       int        `json:"max_score,omitempty" db:"max_score" bson:"max_score"`
	Questions      []Question `json:"questions,omitempty" db:"questions" bson:"questions"` // List of 3 questions in this challenge
	CreatedAt      time.Time  `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
	CompletionDate time.Time  `json:"completion_date,omitempty" db:"completion_date" bson:"completion_date"`
	StartedAt      time.Time  `json:"started_at,omitempty" db:"started_at" bson:"started_at"`
	Language       string     `json:"language,omitempty" db:"language" bson:"language"` // Language the challenge is created for
	IsCompleted    bool       `json:"is_completed,omitempty" db:"is_completed" bson:"is_completed"`
}
type SubmissionStats struct {
	SubmissionID  string    `json:"submission_id,omitempty" bson:"submission_id" db:"submission_id"`
	UserID        string    `json:"user_id,omitempty" bson:"user_id" db:"user_id"`                      // Foreign key to User
	ChallengeID   string    `json:"challenge_id,omitempty" bson:"challenge_id" db:"challenge_id"`       // Foreign key to Challenge
	IsCorrect     bool      `json:"is_correct,omitempty" bson:"is_correct" db:"is_correct"`             // Whether the answer was correct
	Attempts      int       `json:"attempts,omitempty" bson:"attempts" db:"attempts"`                   // Number of attempts made
	TimeTaken     int       `json:"time_taken,omitempty" bson:"time_taken" db:"time_taken"`             // Time in seconds to complete
	CodeQuality   float64   `json:"code_quality,omitempty" bson:"code_quality" db:"code_quality"`       // 0 to 100 score for quality
	PenaltyPoints int       `json:"penalty_points,omitempty" bson:"penalty_points" db:"penalty_points"` // Penalties applied
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at" db:"created_at"`             // Timestamp when submission was created
}

type UserChallenge struct {
	UserID             string         `json:"user_id,omitempty" db:"user_id" bson:"user_id"`
	ChallengeID        string         `json:"challenge_id,omitempty" db:"challenge_id" bson:"challenge_id"`
	UserRank           int            `json:"user_rank,omitempty" db:"user_rank" bson:"user_rank"` // User's rank when they attempt this challenge
	Completed          bool           `json:"completed,omitempty" db:"completed" bson:"completed"`
	CreatedAt          time.Time      `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
	AttemptedQuestions []string       `json:"attempted_questions,omitempty" db:"attempted_questions" bson:"attempted_questions"` // List of question IDs user has faced
	QuestionsAttempted map[string]int `json:"questions_attempted,omitempty" db:"questions_attempted" bson:"questions_attempted"` // Tracks question_id and number of attempts
	SelectedLanguage   string         `json:"selected_language,omitempty" db:"selected_language" bson:"selected_language"`       // User's selected programming language
}

type CodingQuestion struct {
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

type Rank struct {
	ID                          string `json:"id,omitempty" db:"id" bson:"id"`
	Name                        string `json:"name,omitempty" db:"name" bson:"name"`
	IntegerRank                 int    `json:"integer_rank,omitempty" db:"integer_rank" bson:"integer_rank"`
	MinScore                    int    `json:"min_score,omitempty" db:"min_score" bson:"min_score"`
	Description                 string `json:"description,omitempty" db:"description" bson:"description"`
	TotalChallenges             int    `json:"total_challenges,omitempty" db:"total_challenges" bson:"total_challenges"`
	QuestionsPerChallenge       int    `json:"questions_per_challenge,omitempty" db:"questions_per_challenge" bson:"questions_per_challenge"`
	TotalQuestions              int    `json:"total_questions,omitempty" db:"total_questions" bson:"total_questions"`
	PointsRequired              int    `json:"points_required,omitempty" db:"points_required" bson:"points_required"`
	PointsPerChallenge          int    `json:"points_per_challenge,omitempty" db:"points_per_challenge" bson:"points_per_challenge"`
	ChallengesNeededForNextRank int    `json:"challenges_needed_for_next_rank,omitempty" db:"challenges_needed_for_next_rank" bson:"challenges_needed_for_next_rank"`
}

func (r *Rank) CalculateTotalQuestions() int {
	return r.TotalChallenges * r.QuestionsPerChallenge
}

// CalculateChallengesNeeded calculates how many challenges are needed to reach the next rank
func (r *Rank) CalculateChallengesNeeded() int {
	return r.PointsRequired / r.PointsPerChallenge
}

func (r *Rank) PointsPerQuestion() int {
	return r.PointsPerChallenge / r.QuestionsPerChallenge
}

type UpdatePayload struct {
	UserId      string  `json:"user_id,omitempty" db:"user_id" bson:"user_id"`
	QuestionId  string  `json:"question_id,omitempty" db:"question_id" bson:"question_id"`
	ChallengeId string  `json:"challenge_id,omitempty" db:"challenge_id" bson:"challenge_id"`
	CodeQuality float64 `json:"code_quality,omitempty" db:"code_quality" bson:"code_quality"`
	IsCorrect   bool    `json:"is_correct,omitempty" db:"is_correct" bson:"is_correct"`
	TimeTaken   int64   `json:"time_taken,omitempty" db:"time_taken" bson:"time_taken"`
	Language    string  `json:"language,omitempty" db:"language" bson:"language"`
	Score       int64   `json:"score,omitempty" db:"score" bson:"score"`
	IsChallenge bool    `json:"is_challenge,omitempty" db:"is_challenge" bson:"is_challenge"`
}
type UserAnalytics struct {
	UserID          string   `json:"user_id"`
	QuestionID      string   `json:"question_id"`
	Timestamp       string   `json:"timestamp"`
	EventType       string   `json:"event_type"`
	Correct         bool     `json:"correct"`
	TimeTaken       float64  `json:"time_taken"`
	Score           float64  `json:"score"`
	DifficultyLevel int      `json:"difficulty_level"`
	Language        string   `json:"language"`
	Tags            []string `json:"tags"`
	CodeQuality     float64  `json:"code_quality"`
	SessionID       string   `json:"session_id"`
}
