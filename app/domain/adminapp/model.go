package adminapp

import (
	"time"

	"github.com/bentenison/microservice/business/domain/adminbus"
)

type User struct {
	UserID             string    `json:"user_id"`
	Username           string    `json:"username"`
	Rank               int       `json:"rank"`
	CreatedAt          time.Time `json:"created_at"`
	AttemptedQuestions []string  `json:"attempted_questions"` // List of question IDs user has faced
	SelectedLanguage   string    `json:"selected_language"`   // User's selected programming language
}

func ToBusUser(u User) adminbus.User {
	return adminbus.User(u)
}
