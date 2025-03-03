package admindb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/domain/adminbus"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/otel"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *Store) GetUserByUserId(ctx context.Context, userID, language string) (*adminbus.User, error) {
	// 51fc3552-45e0-4982-9adb-50d8cc46c46d
	_, span := otel.AddSpan(context.TODO(), "BROKER.GetUserByUserId", attribute.String("db.Type", "mongo"), attribute.String("db.FindeOne", fmt.Sprintf("{user_id:%s,selected_language:%s}", userID, language)))
	defer span.End()
	filter := bson.M{"user_id": userID, "selected_language": language}
	// Initialize an empty user struct to hold the result
	var user adminbus.User
	// Attempt to find the user in the collection
	err := s.ds.MGO.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		// If no user is found, return an empty user with default values
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		} else {
			return nil, fmt.Errorf("failed to fetch user: %v", err)
		}
	}

	// Return the user
	return &user, nil
}
func (s *Store) InsertUser(ctx context.Context, user adminbus.User) (interface{}, error) {
	otel.AddSpan(context.TODO(), "BROKER.GetUserByUserId", attribute.String("db.Type", "mongo"), attribute.String("db.InsertOne", fmt.Sprintf("{%s}", user)))
	// Get the collection from the database
	usersCollection := s.ds.MGO.Collection("users")
	// Attempt to find the user in the collection
	res, insertErr := usersCollection.InsertOne(context.Background(), user)
	if insertErr != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", insertErr)
	}
	// Return the user (whether it was fetched or newly created)
	return res.InsertedID, nil
}
func (s *Store) InsertUserPerformance(ctx context.Context, userPerformance adminbus.UserPerformance) (interface{}, error) {
	// Get the collection for user performances
	userPerformanceCollection := s.ds.MGO.Collection("user_metrics")

	// Insert the user performance record into the collection
	res, insertErr := userPerformanceCollection.InsertOne(context.Background(), userPerformance)
	if insertErr != nil {
		return nil, fmt.Errorf("failed to insert user performance: %v", insertErr)
	}

	// Return the inserted ID
	return res.InsertedID, nil
}
func (s *Store) InsertChallengeData(ctx context.Context, challengeData adminbus.Challenge) (interface{}, error) {
	// Get the collection for challenge data
	challengeDataCollection := s.ds.MGO.Collection("challenges")

	// Insert the challenge data record into the collection
	res, insertErr := challengeDataCollection.InsertOne(context.Background(), challengeData)
	if insertErr != nil {
		return nil, fmt.Errorf("failed to insert challenge data: %v", insertErr)
	}

	// Return the inserted ID
	return res.InsertedID, nil
}
func (s *Store) InsertUserMetrics(ctx context.Context, userMetrics adminbus.UserMetrics) (interface{}, error) {
	// Get the collection for user metrics
	userMetricsCollection := s.ds.MGO.Collection("user_metrics")

	// Insert the user metrics record into the collection
	res, insertErr := userMetricsCollection.InsertOne(context.Background(), userMetrics)
	if insertErr != nil {
		return nil, fmt.Errorf("failed to insert user metrics: %v", insertErr)
	}

	// Return the inserted ID
	return res.InsertedID, nil
}
func (s *Store) InsertGlobalUserPerformance(ctx context.Context, globalUserPerformance adminbus.GlobalUserPerformance) (interface{}, error) {
	// Get the collection for global user performance
	globalUserPerformanceCollection := s.ds.MGO.Collection("global_user_performance")

	// Insert the global user performance record into the collection
	res, insertErr := globalUserPerformanceCollection.InsertOne(context.Background(), globalUserPerformance)
	if insertErr != nil {
		return nil, fmt.Errorf("failed to insert global user performance: %v", insertErr)
	}

	// Return the inserted ID
	return res.InsertedID, nil
}
func (s *Store) InsertUserChallenge(ctx context.Context, userChallenge adminbus.UserChallenge) (interface{}, error) {
	// Get the collection for user challenges
	userChallengeCollection := s.ds.MGO.Collection("user_challenge")

	// Insert the user challenge record into the collection
	res, insertErr := userChallengeCollection.InsertOne(context.Background(), userChallenge)
	if insertErr != nil {
		return nil, fmt.Errorf("failed to insert user challenge: %v", insertErr)
	}

	// Return the inserted ID
	return res.InsertedID, nil
}
func (s *Store) GetGlobalUserPerformance(ctx context.Context, userID string) (*adminbus.GlobalUserPerformance, error) {
	// Get the collection for global user performance
	globalUserPerformanceCollection := s.ds.MGO.Collection("global_user_performance")

	// Attempt to find the global user performance by userID
	var globalUserPerformance adminbus.GlobalUserPerformance
	filter := bson.M{"user_id": userID}
	err := globalUserPerformanceCollection.FindOne(context.Background(), filter).Decode(&globalUserPerformance)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, fmt.Errorf("failed to fetch global user performance: %v", err)
	}

	// Return the found global user performance
	return &globalUserPerformance, nil
}
func (s *Store) GetUserChallengesByCompletionStatus(ctx context.Context, language, userid string, isCompleted bool) (adminbus.Challenge, error) {
	_, span := otel.AddSpan(ctx, "adminbus.GetUserChallengesByCompletionStatus",
		attribute.String("db.Find", fmt.Sprintf("{ \"is_completed\": %v }", isCompleted)),
		attribute.String("db.type", "mongo"),
		attribute.String("db.Collection", "challenges"))
	defer span.End()

	// Define a slice to hold the user challenges
	var userChallenges adminbus.Challenge

	// Filter on is_completed field
	filter := bson.M{"is_completed": isCompleted, "language": language, "user_id": userid}

	// Perform the query to find documents with the matching is_completed value
	err := s.ds.MGO.Collection("challenges").FindOne(ctx, filter).Decode(&userChallenges)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return userChallenges, mongo.ErrNoDocuments
		}
		return userChallenges, fmt.Errorf("failed to fetch user challenges: %v", err.Error())
	}

	return userChallenges, nil
}
func (s *Store) UpdateGlobalUserPerformance(ctx context.Context, userID string, globalUserPerformance adminbus.GlobalUserPerformance) (*adminbus.GlobalUserPerformance, error) {
	// Get the collection for global user performance
	globalUserPerformanceCollection := s.ds.MGO.Collection("global_user_performance")

	// Update the global user performance by userID
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": globalUserPerformance, // This will replace the document
	}
	_, err := globalUserPerformanceCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update global user performance: %v", err)
	}

	// Return the updated global user performance
	return &globalUserPerformance, nil
}
func (s *Store) UpdateChallengeData(ctx context.Context, challengeID string, challengeData adminbus.Challenge) (*adminbus.Challenge, error) {
	// Get the collection for challenge data
	challengeDataCollection := s.ds.MGO.Collection("challenges")

	// Update the challenge data by challengeID
	filter := bson.M{"challenge_id": challengeID}
	update := bson.M{
		"$set": challengeData, // This will replace the document
	}
	_, err := challengeDataCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update challenge data: %v", err)
	}

	// Return the updated challenge data
	return &challengeData, nil
}
func (s *Store) UpdateUserPerformance(ctx context.Context, userID string, userPerformance adminbus.UserPerformance) (*adminbus.UserPerformance, error) {
	// Get the collection for user performances
	userPerformanceCollection := s.ds.MGO.Collection("user_metrics")

	// Update the user performance by userID
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": userPerformance, // This will replace the document
	}
	_, err := userPerformanceCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update user performance: %v", err)
	}

	// Return the updated user performance
	return &userPerformance, nil
}
func (s *Store) GetUserMetrics(ctx context.Context, userID, language string) (*adminbus.UserMetrics, error) {
	// Get the collection for user metrics
	userMetricsCollection := s.ds.MGO.Collection("user_metrics")

	// Attempt to find the user metrics by userID
	var userMetrics adminbus.UserMetrics
	filter := bson.M{"user_id": userID}
	err := userMetricsCollection.FindOne(context.Background(), filter).Decode(&userMetrics)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, fmt.Errorf("failed to fetch user metrics: %v", err)
	}

	// Return the found user metrics
	return &userMetrics, nil
}
func (s *Store) GetRandomQuestionsByDifficultyAndLanguageDAO(ctx context.Context, difficulty string, language string) ([]adminbus.CodingQuestion, error) {
	_, span := otel.AddSpan(ctx, "adminbus.GetRandomQuestionsByDifficultyAndLanguage",
		attribute.String("db.Aggregate", "[{ $match: { difficulty: ?, language: ? } }, { $sample: { size: 3 } } ]"),
		attribute.String("db.type", "mongo"),
		attribute.String("db.Collection", "qc_questions"))
	defer span.End()

	// Define a slice to hold the randomly selected questions
	questions := []adminbus.CodingQuestion{}

	// Perform the aggregation with match for difficulty and language, then sample 3 questions
	cur, err := s.ds.MGO.Collection("qc_questions").Aggregate(ctx, mongo.Pipeline{
		// Match the difficulty and language
		{
			{"$match", bson.M{
				"difficulty": difficulty,
				"language":   language,
			}},
		},
		// Sample 3 random questions from the filtered results
		{
			{"$sample", bson.M{"size": 3}},
		},
	})
	if err != nil {
		return nil, err
	}

	// Convert the result into a slice of CodingQuestion
	err = cur.All(ctx, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}
func (s *Store) GetQuestionsByIDsDAO(ctx context.Context, ids []string) ([]adminbus.CodingQuestion, error) {
	_, span := otel.AddSpan(ctx, "adminbus.GetQuestionsByIDs",
		attribute.String("db.Aggregate", "[{ $match: { _id: { $in: ? } } }, { $sample: { size: 3 } } ]"),
		attribute.String("db.type", "mongo"),
		attribute.String("db.Collection", "qc_questions"))
	defer span.End()

	// Define a slice to hold the randomly selected questions
	questions := []adminbus.CodingQuestion{}

	// Perform the aggregation with match for the IDs, then sample 3 questions
	cur, err := s.ds.MGO.Collection("qc_questions").Aggregate(ctx, mongo.Pipeline{
		// Match the questions by the provided IDs
		{
			{"$match", bson.M{
				"id": bson.M{"$in": ids}, // Using $in to match any of the provided IDs
			}},
		},
	})
	if err != nil {
		return nil, err
	}

	// Convert the result into a slice of CodingQuestion
	err = cur.All(ctx, &questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

// TODO: get questions by dicfficulty level and languages
func (s *Store) GetAvailableQuestions(ctx context.Context, user *adminbus.User, allQuestions []adminbus.Question) ([]adminbus.Question, error) {
	// var availableQuestions []Question
	// var selectedQuestions []Question

	// // Iterate over all available questions and filter out those that the user has already faced or exceed the attempt limit
	// for _, question := range allQuestions {
	// 	// Ensure the question matches the user's selected language
	// 	if question.Language == user.SelectedLanguage && !contains(user.AttemptedQuestions, question.QuestionID) {
	// 		availableQuestions = append(availableQuestions, question)
	// 	}
	// }

	// // Select 3 random questions from availableQuestions
	// if len(availableQuestions) < 3 {
	// 	return nil, fmt.Errorf("not enough new questions available for the selected language")
	// }

	// // Select 3 random questions for the user
	// for i := 0; i < 3; i++ {
	// 	selectedQuestions = append(selectedQuestions, availableQuestions[i])
	// }

	return []adminbus.Question{}, nil
}

// CreateChallenge creates a new challenge for the user with 3 questions in their selected language
func (s *Store) CreateChallenge(ctx context.Context, user *adminbus.User, allQuestions []adminbus.Question) (*Challenge, error) {
	// Get 3 new questions for the user based on their selected language
	// selectedQuestions, err := s.GetAvailableQuestions(user, allQuestions)
	// if err != nil {
	// 	return nil, err
	// }

	// // Create the new challenge
	challenge := &Challenge{
		// ChallengeID: fmt.Sprintf("%s-challenge-%d", user.UserID, time.Now().Unix()),
		// UserID:      user.UserID,
		// Questions: selectedQuestions,
		// CreatedAt:   time.Now(),
		// Completed:   false,
		// Language:    user.SelectedLanguage, // Set language for the challenge
	}

	// Example: Insert challenge into the database
	// db.Insert("challenges", challenge)

	// Return the created challenge
	return challenge, nil
}

// TrackUserChallenge tracks user attempts and ensures they don't exceed 3 attempts per question
func TrackUserChallenge(ctx context.Context, user *User, challenge *Challenge) error {
	// Create a new UserChallenge entry
	// userChallenge := &UserChallenge{
	// 	UserID:        user.UserID,
	// 	ChallengeID:   challenge.ChallengeID,
	// 	QuestionsAttempted: make(map[string]int),
	// 	Completed:     false,
	// 	CreatedAt:     time.Now(),
	// }

	// Example: Insert into database
	// db.Insert("user_challenges", userChallenge)

	return nil
}
func contains(src []string, id string) bool {
	return false
}
func (s *Store) FetchUserMetricsData(ctx context.Context) ([]UserMetrics, error) {
	var result []UserMetrics
	searchRequest := esapi.SearchRequest{
		Index: []string{"user_performance"},
		Body:  strings.NewReader(`{"query": {"match_all": {}}}`),
	}

	resp, err := searchRequest.Do(context.Background(), s.ds.ES)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.String())
	}

	var esResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&esResponse); err != nil {
		return nil, err
	}

	// Extract hits from the response
	hits, ok := esResponse["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("error: could not extract hits from response")
	}
	for _, hit := range hits {
		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error: could not extract source from hit")
		}
		user := UserMetrics{
			UserID:        source["user_id"].(string),
			Accuracy:      source["accuracy"].(float64),
			SpeedAvg:      source["speed_avg"].(float64),
			PenaltyPoints: int(source["penalty_points"].(float64)),
			Rank:          int(source["rank"].(float64)),
			// CreatedAt:     source["created_at"].(string),
		}
		result = append(result, user)
	}

	return result, nil
}

