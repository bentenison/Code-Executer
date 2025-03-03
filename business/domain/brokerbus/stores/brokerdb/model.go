package brokerdb

import (
	"database/sql"
	"time"

	"github.com/bentenison/microservice/business/domain/brokerbus"
)

// type TestCase struct {
// 	Input          interface{} `bson:"input" json:"input"`                     // Input can be of any type
// 	ExpectedOutput interface{} `bson:"expected_output" json:"expected_output"` // Expected output can be of any type
// }

//	type Question struct {
//		QuestionId   primitive.ObjectID `bson:"_id" json:"questionId"`
//		Title        string             `bson:"title" json:"title"`                 // Title of the problem
//		Description  string             `bson:"description" json:"description"`     // Problem description
//		TemplateCode string             `bson:"template_code" json:"template_code"` // Code template for user logic
//		Language     string             `bson:"language" json:"language"`           // Programming language (e.g., Python)
//		LanguageCode string             `bson:"language_code" json:"language_code"` // Language code (e.g., "py")
//		TestCases    []TestCase         `bson:"test_cases" json:"test_cases"`       // List of test cases with dynamic types
//		Difficulty   string             `bson:"difficulty" json:"difficulty"`       // Difficulty level of the problem
//		Tags         []string           `bson:"tags" json:"tags"`                   // Tags related to the problem
//	}

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
}

type Input struct {
	Description string `json:"description" bson:"description"`
	Expected    string `json:"expected" bson:"expected"`
}

type Output struct {
	Description string `json:"description" bson:"description"`
}

