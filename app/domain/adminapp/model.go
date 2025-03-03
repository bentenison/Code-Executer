package adminapp

import (
	"time"

	"github.com/bentenison/microservice/business/domain/adminbus"
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

func ToBusUser(u User) adminbus.User {
	return adminbus.User(u)
}