func (s *Store) StorePerformanceDataES(ctx context.Context, performanceData []byte) error {
	req := esapi.IndexRequest{
		Index:   "user_performance",
		Body:    bytes.NewReader(performanceData),
		Refresh: "true", // To make the document searchable immediately
	}

	resp, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error storing performance data: %s", resp.String())
	}

	// fmt.Printf("Stored performance data for user %s\n", performanceData.UserID)
	return nil
}

func (s *Store) StoreChallengeDataES(ctx context.Context, challengeData []byte) error {
	req := esapi.IndexRequest{
		Index:   "challenge_data",
		Body:    bytes.NewReader(challengeData),
		Refresh: "true", // To make the document searchable immediately
	}

	resp, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error storing challenge data: %s", resp.String())
	}

	// fmt.Printf("Stored challenge data for challenge %s\n", challengeData.ChallengeID)
	return nil
}
func (s *Store) StoreCodeExecutionStatsES(ctx context.Context, codeStats []byte) error {
	req := esapi.IndexRequest{
		Index:   "code_execution_stats",
		Body:    bytes.NewReader(codeStats),
		Refresh: "true", // To make the document searchable immediately
	}

	resp, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("error storing challenge data: %s", resp.String())
	}

	// fmt.Printf("Stored challenge data for challenge %s\n", challengeData.ChallengeID)
	return nil
}

