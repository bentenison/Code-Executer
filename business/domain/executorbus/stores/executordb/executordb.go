package executordb

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/executorbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/otel"
)

type Store struct {
	log *logger.CustomLogger
	db  mux.DataSource
}

func NewStore(log *logger.CustomLogger, ds mux.DataSource) *Store {
	return &Store{
		log: log,
		db:  ds,
	}
}

func (s *Store) GetLanguages(ctx context.Context) ([]*executorbus.Language, error) {
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

	err := s.db.SQL.Select(&languages, query)
	if err != nil {
		return nil, err
	}
	langs := toBusLanguages(languages)
	return langs, nil
}

func (s *Store) GetAllLangSpecs(ctx context.Context) ([]executorbus.LanguageSpecification, error) {
	// Query to fetch all languages
	rows, err := s.db.SQL.Query(`SELECT * FROM language_specifications`)
	if err != nil {
		return nil, fmt.Errorf("error querying the database: %v", err)
	}
	defer rows.Close()

	var languages []LanguageSpecification
	for rows.Next() {
		var language LanguageSpecification
		// var command []byte
		err := rows.Scan(&language.ID, &language.LanguageName, &language.FileExtension, &language.DockerImage, &language.Command, &language.CreatedAt, &language.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		languages = append(languages, language)
	}

	// Check for any errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}
	res := toBusSpecs(languages)
	return res, nil
}

// GetLanguageByID retrieves a language specification by its ID
func (s *Store) GetLanguageSpecsByID(ctx context.Context, id int) (executorbus.LanguageSpecification, error) {
	_, span := otel.AddSpan(ctx, "executorbus.GetLanguageSpecsByID")
	defer span.End()
	row := s.db.SQL.QueryRow(`SELECT * FROM language_specifications WHERE id = $1`, id)

	var language LanguageSpecification
	err := row.Scan(&language.ID, &language.LanguageName, &language.FileExtension, &language.DockerImage, &language.Command, &language.CreatedAt, &language.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return executorbus.LanguageSpecification{}, fmt.Errorf("no language found with ID %d", id)
		}
		return executorbus.LanguageSpecification{}, fmt.Errorf("error scanning row: %v", err)
	}
	res := toBusSpec(language)
	return res, nil
}

func (s *Store) Get(ctx context.Context, key string, res any) error {
	_, span := otel.AddSpan(ctx, "executorbus.redis.GET")
	defer span.End()
	data, err := s.db.RDB.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), res)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) Set(ctx context.Context, key string, val any, ttl time.Duration) (string, error) {
	_, span := otel.AddSpan(ctx, "executorbus.redis.SET")
	defer span.End()
	var data string
	var err error
	marshalledIn, err := s.MarshalBinary(val)
	if err != nil {
		return "", err
	}
	if ttl != 0 {
		data, err = s.db.RDB.Set(ctx, key, marshalledIn, ttl).Result()
		if err != nil {
			return "", err
		}

	} else {
		data, err = s.db.RDB.Set(ctx, key, marshalledIn, 0).Result()
		if err != nil {
			return "", err
		}
	}
	// s.log.Errorc(ctx, "redis entry created", map[string]interface{}{
	// 	"message": data,
	// })
	return data, nil
}

func (s *Store) MarshalBinary(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

//	question := Question{}
//	err := b.storer.Get(ctx, submission.QuestionId, &question)
