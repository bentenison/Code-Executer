package brokerdb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/brokerbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/otel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/attribute"
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
	_, span := otel.AddSpan(ctx, "brokerbus.GetQuestionTemplate", attribute.String("db.FindOne", fmt.Sprintf("{id:%s}", id)),
		attribute.String("db.type", "mongo"))
	defer span.End()
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

// func (s *Store) AddSubmission(ctx context.Context, submission *brokerbus.Submission) (string, error) {
// 	_, span := otel.AddSpan(ctx, "brokerbus.AddSubmission", attribute.String("db.insert", `INSERT INTO submissions (
// 		user_id,
// 		language_id,
// 		code_snippet,
// 		submission_time,
// 		execution_status,
// 		result_id,
// 		is_public,
// 		created_at,
// 		updated_at,
// 		question_id
// 	)
// 	VALUES (:user_id, :language_id, :code_snippet, :submission_time, :execution_status,
// 			:result_id, :is_public, :created_at, :updated_at, :question_id)
// 	RETURNING id`),
// 		attribute.String("db.type", "pgsql"))
// 	defer span.End()
// 	query := `
// 	INSERT INTO submissions (
// 		user_id,
// 		language_id,
// 		code_snippet,
// 		submission_time,
// 		execution_status,
// 		result_id,
// 		is_public,
// 		created_at,
// 		updated_at,
// 		question_id
// 	)
// 	VALUES (:user_id, :language_id, :code_snippet, :submission_time, :execution_status,
// 			:result_id, :is_public, :created_at, :updated_at, :question_id)
// 	RETURNING id
// `

// 	// Execute the named query
// 	stmt, err := s.ds.SQL.PrepareNamed(query)
// 	if err != nil {
// 		return "", err
// 	}
// 	var id string
// 	err = stmt.Get(&id, submission)
// 	if err != nil {
// 		return "", err
// 	}

