package brokerdb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestCase struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output"`
}

type Question struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	CreatorID      primitive.ObjectID `json:"creator_id" bson:"creator_id"`
	Language       string             `json:"language"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	Difficulty     string             `json:"difficulty"`
	Tags           []string           `json:"tags"`
	TestTemplateID primitive.ObjectID `json:"test_template_id" bson:"test_template_id"`
	TestCases      []TestCase         `json:"test_cases"`
}

type Template struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Template    string             `json:"template"`
	Language    string             `json:"language"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Description string             `json:"description"`
}
