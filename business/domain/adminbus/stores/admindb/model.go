package admindb

import "time"

// this gets created when user logs in for the first time in the system and stays updated as user progresses in mongo and ES
type UserMetrics struct {
	UserID        string    `json:"user_id,omitempty" bson:"user_id" db:"user_id"`
	Username      string    `json:"username,omitempty" bson:"username" db:"username"`
	Level         int       `json:"level,omitempty" bson:"level" db:"level"` // 1 = Easy, 2 = Medium, 3 = Hard
	TotalScore    int       `json:"total_score,omitempty" bson:"total_score" db:"total_score"`
	Accuracy      float64   `json:"accuracy,omitempty" bson:"accuracy" db:"accuracy"`    // Percentage of correct answers
	SpeedAvg      float64   `json:"speed_avg,omitempty" bson:"speed_avg" db:"speed_avg"` // Average time (in seconds)
	PenaltyPoints int       `json:"penalty_points,omitempty" bson:"penalty_points" db:"penalty_points"`
	Rank          int       `json:"rank,omitempty" bson:"rank" db:"rank"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at" db:"created_at"`
}

//	type Challenge struct {
//		ChallengeID     string `json:"challenge_id,omitempty" bson:"challenge_id" db:"challenge_id"`
//		Title           string `json:"title,omitempty" bson:"title" db:"title"`
//		Difficulty      int    `json:"difficulty,omitempty" bson:"difficulty" db:"difficulty"`                   // 1 = Easy, 2 = Medium, 3 = Hard                                // Points for completing the challenge
//		ExecutionTime   int    `json:"execution_time,omitempty" bson:"execution_time" db:"execution_time"`       // Optimal time in milliseconds
//		OptimalSolution string `json:"optimal_solution,omitempty" bson:"optimal_solution" db:"optimal_solution"` // For quality evaluation
//	}

type User struct {
	UserID             string    `json:"user_id"`
	Username           string    `json:"username"`
	Rank               int       `json:"rank"`
	CreatedAt          time.Time `json:"created_at"`
	AttemptedQuestions []string  `json:"attempted_questions"` // List of question IDs user has faced
	SelectedLanguage   string    `json:"selected_language"`   // User's selected programming language
}

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
	PenaltyPoints int     `json:"penalty_points,omitempty" bson:"penalty_points" db:"penalty_points"`
	ExecutionTime int     `json:"execution_time,omitempty" db:"execution_time" bson:"execution_time"`
}

// Question represents a single challenge question
type Question struct {
	QuestionID  string   `json:"question_id,omitempty" db:"question_id" bson:"question_id"`
	Title       string   `json:"title,omitempty" db:"title" bson:"title"`
	Description string   `json:"description,omitempty" db:"description" bson:"description"`
	Logic       string   `json:"logic,omitempty" db:"logic" bson:"logic"`
	Difficulty  string   `json:"difficulty,omitempty" db:"difficulty" bson:"difficulty"`
	Tags        []string `json:"tags,omitempty" db:"tags" bson:"tags"`
	Language    string   `json:"language,omitempty" db:"language" bson:"language"` // Language the question is written in
	IsCompleted bool     `json:"is_completed,omitempty" db:"is_completed" bson:"is_completed"`
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

type Challenge struct {
	ChallengeID    string     `json:"challenge_id,omitempty" db:"challenge_id" bson:"challenge_id"`
	UserID         string     `json:"user_id,omitempty" bson:"user_id" db:"user_id"`
	Title          string     `json:"title,omitempty" db:"title" bson:"title"`
	Description    string     `json:"description,omitempty" db:"description" bson:"description"`
	Logic          string     `json:"logic,omitempty" db:"logic" bson:"logic"`
	Tags           []string   `json:"tags,omitempty" db:"tags" bson:"tags"`
	Difficulty     string     `json:"difficulty,omitempty" db:"difficulty" bson:"difficulty"`
	UserRank       int        `json:"user_rank,omitempty" db:"user_rank" bson:"user_rank"` // rank assigned to the user for this challenge
	Score          int        `json:"score,omitempty" bson:"score" db:"score"`
	Questions      []Question `json:"questions,omitempty" db:"questions" bson:"questions"` // List of 3 questions in this challenge
	CreatedAt      time.Time  `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
	CompletionDate time.Time  `json:"completion_date,omitempty" db:"completion_date" bson:"completion_date"`
	Language       string     `json:"language,omitempty" db:"language" bson:"language"` // Language the challenge is created for
	IsCompleted    bool       `json:"is_completed,omitempty" db:"is_completed" bson:"is_completed"`
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