//		return id, nil
//	}
func (s *Store) AddSubmission(ctx context.Context, submission *brokerbus.Submission) (string, error) {
	_, span := otel.AddSpan(ctx, "brokerbus.AddSubmission",
		attribute.String("db.insert", `INSERT INTO submissions (
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
		RETURNING id`),
		attribute.String("db.type", "pgsql"))
	defer span.End()

	params := map[string]interface{}{
		"question_id": submission.QuestionId,
	}

	// First, check if submission for this question_id already exists
	var existingSubmission brokerbus.Submission
	// Use PrepareNamed for SELECT query
	stmt, err := s.ds.SQL.PrepareNamed(`SELECT id, run_count FROM submissions WHERE question_id = :question_id LIMIT 1`)
	if err != nil {
		return "", err
	}

	// Execute the query with parameters
	err = stmt.Get(&existingSubmission, params)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return "", err // Return error if it's something other than "no rows"
	}

	// If submission exists, increment run_count
	if existingSubmission.ID != "" {
		params["id"] = existingSubmission.ID
		params["updated_at"] = time.Now()

		// Prepare and execute the update query to increment run_count
		updateStmt, err := s.ds.SQL.PrepareNamed(`UPDATE submissions SET run_count = run_count + 1, updated_at = :updated_at WHERE id = :id`)
		if err != nil {
			return "", err
		}

		// Execute the update statement with the parameters
		_, err = updateStmt.Exec(params)
		if err != nil {
			return "", err
		}

		return existingSubmission.ID, nil
	}

	// If no existing submission, insert a new one
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
			question_id,
			run_count
		) 
		VALUES (:user_id, :language_id, :code_snippet, :submission_time, :execution_status, 
				:result_id, :is_public, :created_at, :updated_at, :question_id, 1) 
		RETURNING id
	`

	// Execute the insert query
	stmt, err = s.ds.SQL.PrepareNamed(query)
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
	_, span := otel.AddSpan(ctx, "brokerbus.AddExecutionStats", attribute.String("db.insert", `INSERT INTO code_execution_stats (
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
        )`),
		attribute.String("db.type", "pgsql"))
	defer span.End()
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
	_, span := otel.AddSpan(ctx, "brokerbus.GetLanguages", attribute.String("db.Select", `SELECT 
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
			file_extension,
			description,
			tags,
			logo_url
        FROM languages WHERE is_active=true;`), attribute.String("db.type", "pgsql"))
	defer span.End()
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
			file_extension,
			description,
			tags,
			logo_url
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
	_, span := otel.AddSpan(ctx, "brokerbus.GetAllQuestions", attribute.String("db.Find", "{}"), attribute.String("db.type", "mongo"), attribute.String("db.Collection", "qc_questions"))
	defer span.End()
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
	_, span := otel.AddSpan(ctx, "brokerbus.GetAnswerById")
	defer span.End()
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
	_, span := otel.AddSpan(ctx, "brokerbus.GetAllAnswers", attribute.String("db.Find", "{}"), attribute.String("db.collection", "questions"),
		attribute.String("db.type", "mongo"))
	defer span.End()
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
	_, span := otel.AddSpan(ctx, "brokerbus.UpdateQC", attribute.String("db.Update", `"$set", bson.D{{"is_qc", true}`), attribute.String("db.collection", "qc_questions"),
		attribute.String("db.type", "mongo"))
	defer span.End()
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
	_, span := otel.AddSpan(ctx, "brokerbus.redis.GET")
	defer span.End()
	data, err := s.ds.RDB.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), res)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetQuestionTemplates(ctx context.Context) ([]brokerbus.Question, error) {
	_, span := otel.AddSpan(ctx, "brokerbus.GetQuestionTemplates", attribute.String("db.FInd", `{}`), attribute.String("db.collection", "question_templates"),
		attribute.String("db.type", "mongo"))
	defer span.End()
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
	_, span := otel.AddSpan(ctx, "brokerbus.redis.SET")
	defer span.End()
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
func (s *Store) CreateSnippet(ctx context.Context, snippet *brokerbus.CodeSnippet) (*mongo.InsertOneResult, error) {
	// questions := []Question{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	_, span := otel.AddSpan(ctx, "brokerbus.AddSnippet")
	defer span.End()
	res, err := s.ds.MGO.Collection("snippets").InsertOne(ctx, snippet)
	if err != nil {
		s.logger.Errorc(ctx, "error adding snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return res, fmt.Errorf("error while adding code snippet to DB")
	}
	return res, nil
}
func (s *Store) GetSnippetById(ctx context.Context, id string) (*brokerbus.CodeSnippet, error) {
	_, span := otel.AddSpan(ctx, "brokerbus.GetSnippetByID", attribute.String("db.collection", "snippets"), attribute.String("db.FindOne", fmt.Sprintf(`{"snippet_id": %s`, id)), attribute.String("db.type", `mongo`))
	defer span.End()
	// Retrieve the snippet from the collection
	var snippet brokerbus.CodeSnippet
	err := s.ds.MGO.Collection("snippets").FindOne(ctx, bson.M{"snippet_id": id}).Decode(&snippet)
	if err == mongo.ErrNoDocuments {
		s.logger.Errorc(ctx, "error finding snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("snippet not found")
	}
	if err != nil {
		s.logger.Errorc(ctx, "error finding snippet", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("error while retrieving snippet: %v", err)
	}
	return &snippet, nil
}
func (s *Store) GetAllByUser(ctx context.Context, userId string) ([]brokerbus.CodeSnippet, error) {
	_, span := otel.AddSpan(ctx, "brokerbus.GetSnippetsOFUser", attribute.String("db.collection", "snippets"), attribute.String("db.FindOne", fmt.Sprintf(`{"created_by": %s`, userId)), attribute.String("db.type", `mongo`))
	defer span.End()
	var snippets []brokerbus.CodeSnippet
	cursor, err := s.ds.MGO.Collection("snippets").Find(ctx, bson.M{"created_by": userId})
	if err != nil {
		return nil, fmt.Errorf("error while fetching snippets for user %v: %v", userId, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var snippet brokerbus.CodeSnippet
		if err := cursor.Decode(&snippet); err != nil {
			return nil, fmt.Errorf("error decoding snippet: %v", err)
		}
		snippets = append(snippets, snippet)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return snippets, nil
}
func (s *Store) GetDBQuestByID(ctx context.Context, questionID string) (*brokerbus.SQLQuestion, error) {
	// Create a filter to search for the question by its QuestionId
	_, span := otel.AddSpan(ctx, "brokerbus.GetDBQuestByID", attribute.String("db.collection", "db_questions"), attribute.String("db.FindOne", fmt.Sprintf(`{"id": %s`, questionID)), attribute.String("db.type", `mongo`))
	defer span.End()
	filter := bson.D{{Key: "id", Value: questionID}}

	var question SQLQuestion
	err := s.ds.MGO.Collection("db_questions").FindOne(ctx, filter).Decode(&question)

	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("no question found with ID: %s", questionID)
	} else if err != nil {
		return nil, fmt.Errorf("error retrieving question from DB: %v", err)
	}

	// Return the found question
	res := toBusSQLQuestion(question)
	return &res, nil
}
func (s *Store) GetAllDBQuests(ctx context.Context, questionID string) ([]brokerbus.SQLQuestion, error) {
	// Create a filter to search for the question by its QuestionId
	_, span := otel.AddSpan(ctx, "brokerbus.GetAllDBQuests", attribute.String("db.collection", "db_questions"), attribute.String("db.FindOne", fmt.Sprintf(`{"id": %s`, questionID)), attribute.String("db.type", `mongo`))
	defer span.End()
	var quests []SQLQuestion
	cursor, err := s.ds.MGO.Collection("db_questions").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error while fetching all db questions : %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var quest SQLQuestion
		if err := cursor.Decode(&quest); err != nil {
			return nil, fmt.Errorf("error decoding question: %v", err)
		}
		quests = append(quests, quest)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	// Return the found question
	res := toBusSQLQuestions(quests)
	return res, nil
}
