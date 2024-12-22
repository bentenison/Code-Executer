package creatordb

import (
	"time"

	"github.com/bentenison/microservice/business/domain/creatorbus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	QuestionId        string            `json:"id" bson:"id"`
	Title             string            `json:"title" bson:"title"`
	Description       string            `json:"description" bson:"description"`
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
	TestCases         string            `json:"tstcsc,omitempty" bson:"tstcsc" db:"tstcsc"`
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
type Concept struct {
	Label string `bson:"label" json:"label,omitempty" db:"label"`
	Value string `bson:"value" json:"value,omitempty" db:"value"`
}

// Define the LanguageConcept struct to represent each language with its concepts
type LanguageConcept struct {
	Language string    `bson:"language" json:"language,omitempty" db:"language"`
	Concepts []Concept `bson:"concepts" json:"concepts,omitempty" db:"concepts"`
}

type Answer struct {
	ID        string     `json:"id"`
	Logic     string     `json:"logic"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	TestCases []Testcase `json:"testcases"`
}
type QueryResult struct {
	Documents primitive.A
	Count     int32
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
func toBuslanguageConcept(c LanguageConcept) creatorbus.LanguageConcept {
	var concepts creatorbus.LanguageConcept
	concepts.Language = c.Language
	concepts.Concepts = createConcepts(c.Concepts)
	return concepts
}
func createConcepts(cnspts []Concept) []creatorbus.Concept {
	var outConcepts []creatorbus.Concept
	for _, v := range cnspts {
		outConcepts = append(outConcepts, creatorbus.Concept(v))
	}
	return outConcepts
}
func toBuslanguageConcepts(c []LanguageConcept) []creatorbus.LanguageConcept {
	var res []creatorbus.LanguageConcept
	for _, v := range c {
		out := toBuslanguageConcept(v)
		res = append(res, out)
	}
	return res
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
	busAnswer.CreatedAt = a.CreatedAt
	busAnswer.UpdatedAt = a.UpdatedAt
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

func toBusQueryResult(q QueryResult) creatorbus.QueryResult {
	// var busResults creatorbus.QueryResult
	return creatorbus.QueryResult(q)
}
