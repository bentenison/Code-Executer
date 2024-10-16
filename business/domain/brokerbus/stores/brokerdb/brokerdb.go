package brokerdb

import (
	"context"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ds     mux.DataSource
	logger *logger.CustomLogger
}

func NewStore(ds mux.DataSource, logger *logger.CustomLogger) *Store {
	return &Store{
		ds:     ds,
		logger: logger,
	}
}

func (s *Store) GetQuestionTemplate(ctx context.Context, id string) (brokerbus.Question, error) {
	question := Question{}
	objId, _ := primitive.ObjectIDFromHex(id)
	res := s.ds.MGO.Collection("questions").FindOne(ctx, bson.M{"_id": objId})
	if res.Err() != nil {
		return brokerbus.Question{}, res.Err()
	}
	err := res.Decode(&question)
	if err != nil {
		return brokerbus.Question{}, err
	}
	busQuest := toBusQuestion(question)
	return busQuest, nil
}
func (s *Store) AddSubmission(ctx context.Context, submission *brokerbus.Submission) (string, error) {
	query := `
	INSERT INTO submissions (
		user_id, 
		language_id, 
		code_snippet, 
		submission_time, 
		execution_status, 
		result_id, 
		is_public, 
		created_at, 
		updated_at, 
		question_id
	) 
	VALUES (:user_id, :language_id, :code_snippet, :submission_time, :execution_status, 
			:result_id, :is_public, :created_at, :updated_at, :question_id) 
	RETURNING id
`

	// Execute the named query
	stmt, err := s.ds.SQL.PrepareNamed(query)
	if err != nil {
		return "", err
	}
	var id string
	err = stmt.Get(&id, submission)
	if err != nil {
		return "", err
	}

	return id, nil
}
func (s *Store) AddExecutionStats(ctx context.Context, newStat *brokerbus.CodeExecutionStats) (string, error) {
	var id string
	query := `
        INSERT INTO code_execution_stats (
            user_id, 
            language_id, 
            execution_time, 
            memory_usage, 
			total_memory,
			cpu_usage,
			memory_percentage,
            status, 
            error_message, 
            created_at, 
            updated_at, 
            code_snippet,
            container_id
        ) 
        VALUES (:user_id, :language_id, :execution_time, :memory_usage, :total_memory, :cpu_usage, :memory_percentage, :status, 
                :error_message, :created_at, :updated_at, :code_snippet, :container_id) 
        RETURNING id
    `
	stmt, err := s.ds.SQL.PrepareNamed(query)
	if err != nil {
		return "", err
	}
	// Execute the named query and get the generated ID
	err = stmt.Get(&id, newStat)
	if err != nil {
		// log.Fatalln(err)
		return "", err
	}
	return id, nil
}
func (s *Store) GetLanguages(ctx context.Context) ([]*brokerbus.Language, error) {
	query := `
        SELECT 
            id, 
            code, 
            name, 
            container_id, 
            container_name, 
            version, 
            created_at, 
            updated_at, 
            documentation_url, 
            is_active 
        FROM languages;
    `

	var languages []LanguageDB

	err := s.ds.SQL.Select(&languages, query)
	if err != nil {
		return nil, err
	}
	langs := toBusLanguages(languages)
	return langs, nil
}
