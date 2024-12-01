package brokerdb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	// objId, _ := primitive.ObjectIDFromHex(id)
	res := s.ds.MGO.Collection("qc_questions").FindOne(ctx, bson.M{"id": id})
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
            is_active,
			file_extension
        FROM languages WHERE is_active=true;
    `

	var languages []LanguageDB

	err := s.ds.SQL.Select(&languages, query)
	if err != nil {
		return nil, err
	}
	langs := toBusLanguages(languages)
	return langs, nil
}
func (s *Store) GetAllQuestionsDAO(ctx context.Context) ([]brokerbus.Question, error) {
	questions := []Question{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	cur, err := s.ds.MGO.Collection("qc_questions").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &questions)
	if err != nil {
		return nil, err
	}
	busQuest := toBusQuestions(questions)
	return busQuest, nil
}
func (s *Store) GetAnswerById(ctx context.Context, id string) (brokerbus.Answer, error) {
	answer := Answer{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res := s.ds.MGO.Collection("answers").FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return brokerbus.Answer{}, res.Err()
	}
	err := res.Decode(&answer)
	if err != nil {
		return brokerbus.Answer{}, err
	}
	busQuest := toBusAnswer(answer)
	return busQuest, nil
}
func (s *Store) GetAllAnswersDAO(ctx context.Context) ([]brokerbus.Answer, error) {
	answers := []Answer{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	cur, err := s.ds.MGO.Collection("questions").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &answers)
	if err != nil {
		return nil, err
	}
	busAnswers := toBusAnswers(answers)
	return busAnswers, nil
}
func (s *Store) UpdateQCQuestion(ctx context.Context, id string) (*mongo.UpdateResult, error) {
	// objId, _ := primitive.ObjectIDFromHex(id)
	update := bson.D{{"$set", bson.D{{"is_qc", true}}}}

	res, err := s.ds.MGO.Collection("qc_questions").UpdateOne(
		ctx,
		bson.M{"id": id},
		update,
	)

	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		// No document found with the given ID
		return nil, fmt.Errorf("document with ID %v not found", id)
	}

	if res.ModifiedCount == 0 {
		// Document found but no modification occurred, likely because "is_qc" was already true
		return nil, fmt.Errorf("no changes made, 'is_qc' may already be true")
	}

	return res, nil
}

func (s *Store) Get(ctx context.Context, key string, res any) error {
	data, err := s.ds.RDB.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(data), res)
	if err != nil {
		return nil
	}
	return nil
}
func (s *Store) GetQuestionTemplates(ctx context.Context) ([]brokerbus.Question, error) {
	questions := []Question{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res, err := s.ds.MGO.Collection("question_templates").Find(ctx, bson.M{})
	if err != nil {
		return []brokerbus.Question{}, res.Err()
	}
	err = res.All(ctx, &questions)
	if err != nil {
		return nil, err
	}
	busQuest := toBusQuestions(questions)
	return busQuest, nil
}

func (s *Store) Set(ctx context.Context, key string, val any, ttl time.Duration) (string, error) {
	var data string
	var err error
	marshalledIn, err := s.MarshalBinary(val)
	if err != nil {
		return "", err
	}
	if ttl != 0 {
		data, err = s.ds.RDB.Set(ctx, key, marshalledIn, ttl).Result()
		if err != nil {
			return "", err
		}

	} else {
		data, err = s.ds.RDB.Set(ctx, key, marshalledIn, 0).Result()
		if err != nil {
			return "", err
		}
	}
	s.logger.Errorc(ctx, "redis entry created", map[string]interface{}{
		"message": data,
	})
	return data, nil
}

func (s *Store) MarshalBinary(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
