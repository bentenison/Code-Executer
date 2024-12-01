package creatorapp

import (
	"time"

	"github.com/bentenison/microservice/business/domain/creatorbus"
)

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
	Answer            Answer            `json:"answer,omitempty" bson:"answer" db:"answer"`
	IsQC              bool              `json:"is_qc,omitempty" bson:"is_qc" db:"is_qc"`
	TestCases         string            `json:"-"`
	ClassName         string            `json:"-"`
}
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
	FileExtension   string    `json:"file_extension,omitempty"`
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
type FilterPayload struct {
}
type QueryParams struct {
	Page             string
	Rows             string
	OrderBy          string
	ID               string
	UserID           string
	Lang             string
	Tags             string
	StartCreatedDate string
	EndCreatedDate   string
	IsQC             bool
}

func toBusQuestion(q Question) creatorbus.Question {
	busQuestion := creatorbus.Question{}
	busQuestion.QuestionId = q.QuestionId
	busQuestion.Title = q.Title
	busQuestion.Description = q.Description
	busQuestion.Input = creatorbus.Input(q.Input)
	busQuestion.Output = creatorbus.Output(q.Output)
	busQuestion.UserLogicTemplate = creatorbus.UserLogicTemplate(q.UserLogicTemplate)
	busQuestion.TestcaseTemplate = creatorbus.TestcaseTemplate(q.TestcaseTemplate)
	busQuestion.Difficulty = q.Difficulty
	busQuestion.TemplateCode = q.TemplateCode
	busQuestion.Language = q.Language
	busQuestion.LanguageCode = q.LanguageCode
	busQuestion.Tags = q.Tags
	busQuestion.Testcases = addTestCases(q.Testcases)
	busQuestion.ExecTemplate = q.ExecTemplate
	busQuestion.Answer = toBusAnswer(q.Answer)
	busQuestion.IsQC = q.IsQC
	return busQuestion
}
func addTestCases(cases []Testcase) []creatorbus.Testcase {
	out := []creatorbus.Testcase{}
	for _, v := range cases {
		out = append(out, creatorbus.Testcase(v))
	}
	return out
}

//	func toBusLanguage(lang *LanguageDB) *creatorbus.Language {
//		var lg creatorbus.Language
//		lg.ID = lang.ID.String
//		lg.Code = lang.Code.String
//		lg.Name = lang.Name.String
//		lg.ContainerID = lang.ContainerID.String
//		lg.ContainerName = lang.ContainerName.String
//		lg.Version = lang.Version.String
//		lg.DocumentationURL = lang.DocumentationURL.String
//		lg.IsActive = lang.IsActive.Bool
//		lg.UpdatedAt = lang.UpdatedAt.Time
//		lg.CreatedAt = lang.UpdatedAt.Time
//		return &lg
//	}
//
//	func toBusLanguages(lang []LanguageDB) []*brokerbus.Language {
//		var langs []*brokerbus.Language
//		for _, v := range lang {
//			lg := toBusLanguage(&v)
//			langs = append(langs, lg)
//		}
//		return langs
//	}
func toBusQuestions(q []Question) []creatorbus.Question {
	busQuestions := []creatorbus.Question{}
	for _, v := range q {
		res := toBusQuestion(v)
		busQuestions = append(busQuestions, res)
	}
	return busQuestions
}
func toBusAnswer(a Answer) creatorbus.Answer {
	busAnswer := creatorbus.Answer{}
	busAnswer.ID = a.ID
	busAnswer.Logic = a.Logic
	busAnswer.TestCases = addTestCases(a.TestCases)
	return busAnswer
}
func toBusAnswers(a []Answer) []creatorbus.Answer {
	// busAnswer := creatorbus.Answer{}
	busAnswers := []creatorbus.Answer{}
	for _, v := range a {
		res := toBusAnswer(v)
		busAnswers = append(busAnswers, res)
	}
	return busAnswers
}
