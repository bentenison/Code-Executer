package creatorbus

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

type Answer struct {
	ID        string     `json:"id"`
	Logic     string     `json:"logic"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	TestCases []Testcase `json:"testcases"`
}

type QueryFilter struct {
	ID               string
	UserID           string
	Lang             string
	Tags             string
	IsQc             bool
	StartCreatedDate *time.Time
	EndCreatedDate   *time.Time
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

type QueryResult struct {
	Documents primitive.A
	Count     int32
}

// ExamDetails represents the details of the exam.
type ExamDetails struct {
	ProgramID                   string                 `json:"programId,omitempty" bson:"programID,omitempty"`
	ECourseID                   string                 `json:"eCourseId,omitempty" bson:"eCourseID,omitempty"`
	ECoursePatternID            string                 `json:"eCoursePatternId,omitempty" bson:"eCoursePatternID,omitempty"`
	ClientExamId                string                 `json:"clientExamId,omitempty" bson:"clientExamId,omitempty"`
	ExamCode                    string                 `json:"examCode,omitempty" bson:"examCode,omitempty"`
	ExamName                    string                 `json:"examName,omitempty" bson:"examName,omitempty"`
	DefaultLanguageId           string                 `json:"defaultLanguageId,omitempty" bson:"defaultLanguageId,omitempty"`
	AllowedLanguages            []string               `json:"allowedLanguages,omitempty" bson:"allowedLanguages,omitempty"`
	Instructions                Instructions           `json:"instructions,omitempty" bson:"instructions,omitempty"`
	DurationInSec               int                    `json:"durationInSec,omitempty" bson:"durationInSec,omitempty"`
	TotalItems                  int                    `json:"totalItems,omitempty" bson:"totalItems,omitempty"`
	TotalMarks                  float64                `json:"totalMarks,omitempty" bson:"totalMarks,omitempty"`
	OptionRandamization         bool                   `json:"optionRandamization,omitempty" bson:"optionRandamization,omitempty"`
	ExamType                    string                 `json:"examType,omitempty" bson:"examType,omitempty"`
	AllowedAttempts             int                    `json:"allowedAttempts,omitempty" bson:"allowedAttempts,omitempty"`
	ExamScheduledFrom           *time.Time             `json:"examScheduledFrom,omitempty" bson:"examScheduledFrom,omitempty"`
	ExamScheduledTo             *time.Time             `json:"examScheduledTo,omitempty" bson:"examScheduledTo,omitempty"`
	ShowAnalysisReport          bool                   `json:"showAnalysisReport,omitempty" bson:"showAnalysisReport,omitempty"`
	Sections                    []ExamSection          `json:"sections,omitempty" bson:"sections,omitempty"`
	IsQuestionShufflingRequired bool                   `json:"isQuestionShufflingRequired,omitempty" bson:"isQuestionShufflingRequired,omitempty"`
	MinimumEndExamDurationInSec int                    `json:"minimumEndExamDurationInSec,omitempty" bson:"minimumEndExamDurationInSec,omitempty"`
	ExecuteInNewMode            bool                   `json:"executeInNewMode,omitempty" bson:"executeInNewMode,omitempty"`
	MascotImagePath             string                 `json:"mascotImagePath,omitempty" bson:"mascotImagePath,omitempty"`
	MinimumPassingMarks         float64                `json:"minimumPassingMarks,omitempty" bson:"minimumPassingMarks,omitempty"`
	ExamConf                    ExamRelated            `json:"examConf,omitempty" bson:"examConf,omitempty"`
	IsMockEnabled               bool                   `json:"isMockEnabled,omitempty" bson:"isMockEnabled,omitempty"`
	ExtraDetails                map[string]interface{} `json:"extraDetails,omitempty" bson:"extraDetails,omitempty"`
}

// ExamSection represents a section in the exam.
type ExamSection struct {
	SectionID        string          `json:"sectionId,omitempty" bson:"sectionID,omitempty"`
	SectionName      string          `json:"sectionName,omitempty" bson:"sectionName,omitempty"`
	Ordinality       int             `json:"ordinality,omitempty" bson:"ordinality,omitempty"`
	TotalItems       int             `json:"totalItems,omitempty" bson:"totalItems,omitempty"`
	TotalMarks       float64         `json:"totalMarks,omitempty" bson:"totalMarks,omitempty"`
	DurationInSec    int             `json:"durationInSec,omitempty" bson:"durationInSec,omitempty"`
	QuestionItemList []QuestionItem  `json:"questionItemList,omitempty" bson:"questionItemList,omitempty"`
	SectionConfigs   []SectionConfig `json:"sectionConfigs,omitempty" bson:"sectionConfigs,omitempty"`
}

// SectionConfig represents the configuration for a section.
type SectionConfig struct {
	Tag                string  `json:"tag,omitempty" bson:"tag,omitempty"`
	NoOfLowItems       int     `json:"noOfLowItems,omitempty" bson:"noOfLowItems,omitempty"`
	MarksForLowItem    float64 `json:"marksForLowItem,omitempty" bson:"marksForLowItem,omitempty"`
	NoOfMediumItems    int     `json:"noOfMediumItems,omitempty" bson:"noOfMediumItems,omitempty"`
	MarksForMediumItem float64 `json:"marksForMediumItem,omitempty" bson:"marksForMediumItem,omitempty"`
	NoOfHighItems      int     `json:"noOfHighItems,omitempty" bson:"noOfHighItems,omitempty"`
	MarksForHighItem   float64 `json:"marksForHighItem,omitempty" bson:"marksForHighItem,omitempty"`
}

// QuestionItem represents an individual programming question item in the exam.
type QuestionItem struct {
	Ordinality                int     `json:"ordinality,omitempty" bson:"ordinality,omitempty"`
	QuestionID                string  `json:"questionId,omitempty" bson:"questionID,omitempty"`
	QuestionType              string  `json:"questionType,omitempty" bson:"questionType,omitempty"`
	QuestionItemDataPath      string  `json:"questionItemDataPath,omitempty" bson:"questionItemDataPath,omitempty"`
	QuestionDifficultyLevel   string  `json:"questionDifficultyLevel,omitempty" bson:"questionDifficultyLevel,omitempty"`
	QuestionDisplayLanguageID string  `json:"questionDisplayLanguageId,omitempty" bson:"questionDisplayLanguageID,omitempty"`
	OutOfMarks                float64 `json:"outOfMarks,omitempty" bson:"outOfMarks,omitempty"`
	ExampleInput              string  `json:"exampleInput,omitempty" bson:"exampleInput,omitempty"`
	ExampleOutput             string  `json:"exampleOutput,omitempty" bson:"exampleOutput,omitempty"`
	Constraints               string  `json:"constraints,omitempty" bson:"constraints,omitempty"`
	LanguageSpecificHints     string  `json:"languageSpecificHints,omitempty" bson:"languageSpecificHints,omitempty"`
}

// Instructions represents exam instructions.
type Instructions struct {
	Text string `json:"text,omitempty" bson:"text,omitempty"`
}

// ExamRelated represents details related to the exam configuration.
type ExamRelated struct {
	// Define any specific fields for exam configuration here
}
