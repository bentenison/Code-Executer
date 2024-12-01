package executordb

import (
	"database/sql"
	"time"

	"github.com/bentenison/microservice/business/domain/executorbus"
)

type SubmissionDB struct {
	ID              sql.NullString `json:"id"`
	UserID          sql.NullString `json:"user_id"`
	LanguageID      sql.NullString `json:"language_id"`
	CodeSnippet     sql.NullString `json:"code_snippet"`
	SubmissionTime  sql.NullTime   `json:"submission_time"`
	ExecutionStatus sql.NullString `json:"execution_status"`
	ResultID        sql.NullString `json:"result_id"`
	IsPublic        sql.NullBool   `json:"is_public"`
	CreatedAt       sql.NullTime   `json:"created_at"`
	UpdatedAt       sql.NullTime   `json:"updated_at"`
}
type PerformanceMetricDB struct {
	ID            sql.NullInt64  `json:"id"`
	SubmissionID  sql.NullString `json:"submission_id"`
	ExecutionTime sql.NullTime   `json:"execution_time"`
	MemoryUsage   sql.NullInt64  `json:"memory_usage"`
	Status        sql.NullString `json:"status"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
}
type CodeExecutionStatDB struct {
	ID            sql.NullString `json:"id"`
	UserID        sql.NullString `json:"user_id"`
	LanguageID    sql.NullString `json:"language_id"`
	ExecutionTime sql.NullTime   `json:"execution_time"`
	MemoryUsage   sql.NullInt64  `json:"memory_usage"`
	Status        sql.NullString `json:"status"`
	ErrorMessage  sql.NullString `json:"error_message"`
	CodeSnippet   sql.NullString `json:"code_snippet"`
	ContainerID   sql.NullString `json:"container_id"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
}
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
}
type LanguageSpecification struct {
	ID            int       `json:"id,omitempty" db:"id"`
	LanguageName  string    `json:"language_name,omitempty" db:"language_name"`
	FileExtension string    `json:"file_extension,omitempty" db:"file_extension"`
	DockerImage   string    `json:"docker_image,omitempty" db:"docker_image"`
	Command       string    `json:"command,omitempty" db:"command"`
	CreatedAt     time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func toBusLanguage(lang *LanguageDB) *executorbus.Language {
	var lg executorbus.Language
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
	return &lg
}
func toBusLanguages(lang []LanguageDB) []*executorbus.Language {
	var langs []*executorbus.Language
	for _, v := range lang {
		lg := toBusLanguage(&v)
		langs = append(langs, lg)
	}
	return langs
}
func toBusSpec(specs LanguageSpecification) executorbus.LanguageSpecification {
	var spcs executorbus.LanguageSpecification
	spcs.ID = specs.ID
	spcs.LanguageName = specs.LanguageName
	spcs.Command = specs.Command
	spcs.FileExtension = specs.FileExtension
	spcs.UpdatedAt = specs.UpdatedAt
	return spcs
}
func toBusSpecs(specs []LanguageSpecification) []executorbus.LanguageSpecification {
	var res []executorbus.LanguageSpecification
	for _, v := range specs {
		spc := toBusSpec(v)
		res = append(res, spc)
	}
	return res
}