// performanceMappings := `
// 	{
// 		"mappings": {
// 			"properties": {
// 				"user_id": { "type": "keyword" },
// 				"accuracy": { "type": "float" },
// 				"speed_avg": { "type": "float" },
// 				"penalty_points": { "type": "integer" },
// 				"rank": { "type": "integer" },
// 				"created_at": { "type": "date" }
// 			}
// 		}
// 	}
// 	`

// 	// Define the mappings for challenge_data
// 	challengeMappings := `
// 	{
// 		"mappings": {
// 			"properties": {
// 				"challenge_id": { "type": "keyword" },
// 				"difficulty": { "type": "integer" },
// 				"score": { "type": "integer" },
// 				"execution_time": { "type": "integer" }
// 			}
// 		}
// 	}
// 	`

//	// Create indices with mappings
//	if err := createIndex(client, "user_performance", performanceMappings); err != nil {
//		log.Fatalf("Error creating user_performance index: %s", err)
//	}
//	if err := createIndex(client, "challenge_data", challengeMappings); err != nil {
//		log.Fatalf("Error creating challenge_data index: %s", err)
//	}
//
// Top N Users by Total Score
func (s *Store) GetTopUsersByTotalScore(ctx context.Context, limit int64) ([]UserMetrics, error) {
	var users []UserMetrics
	sort := bson.D{{"total_score", -1}} // Sort by total score in descending order
	cursor, err := s.ds.MGO.Collection("").Find(ctx, bson.M{}, options.Find().SetSort(sort).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user UserMetrics
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUsersByAccuracy fetches users whose accuracy is above a given threshold.
func (s *Store) GetUsersByAccuracy(ctx context.Context, accuracyThreshold float64) ([]UserMetrics, error) {
	var users []UserMetrics
	filter := bson.M{"accuracy": bson.M{"$gte": accuracyThreshold}}
	cursor, err := s.ds.MGO.Collection("").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user UserMetrics
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUsersWithLowSpeed fetches users who have a higher average completion time than a given threshold.
func (s *Store) GetUsersWithLowSpeed(ctx context.Context, speedThreshold float64) ([]UserMetrics, error) {
	var users []UserMetrics
	filter := bson.M{"speed_avg": bson.M{"$gte": speedThreshold}}
	cursor, err := s.ds.MGO.Collection("").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user UserMetrics
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserRankingHistory fetches ranking history for a user, sorted by creation date.
func (s *Store) GetUserRankingHistory(ctx context.Context, userID string) ([]UserMetrics, error) {
	var users []UserMetrics
	filter := bson.M{"user_id": userID}
	sort := bson.D{{"created_at", 1}} // Sort by creation date in ascending order
	cursor, err := s.ds.MGO.Collection("").Find(ctx, filter, options.Find().SetSort(sort))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user UserMetrics
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
func (s *Store) GetMostAttemptedChallenges(ctx context.Context, limit int64) ([]Challenge, error) {
	agg := bson.A{
		bson.M{"$group": bson.M{
			"_id":            "$challenge_id",
			"total_attempts": bson.M{"$sum": 1},
		}},
		bson.M{"$sort": bson.M{"total_attempts": -1}}, // Sort by total attempts
		bson.M{"$limit": limit},                       // Limit to top N
	}
	cursor, err := s.ds.MGO.Collection("").Aggregate(ctx, agg)
	if err != nil {
		return nil, err
	}

	var challenges []Challenge
	for cursor.Next(ctx) {
		var challenge Challenge
		if err := cursor.Decode(&challenge); err != nil {
			return nil, err
		}
		challenges = append(challenges, challenge)
	}

	return challenges, nil
}

// GetChallengesByDifficulty fetches challenges grouped by difficulty level.
func (s *Store) GetChallengesByDifficulty(ctx context.Context) ([]Challenge, error) {
	agg := bson.A{
		bson.M{"$group": bson.M{
			"_id":   "$difficulty",
			"count": bson.M{"$sum": 1},
		}},
		bson.M{"$sort": bson.M{"_id": 1}}, // Sort by difficulty
	}
	cursor, err := s.ds.MGO.Collection("").Aggregate(ctx, agg)
	if err != nil {
		return nil, err
	}

	var challenges []Challenge
	for cursor.Next(ctx) {
		var challenge Challenge
		if err := cursor.Decode(&challenge); err != nil {
			return nil, err
		}
		challenges = append(challenges, challenge)
	}

	return challenges, nil
}

func (s *Store) GetChallengesWithHighestCorrectRate(ctx context.Context, limit int64) ([]Challenge, error) {
	agg := bson.A{
		// Match submissions where is_correct is present
		bson.M{"$group": bson.M{
			"_id": "$challenge_id",
			"correct_count": bson.M{"$sum": bson.M{"$cond": bson.A{
				bson.M{"$eq": bson.A{"$is_correct", true}}, 1, 0,
			}}},
			"total_count": bson.M{"$sum": 1},
		}},
		bson.M{"$project": bson.M{
			"challenge_id": "$_id",
			"correct_rate": bson.M{"$divide": bson.A{"$correct_count", "$total_count"}},
		}},
		bson.M{"$sort": bson.M{"correct_rate": -1}}, // Sort by correct answer rate
		bson.M{"$limit": limit},                     // Limit to top N
	}

	cursor, err := s.ds.MGO.Collection("").Aggregate(ctx, agg)
	if err != nil {
		return nil, err
	}

	var challenges []Challenge
	for cursor.Next(ctx) {
		var challenge Challenge
		if err := cursor.Decode(&challenge); err != nil {
			return nil, err
		}
		challenges = append(challenges, challenge)
	}

	return challenges, nil
}

func (s *Store) GetUsersWithMostAttemptsForChallenge(ctx context.Context, challengeID string, limit int64) ([]SubmissionStats, error) {
	agg := bson.A{
		bson.M{"$match": bson.M{"challenge_id": challengeID}}, // Filter by challenge_id
		bson.M{"$group": bson.M{
			"_id":            "$user_id",
			"total_attempts": bson.M{"$sum": 1}, // Sum up attempts
		}},
		bson.M{"$sort": bson.M{"total_attempts": -1}}, // Sort by total attempts
		bson.M{"$limit": limit},                       // Limit to top N
	}

	cursor, err := s.ds.MGO.Collection("").Aggregate(ctx, agg)
	if err != nil {
		return nil, err
	}

	var submissions []SubmissionStats
	for cursor.Next(ctx) {
		var submission SubmissionStats
		if err := cursor.Decode(&submission); err != nil {
			return nil, err
		}
		submissions = append(submissions, submission)
	}

	return submissions, nil
}

func (s *Store) GetChallengesWithHighestCorrectRateES(ctx context.Context, limit int) ([]map[string]interface{}, error) {
	query := `{
		"size": 0,
		"aggs": {
			"challenges": {
				"terms": {
					"field": "challenge_id",
					"size": ` + string(limit) + `
				},
				"aggs": {
					"correct_answer_rate": {
						"avg": {
							"field": "accuracy"
						}
					},
					"total_submissions": {
						"value_count": {
							"field": "user_id"
						}
					},
					"correct_answers": {
						"sum": {
							"field": "correct_count"
						}
					}
				}
			}
		}
	}`

	req := esapi.SearchRequest{
		Body: bytes.NewReader([]byte(query)),
	}

	res, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Handle the response and return the data in a suitable format
	var result []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	return result, nil
}
func (s *Store) GetUsersWithMostAttemptsForChallengeES(ctx context.Context, challengeID string, limit int) ([]map[string]interface{}, error) {
	query := `{
		"query": {
			"match": {
				"challenge_id": "` + challengeID + `"
			}
		},
		"aggs": {
			"users_with_most_attempts": {
				"terms": {
					"field": "user_id",
					"size": ` + string(limit) + `
				},
				"aggs": {
					"total_attempts": {
						"value_count": {
							"field": "user_id"
						}
					}
				}
			}
		}
	}`

	req := esapi.SearchRequest{
		Index: []string{},
		Body:  strings.NewReader(query),
	}

	res, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Handle the response and return the data in a suitable format
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	results, ok := result["aggregations"].([]map[string]interface{})
	if !ok {

	}
	return results, nil
}

func (s *Store) GetCorrectIncorrectSubmissionsPerUser(ctx context.Context, userID string) (map[string]int, error) {
	query := `{
		"query": {
			"match": {
				"user_id": "` + userID + `"
			}
		},
		"aggs": {
			"correct_vs_incorrect": {
				"terms": {
					"field": "is_correct",
					"size": 2
				}
			}
		}
	}`

	req := esapi.SearchRequest{
		Body: strings.NewReader(query),
	}

	res, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Handle the response and return the data in a suitable format
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	results, ok := result["aggregations"].(map[string]int)
	if !ok {

	}
	return results, nil
}
func (s *Store) GetChallengeDifficultyDistribution(ctx context.Context) ([]map[string]interface{}, error) {
	query := `{
		"size": 0,
		"aggs": {
			"difficulty_distribution": {
				"terms": {
					"field": "difficulty"
				},
				"aggs": {
					"average_score": {
						"avg": {
							"field": "score"
						}
					}
				}
			}
		}
	}`

	req := esapi.SearchRequest{
		Body: strings.NewReader(query),
	}

	res, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	results, ok := result["aggregations"].([]map[string]interface{})
	if !ok {

	}
	return results, nil
}

func (s *Store) GetUserSegmentation(ctx context.Context) ([]map[string]interface{}, error) {
	query := `{
		"size": 0,
		"aggs": {
			"accuracy_bucket": {
				"range": {
					"field": "accuracy",
					"ranges": [
						{ "to": 50 },
						{ "from": 50, "to": 75 },
						{ "from": 75 }
					]
				},
				"aggs": {
					"speed_avg_bucket": {
						"range": {
							"field": "speed_avg",
							"ranges": [
								{ "to": 2 },
								{ "from": 2, "to": 5 },
								{ "from": 5 }
							]
						}
					}
				}
			}
		}
	}`

	req := esapi.SearchRequest{
		Body: strings.NewReader(query),
	}

	res, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	results, ok := result["aggregations"].([]map[string]interface{})
	if !ok {

	}
	return results, nil
}
func (s *Store) GetUserPerformanceOverTime(ctx context.Context, startDate, endDate string) ([]map[string]interface{}, error) {
	query := `{
		"query": {
			"range": {
				"created_at": {
					"gte": "` + startDate + `",
					"lte": "` + endDate + `"
				}
			}
		},
		"aggs": {
			"user_performance_over_time": {
				"date_histogram": {
					"field": "created_at",
					"interval": "month"
				},
				"aggs": {
					"avg_accuracy": {
						"avg": {
							"field": "accuracy"
						}
					},
					"avg_speed": {
						"avg": {
							"field": "speed_avg"
						}
					},
					"avg_penalty_points": {
						"avg": {
							"field": "penalty_points"
						}
					}
				}
			}
		}
	}`

	req := esapi.SearchRequest{
		Body: strings.NewReader(query),
	}

	res, err := req.Do(ctx, s.ds.ES)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	results, ok := result["aggregations"].([]map[string]interface{})
	if !ok {

	}
	return results, nil
}

func (s *Store) GetRankByID(ctx context.Context, rnk int64) (adminbus.Rank, error) {
	var rank adminbus.Rank
	err := s.ds.MGO.Collection("ranks").FindOne(ctx, bson.M{"integer_rank": rank}).Decode(&rank)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return rank, fmt.Errorf("rank not found")
		}
		return rank, err
	}
	return rank, nil
}
func (s *Store) GetAllRanks(ctx context.Context) ([]adminbus.Rank, error) {
	var ranks []adminbus.Rank
	cur, err := s.ds.MGO.Collection("ranks").Find(ctx, bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ranks, fmt.Errorf("rank not found")
		}
		return ranks, err
	}
	err = cur.All(ctx, &ranks)
	if err != nil {
		return nil, err
	}
	return ranks, nil
}

// :TODO create a getall ranks dao
// add data to redis
func (s *Store) Get(ctx context.Context, key string, res any) error {
	_, span := otel.AddSpan(ctx, "adminbus.redis.GET")
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

func (s *Store) Set(ctx context.Context, key string, val any, ttl time.Duration) (string, error) {
	_, span := otel.AddSpan(ctx, "adminbus.redis.SET")
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
	s.logger.Errorc(ctx, fmt.Sprintf("redis entry created for %s", key), map[string]interface{}{
		"message": data,
	})
	return data, nil
}

func (s *Store) MarshalBinary(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// UpdateUserStats(ctx context.Context, userID string, questionID string, correct bool) error
// UpdateUserMetrics(ctx context.Context, userID string, correct bool, timeTaken float64, codeQualityScore float64) error
// UpdateChallengeQuestion(ctx context.Context, challengeID string, questionID string, isCompleted bool, startedAt int64, endedAt int64) (*Question, error)
func (s *Store) UpdateUserStats(ctx context.Context, userID, questionID, language string, correct bool) error {
	// Prepare the update data
	update := bson.M{
		"$addToSet": bson.M{"attempted_questions": questionID}, // Add questionID to attempted_questions
		"$inc": bson.M{
			"no_attempted":      1, // Increment NoAttempted
			"total_submissions": 1, // Increment TotalSubmissions
		},
	}
	// Increment TotalCorrect or TotalWrong based on the answer
	if correct {
		update["$inc"].(bson.M)["total_correct"] = 1 // Increment TotalCorrect if the answer is correct
	} else {
		update["$inc"].(bson.M)["total_wrong"] = 1 // Increment TotalWrong if the answer is incorrect
	}
	_, span := otel.AddSpan(ctx, "ADMIN.UpdateUserStats", attribute.String("db.type", "mongo"), attribute.String("db.collection", "users"), attribute.String("db.UpdateOne", fmt.Sprintf("%v", update)))
	defer span.End()

	// Get the collection from the database
	usersCollection := s.ds.MGO.Collection("users")

	// Attempt to update the user in the collection
	_, err := usersCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID, "selected_language": language}, // Filter by userID
		update,
	)
	if err != nil {
		return fmt.Errorf("failed to update user stats: %v", err)
	}

	return nil
}
func (s *Store) UpdateUserMetrics(ctx context.Context, userID, language string, correct bool, timeTaken, codeQualityScore float64, score int64) error {
	// Get the collections from the database
	userMetricsCollection := s.ds.MGO.Collection("user_metrics")
	globalPerformanceCollection := s.ds.MGO.Collection("global_user_performance")

	// Fetch current user metrics
	var userMetrics adminbus.UserMetrics
	var globalPerf adminbus.GlobalUserPerformance
	if err := userMetricsCollection.FindOne(ctx, bson.M{"user_id": userID, "language": language}).Decode(&userMetrics); err != nil {
		return fmt.Errorf("failed to retrieve user metrics: %v", err)
	}

	// Prepare updates
	userMetrics.CorrectAnswers += 0 // This is for correct increment if correct is true
	if correct {
		userMetrics.CorrectAnswers++
	}

	userMetrics.TotalQuestions++
	userMetrics.TotalTime += timeTaken
	userMetrics.TotalSubmissions++
	userMetrics.TotalScore += int(score)
	userMetrics.CodeQualityScores = append(userMetrics.CodeQualityScores, codeQualityScore)
	acc := userMetrics.CalculateAccuracy()
	avg := userMetrics.CalculateSpeedAvg()
	// Prepare the update for UserMetrics
	userMetricsUpdate := bson.M{
		"$set": bson.M{
			"correct_answers":     userMetrics.CorrectAnswers,
			"total_questions":     userMetrics.TotalQuestions,
			"total_time":          userMetrics.TotalTime,
			"total_submissions":   userMetrics.TotalSubmissions,
			"total_score":         userMetrics.TotalScore,
			"accuracy":            acc,
			"speed_avg":           avg,
			"code_quality_scores": userMetrics.CodeQualityScores,
		},
	}

	// Update UserMetrics
	_, span := otel.AddSpan(ctx, "ADMIN.userMetricsUpdate", attribute.String("db.type", "mongo"), attribute.String("db.collection", "user_metrics"), attribute.String("db.UpdateOne", fmt.Sprintf("%v", userMetricsUpdate)))
	defer span.End()

	_, err := userMetricsCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID}, // Filter by userID
		userMetricsUpdate,
	)
	if err != nil {
		return fmt.Errorf("failed to update user metrics: %v", err)
	}
	if err := globalPerformanceCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&globalPerf); err != nil {
		return fmt.Errorf("failed to retrieve user metrics: %v", err)
	}
	// Prepare the update for GlobalUserPerformance (same update logic)

	globalPerf.CorrectAnswers += 0 // This is for correct increment if correct is true
	if correct {
		globalPerf.CorrectAnswers++
	}

	globalPerf.TotalQuestions++
	globalPerf.TotalTime += timeTaken
	globalPerf.TotalSubmissions++
	globalPerf.TotalScore += int(score)
	globalPerf.CodeQualityScores = append(globalPerf.CodeQualityScores, codeQualityScore)
	globalacc := globalPerf.CalculateAccuracy()
	globalavg := globalPerf.CalculateSpeedAvg()

	globalPerformanceUpdate := bson.M{
		"$set": bson.M{
			"correct_answers":     globalPerf.CorrectAnswers,
			"total_questions":     globalPerf.TotalQuestions,
			"total_time":          globalPerf.TotalTime,
			"total_submissions":   globalPerf.TotalSubmissions,
			"total_score":         globalPerf.TotalScore,
			"accuracy":            globalacc,
			"speed_avg":           globalavg,
			"code_quality_scores": globalPerf.CodeQualityScores,
		},
	}
	// Update GlobalUserPerformance
	_, err = globalPerformanceCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID}, // Filter by userID
		globalPerformanceUpdate,
	)
	if err != nil {
		return fmt.Errorf("failed to update global user performance: %v", err)
	}

	return nil
}

