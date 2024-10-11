package brokerbus

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
	Template    string    `json:"template,omitempty" bson:"template,omitempty"`
	Language    string    `json:"language,omitempty" bson:"language,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
}
