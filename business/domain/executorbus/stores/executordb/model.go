package executordb

import "database/sql"

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