func (s *Store) UpdateChallengeQuestion(ctx context.Context, challengeID string, questionID string, isCompleted bool, endedAt int64, scored int64) (*adminbus.Question, error) {
	// Add OpenTelemetry span

	// Prepare the update data
	update := bson.M{
		"$set": bson.M{
			"questions.$.is_completed": isCompleted, // Update IsCompleted
			"questions.$.ended_at":     endedAt,     // Update EndedAt
			"questions.$.score":        scored,
		},
	}

	// Create the filter to find the specific challenge and question
	filter := bson.M{
		"challenge_id":          challengeID,
		"questions.question_id": questionID, // Match the question by ID
	}
	_, span := otel.AddSpan(ctx, "ADMIN.UpdateChallengeQuestion", attribute.String("db.type", "mongo"),
		attribute.String("db.collection", "challenges"),
		attribute.String("db.FindOneAndUpdate", fmt.Sprintf("%v", update)),
		attribute.String("filter", fmt.Sprintf("%v", filter)),
	)
	defer span.End()

	// Get the collection from the database
	challengesCollection := s.ds.MGO.Collection("challenges")

	// Use FindOneAndUpdate to get the updated question
	var updatedChallenge struct {
		Questions []adminbus.Question `bson:"questions"`
	}
	err := challengesCollection.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&updatedChallenge)

	if err != nil {
		// Log the error and return
		span.RecordError(err) // Record error in the span
		return nil, fmt.Errorf("failed to update question in challenge: %v", err)
	}

	// Find the updated question by ID in the updatedChallenge
	for _, question := range updatedChallenge.Questions {
		if question.QuestionID == questionID {
			return &question, nil // Return the updated question
		}
	}

	return nil, fmt.Errorf("question not found after update")
}

