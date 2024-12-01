package creatordb

import (
	"context"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/creatorbus"
	"github.com/bentenison/microservice/business/sdk/page"
	"github.com/bentenison/microservice/foundation/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (s *Store) GetSingleQuestion(ctx context.Context, id string) (creatorbus.Question, error) {
	question := Question{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res := s.ds.MGO.Collection("questions").FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return creatorbus.Question{}, res.Err()
	}
	err := res.Decode(&question)
	if err != nil {
		return creatorbus.Question{}, err
	}
	busQuest := toBusQuestion(question)
	return busQuest, nil
}
func (s *Store) GetQuestionTemplates(ctx context.Context) ([]creatorbus.Question, error) {
	questions := []Question{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res, err := s.ds.MGO.Collection("question_templates").Find(ctx, bson.M{})
	if err != nil {
		return []creatorbus.Question{}, res.Err()
	}
	err = res.All(ctx, &questions)
	if err != nil {
		return nil, err
	}
	busQuest := toBusQuestions(questions)
	return busQuest, nil
}

func (s *Store) GetAllQuestionsDAO(ctx context.Context) ([]creatorbus.Question, error) {
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
func (s *Store) GetAnswerById(ctx context.Context, id string) (creatorbus.Answer, error) {
	answer := Answer{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res := s.ds.MGO.Collection("answers").FindOne(ctx, bson.M{"id": id})
	if res.Err() != nil {
		return creatorbus.Answer{}, res.Err()
	}
	err := res.Decode(&answer)
	if err != nil {
		return creatorbus.Answer{}, err
	}
	busQuest := toBusAnswer(answer)
	return busQuest, nil
}
func (s *Store) GetAllAnswersDAO(ctx context.Context) ([]creatorbus.Answer, error) {
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
func (s *Store) DeleteQuestion(ctx context.Context, id string) (int64, error) {
	// answers := []Answer{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res, err := s.ds.MGO.Collection("questions").DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
func (s *Store) DeleteAnswer(ctx context.Context, id string) (int64, error) {
	// answers := []Answer{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res, err := s.ds.MGO.Collection("answers").DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
func (s *Store) DeleteQuestions(ctx context.Context, ids []string) (int64, error) {
	// answers := []Answer{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res, err := s.ds.MGO.Collection("questions").DeleteMany(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		s.logger.Errorc(ctx, "error while getting data from DB", map[string]interface{}{
			"error": err.Error(),
		})
		return 0, err
	}
	return res.DeletedCount, nil
}
func (s *Store) DeleteAnswers(ctx context.Context, ids []string) (int64, error) {
	// answers := []Answer{}
	// objId, _ := primitive.ObjectIDFromHex(id)
	res, err := s.ds.MGO.Collection("answers").DeleteMany(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
func (s *Store) AddQuestions(ctx context.Context, qts []creatorbus.Question) ([]interface{}, error) {
	insertDocs := make([]interface{}, len(qts))
	for _, v := range qts {
		insertDocs = append(insertDocs, v)
	}
	res, err := s.ds.MGO.Collection("questions").InsertMany(ctx, insertDocs)
	if err != nil {
		return nil, err
	}
	return res.InsertedIDs, nil
}
func (s *Store) AddQCQuestions(ctx context.Context, qts []creatorbus.Question) ([]interface{}, error) {
	insertDocs := []interface{}{}
	for _, v := range qts {
		insertDocs = append(insertDocs, v)
	}
	res, err := s.ds.MGO.Collection("qc_questions").InsertMany(ctx, insertDocs)
	if err != nil {
		s.logger.Errorc(ctx, "error while getting data from DB", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	return res.InsertedIDs, nil
}
func (s *Store) AddAnswer(ctx context.Context, ans []creatorbus.Answer) ([]interface{}, error) {
	insertDocs := make([]interface{}, len(ans))
	res, err := s.ds.MGO.Collection("answers").InsertMany(ctx, insertDocs)
	if err != nil {
		return nil, err
	}
	return res.InsertedIDs, nil
}
func (s *Store) SearchQuestionByTag(ctx context.Context, tag string) ([]creatorbus.Question, error) {

	filter := bson.M{"tags": bson.M{"$in": []string{tag}}}

	cursor, err := s.ds.MGO.Collection("qc_questions").Find(context.Background(), filter)
	if err != nil {
		s.logger.Errorc(ctx, "error while getting data from DB", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []Question
	for cursor.Next(context.Background()) {
		var question Question
		if err := cursor.Decode(&question); err != nil {
			s.logger.Errorc(ctx, "error in decoding question", map[string]interface{}{
				"error": err.Error(),
			})
			continue
		}
		results = append(results, question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toBusQuestions(results), nil
}
func (s *Store) Query(ctx context.Context, filter creatorbus.QueryFilter, page page.Page) ([]creatorbus.Question, error) {
	query := s.applyFilter(filter)

	// Pagination logic
	skip := (page.PageNumber() - 1) * page.RowsPerPage()
	limit := page.RowsPerPage()

	cursor, err := s.ds.MGO.Collection("qc_questions").Find(context.Background(), query, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		s.logger.Errorc(ctx, "error while getting data from DB", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []Question
	for cursor.Next(context.Background()) {
		var question Question
		if err := cursor.Decode(&question); err != nil {
			s.logger.Errorc(ctx, "error in decoding question", map[string]interface{}{
				"error": err.Error(),
			})
			continue
		}
		results = append(results, question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toBusQuestions(results), nil
}

func (s *Store) SearchQuestionByLang(ctx context.Context, lang string) ([]creatorbus.Question, error) {

	filter := bson.M{"language": lang}

	cursor, err := s.ds.MGO.Collection("qc_questions").Find(context.Background(), filter)
	if err != nil {
		s.logger.Errorc(ctx, "error while getting data from DB", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []Question
	for cursor.Next(context.Background()) {
		var question Question
		if err := cursor.Decode(&question); err != nil {
			s.logger.Errorc(ctx, "error in decoding question", map[string]interface{}{
				"error": err.Error(),
			})
			continue
		}
		results = append(results, question)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return toBusQuestions(results), nil
}
func (s *Store) GetAllLanguageConcepts(ctx context.Context) ([]creatorbus.LanguageConcept, error) {
	var languageConcepts []LanguageConcept

	// Create an empty filter (this means no filter, so all documents will be retrieved)
	filter := bson.D{}

	// Find all language concepts in the collection
	cursor, err := s.ds.MGO.Collection("concepts").Find(context.Background(), filter)
	if err != nil {
		s.logger.Errorc(ctx, "error while getting data from DB", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document into the `languageConcepts` slice
	for cursor.Next(context.Background()) {
		var languageConcept LanguageConcept
		if err := cursor.Decode(&languageConcept); err != nil {
			s.logger.Errorc(ctx, "error while getting data from DB", map[string]interface{}{
				"error": err.Error(),
			})
			return nil, err
		}
		languageConcepts = append(languageConcepts, languageConcept)
	}

	if err := cursor.Err(); err != nil {
		s.logger.Errorc(ctx, "error while getting data from DB", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	res := toBuslanguageConcepts(languageConcepts)
	return res, nil
}