type UserLogicTemplate struct {
	Description     string `json:"description,omitempty" bson:"description" db:"description"`
	Code            string `json:"code,omitempty" bson:"code" db:"code"`
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

// Submission struct for the 'submissions' table
type SubmissionDB struct {
	ID              sql.NullString `db:"id"`
	UserID          sql.NullString `db:"user_id"`
	LanguageID      sql.NullString `db:"language_id"`
	CodeSnippet     sql.NullString `db:"code_snippet"`
	SubmissionTime  sql.NullTime   `db:"submission_time"`
	ExecutionStatus sql.NullString `db:"execution_status"`
	ResultID        sql.NullString `db:"result_id"`
	IsPublic        sql.NullBool   `db:"is_public"`
	CreatedAt       sql.NullTime   `db:"created_at"`
	UpdatedAt       sql.NullTime   `db:"updated_at"`
	FileExtension   sql.NullString `json:"file_extension,omitempty" db:"file_extension"`
}

// PerformanceMetrics struct for the 'performance_metrics' table
type PerformanceMetricsDB struct {
	ID            sql.NullString `db:"id"`
	SubmissionID  sql.NullString `db:"submission_id"`
	ExecutionTime sql.NullTime   `db:"execution_time"`
	MemoryUsage   sql.NullInt64  `db:"memory_usage"`
	Status        sql.NullString `db:"status"`
	CreatedAt     sql.NullTime   `db:"created_at"`
	UpdatedAt     sql.NullTime   `db:"updated_at"`
}

// CodeExecutionStats struct for the 'code_execution_stats' table
type CodeExecutionStatsDB struct {
	ID            sql.NullString `db:"id"`
	UserID        sql.NullString `db:"user_id"`
	LanguageID    sql.NullString `db:"language_id"`
	ExecutionTime sql.NullTime   `db:"execution_time"`
	MemoryUsage   sql.NullInt64  `db:"memory_usage"`
	Status        sql.NullString `db:"status"`
	ErrorMessage  sql.NullString `db:"error_message"`
	CodeSnippet   sql.NullString `db:"code_snippet"`
	ContainerID   sql.NullString `db:"container_id"`
	CreatedAt     sql.NullTime   `db:"created_at"`
	UpdatedAt     sql.NullTime   `db:"updated_at"`
}

// User struct for the 'users' table
type UserDB struct {
	ID           sql.NullString `db:"id"`
	Username     sql.NullString `db:"username"`
	Email        sql.NullString `db:"email"`
	PasswordHash sql.NullString `db:"password_hash"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	Role         sql.NullString `db:"role"`
	CreatedAt    sql.NullTime   `db:"created_at"`
	UpdatedAt    sql.NullTime   `db:"updated_at"`
}

// Language struct for the 'languages' table
type LanguageDB struct {
	ID               sql.NullString `db:"id"`
	Code             sql.NullString `db:"code"`
	Name             sql.NullString `db:"name"`
	ContainerID      sql.NullString `db:"container_id"`
	ContainerName    sql.NullString `db:"container_name"`
	Version          sql.NullString `db:"version"`
	DocumentationURL sql.NullString `db:"documentation_url"`
	IsActive         sql.NullBool   `db:"is_active"`
	CreatedAt        sql.NullTime   `db:"created_at"`
	UpdatedAt        sql.NullTime   `db:"updated_at"`
	FileExtension    sql.NullString `db:"file_extension"`
	Description      sql.NullString `json:"description,omitempty" db:"description" bson:"description"` // Optional field
	Tags             []byte         `json:"tags,omitempty" db:"tags" bson:"tags"`                      // Optional field (array of strings)
	LogoURL          sql.NullString `json:"logo_url,omitempty" db:"logo_url" bson:"logo_url"`
}

type CodeSnippet struct {
	SnippetID string    `bson:"snippet_id,omitempty" json:"snippet_id,omitempty" db:"id"`
	Code      string    `bson:"code" json:"code,omitempty" db:"code"`
	Language  string    `bson:"language" json:"language,omitempty" db:"language"`
	CreatedBy string    `bson:"created_by" json:"created_by,omitempty" db:"created_by"`
	CreatedAt time.Time `bson:"createdAt" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updated_at,omitempty" db:"updated_at"`
}
type SubmissionQueryParams struct {
	ID         string    `db:"id"`
	QuestionID string    `db:"question_id"`
	UpdatedAt  time.Time `db:"updated_at"`
}
type ExampleData struct {
	ActorID    int       `json:"actor_id" bson:"actor_id"`
	FirstName  string    `json:"first_name" bson:"first_name"`
	LastName   string    `json:"last_name" bson:"last_name"`
	LastUpdate time.Time `json:"last_update" bson:"last_update"`
}

type Table struct {
	TableName        string                   `json:"table_name" bson:"table_name"`
	CreateTableQuery string                   `json:"create_table_query" bson:"create_table_query"`
	Columns          []string                 `json:"columns" bson:"columns"`
	ExampleData      []map[string]interface{} `json:"example_data" bson:"example_data"`
	RestoreQuery     string                   `json:"restore_query" bson:"restore_query"`
}

type DBQuestion struct {
	QueryType         string                   `json:"query_type" bson:"query_type"`
	QueryModifiesData bool                     `json:"query_modifies_data" bson:"query_modifies_data"`
	QuestionText      string                   `json:"question_text" bson:"question_text"`
	ExpectedResult    []map[string]interface{} `json:"expected_result" bson:"expected_result"`
	Hints             []string                 `json:"hints" bson:"hints"`
	ExpectedQuery     string                   `json:"expected_query" bson:"expected_query"`
}

type Validation struct {
	StrictOrdering   bool `json:"strict_ordering" bson:"strict_ordering"`
	IgnoreCase       bool `json:"ignore_case" bson:"ignore_case"`
	IgnoreWhitespace bool `json:"ignore_whitespace" bson:"ignore_whitespace"`
}

type SQLQuestion struct {
	ID          string     `json:"id" bson:"id"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Database    string     `json:"database" bson:"database"`
	Difficulty  string     `json:"difficulty" bson:"difficulty"`
	Tags        []string   `json:"tags" bson:"tags"`
	Tables      []Table    `json:"tables" bson:"tables"`
	DBQuestion  DBQuestion `json:"question" bson:"question"`
	Validation  Validation `json:"validation" bson:"validation"`
}

func toBusQuestion(q Question) brokerbus.Question {
	busQuestion := brokerbus.Question{}
	busQuestion.QuestionId = q.QuestionId
	busQuestion.Title = q.Title
	busQuestion.Description = q.Description
	busQuestion.Input = brokerbus.Input(q.Input)
	busQuestion.Output = brokerbus.Output(q.Output)
	busQuestion.UserLogicTemplate = brokerbus.UserLogicTemplate(q.UserLogicTemplate)
	busQuestion.TestcaseTemplate = brokerbus.TestcaseTemplate(q.TestcaseTemplate)
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
func addTestCases(cases []Testcase) []brokerbus.Testcase {
	out := []brokerbus.Testcase{}
	for _, v := range cases {
		out = append(out, brokerbus.Testcase(v))
	}
	return out
}
func toBusLanguage(lang *LanguageDB) *brokerbus.Language {
	var lg brokerbus.Language
	lg.ID = lang.ID.String
	lg.Code = lang.Code.String
	lg.Name = lang.Name.String
	lg.ContainerID = lang.ContainerID.String
	lg.ContainerName = lang.ContainerName.String
	lg.Version = lang.Version.String
	lg.DocumentationURL = lang.DocumentationURL.String
	lg.IsActive = lang.IsActive.Bool
	lg.UpdatedAt = lang.UpdatedAt.Time
	lg.CreatedAt = lang.UpdatedAt.Time
	lg.FileExtension = lang.FileExtension.String
	lg.Description = lang.Description.String
	lg.LogoURL = lang.LogoURL.String
	tgs := []string{}
	// err := json.Unmarshal(lang.Tags, &tgs)
	// if err != nil {
	// 	log.Println("error in unmarshaling tags", err)
	// }
	lg.Tags = tgs
	return &lg
}
func toBusLanguages(lang []LanguageDB) []*brokerbus.Language {
	var langs []*brokerbus.Language
	for _, v := range lang {
		lg := toBusLanguage(&v)
		langs = append(langs, lg)
	}
	return langs
}
func toBusQuestions(q []Question) []brokerbus.Question {
	busQuestions := []brokerbus.Question{}
	for _, v := range q {
		res := toBusQuestion(v)
		busQuestions = append(busQuestions, res)
	}
	return busQuestions
}
func toBusAnswer(a Answer) brokerbus.Answer {
	busAnswer := brokerbus.Answer{}
	busAnswer.ID = a.ID
	busAnswer.Logic = a.Logic
	busAnswer.TestCases = addTestCases(a.TestCases)
	return busAnswer
}
func toBusAnswers(a []Answer) []brokerbus.Answer {
	// busAnswer := brokerbus.Answer{}
	busAnswers := []brokerbus.Answer{}
	for _, v := range a {
		res := toBusAnswer(v)
		busAnswers = append(busAnswers, res)
	}
	return busAnswers
}

func toBusSQLQuestions(q []SQLQuestion) []brokerbus.SQLQuestion {
	var quests []brokerbus.SQLQuestion
	for _, v := range q {
		quest := toBusSQLQuestion(v)
		quests = append(quests, quest)
	}
	return quests
}
func toBusSQLQuestion(q SQLQuestion) brokerbus.SQLQuestion {
	busQuestion := brokerbus.SQLQuestion{}
	busQuestion.ID = q.ID
	busQuestion.Title = q.Title
	busQuestion.Description = q.Description
	busQuestion.Database = q.Database
	busQuestion.Difficulty = q.Difficulty
	busQuestion.Tags = q.Tags
	busQuestion.Tables = toBusTables(q.Tables)
	busQuestion.DBQuestion.QuestionText = q.DBQuestion.QuestionText
	busQuestion.DBQuestion.QueryType = q.DBQuestion.QueryType
	busQuestion.DBQuestion.ExpectedResult = q.DBQuestion.ExpectedResult
	busQuestion.DBQuestion.Hints = q.DBQuestion.Hints
	busQuestion.DBQuestion.ExpectedQuery = q.DBQuestion.ExpectedQuery
	busQuestion.Validation = toBusValidation(q.Validation)

	return busQuestion
}

func toBusTables(tables []Table) []brokerbus.Table {
	var busTables []brokerbus.Table
	for _, table := range tables {
		busTable := brokerbus.Table{}
		busTable.TableName = table.TableName
		busTable.CreateTableQuery = table.CreateTableQuery
		busTable.Columns = table.Columns
		busTable.ExampleData = table.ExampleData
		busTable.RestoreQuery = table.RestoreQuery
		busTables = append(busTables, busTable)
	}
	return busTables
}

func toBusExampleData(data []ExampleData) []brokerbus.ExampleData {
	var busData []brokerbus.ExampleData
	for _, row := range data {
		busRow := brokerbus.ExampleData{}
		busRow.ActorID = row.ActorID
		busRow.FirstName = row.FirstName
		busRow.LastName = row.LastName
		busRow.LastUpdate = row.LastUpdate // Convert time to string format
		busData = append(busData, busRow)
	}
	return busData
}

func toBusExpectedResult(result []ExampleData) []brokerbus.ExampleData {
	var busResult []brokerbus.ExampleData
	for _, row := range result {
		busRow := brokerbus.ExampleData{}
		busRow.ActorID = row.ActorID
		busRow.FirstName = row.FirstName
		busRow.LastName = row.LastName
		busRow.LastUpdate = row.LastUpdate // Convert time to string format
		busResult = append(busResult, busRow)
	}
	return busResult
}

func toBusValidation(v Validation) brokerbus.Validation {
	busValidation := brokerbus.Validation{}
	busValidation.StrictOrdering = v.StrictOrdering
	busValidation.IgnoreCase = v.IgnoreCase
	busValidation.IgnoreWhitespace = v.IgnoreWhitespace
	return busValidation
}