func (s *Store) GetQuestionByID(ctx context.Context, challengeID, questionID string) (*adminbus.Question, error) {
	// Add OpenTelemetry span
	_, span := otel.AddSpan(ctx, "GetQuestionByID",
		attribute.String("db.type", "mongo"),
		attribute.String("db.collection", "challenges"),
		attribute.String("challenge_id", challengeID),
		attribute.String("question_id", questionID),
	)
	defer span.End()

	// Prepare the aggregation pipeline
	pipeline := mongo.Pipeline{
		{
			{"$match", bson.M{"challenge_id": challengeID}},
		},
		{
			{"$unwind", "$questions"},
		},
		{
			{"$match", bson.M{"questions.question_id": questionID}},
		},
		{
			{"$replaceRoot", bson.M{"newRoot": "$questions"}},
		},
	}

	// Get the collection from the database
	challengesCollection := s.ds.MGO.Collection("challenges")

	// Execute the aggregation
	var questions []adminbus.Question
	cur, err := challengesCollection.Aggregate(ctx, pipeline)
	if err != nil {
		// Log the error and return
		span.RecordError(err) // Record error in the span
		return nil, fmt.Errorf("failed to find question in challenge: %v", err)
	}
	err = cur.All(ctx, &questions)
	if err != nil {
		// Log the error and return
		span.RecordError(err) // Record error in the span
		return nil, fmt.Errorf("failed to find question in challenge: %v", err)
	}
	var question adminbus.Question
	if len(questions) > 0 {
		question = questions[0]
	}
	if question.QuestionID == "" {
		return nil, fmt.Errorf("question not found in the specified challenge")
	}

	return &question, nil
}
