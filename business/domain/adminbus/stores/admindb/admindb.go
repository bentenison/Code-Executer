package admindb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

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

func (s *Store) GetUserByUserId(userID, language string) (*adminbus.User, error) {
	// 51fc3552-45e0-4982-9adb-50d8cc46c46d
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
func (s *Store) InsertUser(user adminbus.User) (interface{}, error) {
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
func (s *Store) InsertUserPerformance(userPerformance adminbus.UserPerformance) (interface{}, error) {
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
func (s *Store) InsertChallengeData(challengeData adminbus.Challenge) (interface{}, error) {
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
func (s *Store) InsertUserMetrics(userMetrics adminbus.UserMetrics) (interface{}, error) {
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
func (s *Store) InsertGlobalUserPerformance(globalUserPerformance adminbus.GlobalUserPerformance) (interface{}, error) {
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
func (s *Store) InsertUserChallenge(userChallenge adminbus.UserChallenge) (interface{}, error) {
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
func (s *Store) GetGlobalUserPerformance(userID string) (*adminbus.GlobalUserPerformance, error) {
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
func (s *Store) UpdateUserMetrics(userID string, userMetrics adminbus.UserMetrics) (*adminbus.UserMetrics, error) {
	// Get the collection for user metrics
	userMetricsCollection := s.ds.MGO.Collection("user_metrics")

	// Update the user metrics by userID
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$set": userMetrics, // This will replace the document
	}
	_, err := userMetricsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update user metrics: %v", err)
	}

	// Return the updated user metrics
	return &userMetrics, nil
}
func (s *Store) UpdateGlobalUserPerformance(userID string, globalUserPerformance adminbus.GlobalUserPerformance) (*adminbus.GlobalUserPerformance, error) {
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
func (s *Store) UpdateChallengeData(challengeID string, challengeData adminbus.Challenge) (*adminbus.Challenge, error) {
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
func (s *Store) UpdateUserPerformance(userID string, userPerformance adminbus.UserPerformance) (*adminbus.UserPerformance, error) {
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
func (s *Store) GetUserMetrics(userID, language string) (*adminbus.UserMetrics, error) {
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
func (s *Store) GetAvailableQuestions(user *adminbus.User, allQuestions []adminbus.Question) ([]adminbus.Question, error) {
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
func (s *Store) CreateChallenge(user *adminbus.User, allQuestions []adminbus.Question) (*Challenge, error) {
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
func TrackUserChallenge(user *User, challenge *Challenge) error {
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
func (s *Store) FetchUserMetricsData() ([]UserMetrics, error) {
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
