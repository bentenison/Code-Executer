package creatorapi

import "time"

type Question struct {
	QuestionId        string            `json:"id,omitempty" bson:"id" db:"question_id"`
	Title             string            `json:"title,omitempty" bson:"title" db:"title"`
	Description       string            `json:"description,omitempty" bson:"description" db:"description"`
	UserLogic         string            `json:"logic,omitempty" bson:"logic" db:"user_logic"`
	Input             Input             `json:"input,omitempty" bson:"input" db:"input"`
	Output            Output            `json:"output,omitempty" bson:"output" db:"output"`
	TemplateCode      string            `json:"template_code,omitempty" bson:"template_code" db:"template_code"`
	Language          string            `json:"language,omitempty" bson:"language" db:"language"`
	LanguageCode      string            `json:"language_code,omitempty" bson:"language_code" db:"language_code"`
	Difficulty        string            `json:"difficulty,omitempty" bson:"difficulty" db:"difficulty"`
	Tags              []string          `json:"tags,omitempty" bson:"tags" db:"tags"`
	UserLogicTemplate UserLogicTemplate `json:"user_logic_template,omitempty" bson:"user_logic_template" db:"user_logic_template"`
	TestcaseTemplate  TestcaseTemplate  `json:"testcase_template,omitempty" bson:"testcase_template" db:"testcase_template"`
	Testcases         []Testcase        `json:"testcases,omitempty" bson:"testcases" db:"testcases"`
	ExecTemplate      string            `json:"exec_template,omitempty" bson:"exec_template" db:"exec_template"`
	TestCases         string            `json:"test_cases,omitempty" db:"test_cases"`
	Answer            Answer            `json:"answer,omitempty" bson:"answer" db:"answer"`
	IsQC              bool              `json:"is_qc,omitempty" db:"is_qc"`
}

type Input struct {
	Description string `json:"description" bson:"description"`
	Expected    string `json:"expected" bson:"expected"`
}

type Output struct {
	Description string `json:"description" bson:"description"`
}

type UserLogicTemplate struct {
	Description string `json:"description" bson:"description"`
	Code        string `json:"code" bson:"code"`
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
