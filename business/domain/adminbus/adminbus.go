package adminbus

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/async/rabbit/rabbitconsumer"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gonum.org/v1/gonum/mat"
)

type Storer interface {
	GetUserByUserId(ctx context.Context, userID, language string) (*User, error)
	InsertUser(ctx context.Context, user User) (interface{}, error)
	GetRankByID(ctx context.Context, rnk int64) (Rank, error)
	InsertUserPerformance(ctx context.Context, userPerformance UserPerformance) (interface{}, error)
	InsertChallengeData(ctx context.Context, challengeData Challenge) (interface{}, error)
	InsertUserMetrics(ctx context.Context, userMetrics UserMetrics) (interface{}, error)
	InsertGlobalUserPerformance(ctx context.Context, globalUserPerformance GlobalUserPerformance) (interface{}, error)
	InsertUserChallenge(ctx context.Context, userChallenge UserChallenge) (interface{}, error)
	GetGlobalUserPerformance(ctx context.Context, userID string) (*GlobalUserPerformance, error)
	UpdateGlobalUserPerformance(ctx context.Context, userID string, globalUserPerformance GlobalUserPerformance) (*GlobalUserPerformance, error)
	UpdateChallengeData(ctx context.Context, challengeID string, challengeData Challenge) (*Challenge, error)
	UpdateUserPerformance(ctx context.Context, userID string, userPerformance UserPerformance) (*UserPerformance, error)
	GetUserMetrics(ctx context.Context, userID, language string) (*UserMetrics, error)
	GetRandomQuestionsByDifficultyAndLanguageDAO(ctx context.Context, difficulty string, language string) ([]CodingQuestion, error)
	GetUserChallengesByCompletionStatus(ctx context.Context, language string, userid string, isCompleted bool) (Challenge, error)
	GetQuestionsByIDsDAO(ctx context.Context, ids []string) ([]CodingQuestion, error)
	StoreCodeExecutionStatsES(ctx context.Context, codeStats []byte) error
	StoreChallengeDataES(ctx context.Context, challengeData []byte) error
	StorePerformanceDataES(ctx context.Context, performanceData []byte) error
	GetAllRanks(ctx context.Context) ([]Rank, error)
	Get(ctx context.Context, key string, res any) error
	Set(ctx context.Context, key string, val any, ttl time.Duration) (string, error)
	UpdateUserStats(ctx context.Context, userID string, questionID string, language string, correct bool) error
	// UpdateUserMetrics(ctx context.Context, userID string, correct bool, timeTaken float64, codeQualityScore float64) error
	UpdateUserMetrics(ctx context.Context, userID, language string, correct bool, timeTaken, codeQualityScore float64, score int64) error

	UpdateChallengeQuestion(ctx context.Context, challengeID string, questionID string, isCompleted bool, endedAt int64, scored int64) (*Question, error)
	GetQuestionByID(ctx context.Context, challengeID string, questionID string) (*Question, error)
	// GetAvailableQuestions(user *User, allQuestions []Question) ([]Question, error)
	// CreateChallenge(user *User, allQuestions []Question) (*Challenge, error)
}

type Business struct {
	log      *logger.CustomLogger
	delegate *delegate.Delegate
	storer   Storer
	// consumer *kafkaconsumer.Consumer
	rabbitConsumer *rabbitconsumer.Consumer
}

func NewBusiness(logger *logger.CustomLogger, delegate *delegate.Delegate, storer Storer, consumer *rabbitconsumer.Consumer) *Business {
	business := &Business{
		log:            logger,
		delegate:       delegate,
		storer:         storer,
		rabbitConsumer: consumer,
	}
	go consumer.ConsumeMessages()
	return business
}

func PredictRanking(userData []UserMetrics) {
	// Create a matrix for features (X)
	features := mat.NewDense(len(userData), 3, nil)
	// Create a vector for target (Y)
	target := mat.NewVecDense(len(userData), nil)

	// Fill the feature matrix and target vector
	for i, user := range userData {
		features.Set(i, 0, user.Accuracy)               // Accuracy
		features.Set(i, 1, user.SpeedAvg)               // SpeedAvg
		features.Set(i, 2, float64(user.PenaltyPoints)) // PenaltyPoints
		target.SetVec(i, float64(user.Rank))            // Rank (target)
	}

	// Use Gonum for linear regression using the normal equation: (X^T * X)^(-1) * X^T * Y
	// First, calculate the transpose of X
	var Xt mat.Dense
	Xt.CloneFrom(features.T())

	// Calculate (X^T * X)
	var XtX mat.Dense
	XtX.Mul(&Xt, features)

	// Calculate the inverse of (X^T * X)
	var XtX_inv mat.Dense
	if err := XtX.Inverse(&XtX_inv); err != nil {
		log.Fatalf("Error inverting (X^T * X): %v", err)
	}

	// Now calculate (X^T * X)^(-1) * X^T
	var XtX_inv_Xt mat.Dense
	XtX_inv_Xt.Mul(&XtX_inv, &Xt)

	// Calculate the coefficients (weights) for the model (W)
	var weights mat.Dense
	weights.Mul(&XtX_inv_Xt, target)

	// Print the weights (coefficients)
	fmt.Printf("Model Weights: \n%v\n", &weights)

	// Predict the rank for a new user based on the learned model
	newUser := []float64{85.0, 2.5, 0} // Example: accuracy=85, speed=2.5, penalties=0

	// Create a row vector from the new user's data
	newUserVec := mat.NewVecDense(3, newUser)

	// Perform the prediction by multiplying the new user's data with the weights
	// var prediction float64
	prediction := mat.Dot(newUserVec, weights.ColView(0))

	// Print the predicted rank
	fmt.Printf("Predicted Rank: %.2f\n", prediction)
}

// challenge creation
// 1) challenges are language based
// challenges are series of 3 questions of a perticular language
// inside challenge there is array of 3 questions with score points etc
// 2) user will get 3 attempts per question for a single day
// 3) we will record accuracy score avd speed
func (b *Business) AddPreRequisites(ctx context.Context, user User, language string) error {
	// get user if not found insert with default
	_, err := b.storer.GetGlobalUserPerformance(ctx, user.UserID)
	if err == mongo.ErrNoDocuments {
		var perf GlobalUserPerformance
		perf.CreatedAt = time.Now()
		perf.UserID = user.UserID
		perf.Username = user.Username
		// perf.Level =
		res, err := b.storer.InsertGlobalUserPerformance(ctx, perf)
		if err != nil {
			b.log.Errorc(ctx, "error in inserting user", map[string]interface{}{
				"error": err.Error(),
			})
			return err
		}
		_ = res
	}
	u, err := b.storer.GetUserByUserId(ctx, user.UserID, user.SelectedLanguage)
	if err == mongo.ErrNoDocuments {
		user.CreatedAt = time.Now()
		user.SelectedLanguage = language
		res, err := b.storer.InsertUser(ctx, user)
		if err != nil {
			b.log.Errorc(ctx, "error in inserting user", map[string]interface{}{
				"error": err.Error(),
			})
			return err
		}
		_ = res
	}
	//add user_metrics
	_, err = b.storer.GetUserMetrics(ctx, user.UserID, language)
	if err == mongo.ErrNoDocuments {
		var userMetrics UserMetrics
		userMetrics.UserID = user.UserID
		userMetrics.Language = user.SelectedLanguage
		userMetrics.CreatedAt = time.Now()
		userMetrics.Username = user.Username
		res, err := b.storer.InsertUserMetrics(ctx, userMetrics)
		if err != nil {
			b.log.Errorc(ctx, "error in inserting user", map[string]interface{}{
				"error": err.Error(),
			})
			return err
		}
		_ = res
	}
	b.log.Infoc(ctx, "user object", map[string]interface{}{
		"user": u,
	})
	//get available questions by difficulty and language
	//check if any challenge exists and if exists load that challange if not create new and insert
	// track user challenge progress
	return nil
}

func (b *Business) CreateChallengeService(ctx context.Context, user User, language string) (Challenge, error) {
	// b.storer.GetUserByUserId()
	var challange Challenge
	u, err := b.storer.GetUserByUserId(ctx, user.UserID, user.SelectedLanguage)
	if err != nil {
		b.log.Errorc(ctx, "error in getting user from db", map[string]interface{}{
			"error": err.Error(),
		})
		return challange, err
	}
	// get user attempted questions from users collection
	// get challengeData from challenges collections and track if user challenge exists which is not completed if it exists then get that challenge by challenge id aand load it
	challange, err = b.storer.GetUserChallengesByCompletionStatus(ctx, language, u.UserID, false)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// if not create challenges based on the previous questions
			// TODO get all ranks
			var ranks []Rank
			err := b.storer.Get(ctx, "ranks", &ranks)
			if err != nil {
				b.log.Errorc(ctx, "error while getting ranks redis!! going to set from mongo", map[string]interface{}{
					"error": err.Error(),
				})
				ranks, err = b.storer.GetAllRanks(ctx)
				if err != nil {
					b.log.Errorc(ctx, "error while getting ranks from DB! aborting...", map[string]interface{}{
						"error": err.Error(),
					})
					return Challenge{}, err
				}
				res, err := b.storer.Set(ctx, "ranks", &ranks, 0)
				if err != nil {
					b.log.Errorc(ctx, "error while setting ranks to Redis!", map[string]interface{}{
						"error": err.Error(),
						"res":   res,
					})
					return Challenge{}, err
				}
			}
			rank := getRankById(u.Rank, ranks)
			pts := float64(rank.PointsPerQuestion())
			challengeQuestions, err := b.storer.GetRandomQuestionsByDifficultyAndLanguageDAO(ctx, getdifficultyFromRank(u.Rank), strings.ToLower(u.SelectedLanguage))
			if err != nil {
				b.log.Errorc(ctx, "error in getting sample questions from db", map[string]interface{}{
					"error": err.Error(),
				})
				return challange, nil
			}
			//TODO get the ranks from DB and store to redis
			quests, tagMap := createQuestionFromCodingQuestion(challengeQuestions, pts)
			tags := extractTags(tagMap)
			challange.ChallengeID = uuid.NewString()
			challange.UserID = u.UserID
			challange.Tags = tags
			challange.Difficulty = getdifficultyFromRank(u.Rank)
			challange.CreatedAt = time.Now()
			challange.Language = language
			challange.Questions = quests
			challange.MaxScore = rank.PointsPerChallenge
			// 	ChallengeID    string     `json:"challenge_id,omitempty" db:"challenge_id" bson:"challenge_id"`
			// UserID         string     `json:"user_id,omitempty" bson:"user_id" db:"user_id"`
			// Tags           []string   `json:"tags,omitempty" db:"tags" bson:"tags"`
			// Difficulty     string     `json:"difficulty,omitempty" db:"difficulty" bson:"difficulty"`
			// UserRank       int        `json:"user_rank,omitempty" db:"user_rank" bson:"user_rank"` // rank assigned to the user for this challenge
			// Score          int        `json:"score,omitempty" bson:"score" db:"score"`
			// Questions      []Question `json:"questions,omitempty" db:"questions" bson:"questions"` // List of 3 questions in this challenge
			// CreatedAt      time.Time  `json:"created_at,omitempty" db:"created_at" bson:"created_at"`
			// CompletionDate time.Time  `json:"completion_date,omitempty" db:"completion_date" bson:"completion_date"`
			// Language       string     `json:"language,omitempty" db:"language" bson:"language"` // Language the challenge is created for
			// IsCompleted    bool       `json:"is_completed,omitempty" db:"is_completed" bson:"is_completed"`
			_, err = b.storer.InsertChallengeData(ctx, challange)
			if err != nil {
				b.log.Errorc(ctx, "error in adding user challenges", map[string]interface{}{
					"error": err.Error(),
				})
				return challange, err
			}
			return challange, nil
		}
		b.log.Errorc(ctx, "error in getting user challenges", map[string]interface{}{
			"error": err.Error(),
		})
		return challange, err
	}

	// if not create challenges based on the previous questions
	// return the created challenge
	return challange, nil
}

func getdifficultyFromRank(rank int) string {
	switch rank {
	case 0:
		return "easy"
	case 1:
		return "medium"
	case 2:
		return "high"
	default:
		return "hign"
	}
}
func createQuestionFromCodingQuestion(codingQuestions []CodingQuestion, score float64) ([]Question, map[string]bool) {
	// Create and populate the Question struct from the CodingQuestion struct
	questions := []Question{}
	tags := make(map[string]bool)
	for _, v := range codingQuestions {
		questions = append(questions, Question{
			QuestionID:  v.QuestionId,
			Title:       v.Title,
			Description: v.Description,
			Logic:       v.UserLogic,
			Difficulty:  v.Difficulty,
			Tags:        v.Tags,
			Language:    v.Language,
			StartedAt:   time.Now().UnixNano(),
			EndedAt:     time.Now().UnixNano(),
			MaxScore:    score,
			IsCompleted: false, // Assuming that the question is not completed when created, can be modified based on your requirements
		})
		for _, tg := range v.Tags {
			tags[tg] = true
		}
	}
	return questions, tags
}
func extractTags(tagMap map[string]bool) []string {
	tags := []string{}
	for k, _ := range tagMap {
		tags = append(tags, k)
	}
	return tags
}
func (b *Business) FetchQuestionsByIds(ctx context.Context, ids []string) ([]CodingQuestion, error) {
	res, err := b.storer.GetQuestionsByIDsDAO(ctx, ids)
	if err != nil {
		b.log.Errorc(ctx, "error in getting questions", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	return res, nil
}

func getRankById(rankId int, ranks []Rank) Rank {
	if len(ranks) > 0 {
		for _, v := range ranks {
			if v.IntegerRank == rankId {
				return v
			}
		}
	}
	return Rank{}
}

func (b *Business) MarkQuestionCompletion(ctx context.Context, payload UpdatePayload) error {
	// get the challange id questionId and isCorrect from the request
	//TODO add result modify the performnace, score, aacuracy mark question aatempted
	// TODO correct answer modify
	u, err := b.storer.GetUserByUserId(ctx, payload.UserId, payload.Language)
	if err != nil {
		b.log.Errorc(ctx, "error in getting user from db", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	var ranks []Rank
	err = b.storer.Get(ctx, "ranks", &ranks)
	if err != nil {
		b.log.Errorc(ctx, "error while getting ranks redis!! going to set from mongo", map[string]interface{}{
			"error": err.Error(),
		})
		ranks, err = b.storer.GetAllRanks(ctx)
		if err != nil {
			b.log.Errorc(ctx, "error while getting ranks from DB! aborting...", map[string]interface{}{
				"error": err.Error(),
			})
			return err
		}
		res, err := b.storer.Set(ctx, "ranks", &ranks, 0)
		if err != nil {
			b.log.Errorc(ctx, "error while setting ranks to Redis!", map[string]interface{}{
				"error": err.Error(),
				"res":   res,
			})
			return err
		}
	}
	rank := getRankById(u.Rank, ranks)
	pts := float64(rank.PointsPerQuestion())
	payload.Score = int64(pts)
	ended_at := time.Now().UnixNano()
	_, err = b.storer.UpdateChallengeQuestion(ctx, payload.ChallengeId, payload.QuestionId, true, ended_at, payload.Score)
	if err != nil {
		b.log.Errorc(ctx, "error in UpdateChallengeQuestion:", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (b *Business) UpdateUserMetrics(ctx context.Context, payload UpdatePayload) error {
	// type GlobalUserPerformance struct {
	// 	TotalScore    int       `json:"total_score,omitempty" bson:"total_score" db:"total_score"`
	// 	Accuracy      float64   `json:"accuracy,omitempty" bson:"accuracy" db:"accuracy"`    // Percentage of correct answers
	// 	SpeedAvg      float64   `json:"speed_avg,omitempty" bson:"speed_avg" db:"speed_avg"` // Average time (in seconds)
	// 	PenaltyPoints int       `json:"penalty_points,omitempty" bson:"penalty_points" db:"penalty_points"`
	// 	// New fields
	// 	CorrectAnswers    int       `json:"correct_answers,omitempty" bson:"correct_answers" db:"correct_answers"`
	// 	TotalQuestions    int       `json:"total_questions,omitempty" bson:"total_questions" db:"total_questions"`
	// 	TotalTime         float64   `json:"total_time,omitempty" bson:"total_time" db:"total_time"`
	// 	TotalSubmissions  int       `json:"total_submissions,omitempty" bson:"total_submissions" db:"total_submissions"`
	// 	CodeQualityScores []float64 `json:"code_quality_scores,omitempty" bson:"code_quality_scores" db:"code_quality_scores"`
	// }
	// type UserMetrics struct {
	// 	TotalScore    int       `json:"total_score,omitempty" bson:"total_score" db:"total_score"`
	// 	Accuracy      float64   `json:"accuracy,omitempty" bson:"accuracy" db:"accuracy"`    // Percentage of correct answers
	// 	SpeedAvg      float64   `json:"speed_avg,omitempty" bson:"speed_avg" db:"speed_avg"` // Average time (in seconds)
	// 	PenaltyPoints int       `json:"penalty_points,omitempty" bson:"penalty_points" db:"penalty_points"`
	// 	Rank          int       `json:"rank,omitempty" bson:"rank" db:"rank"`
	// 	// New fields
	// 	CorrectAnswers    int       `json:"correct_answers,omitempty" bson:"correct_answers" db:"correct_answers"`             // Total correct answers
	// 	TotalQuestions    int       `json:"total_questions,omitempty" bson:"total_questions" db:"total_questions"`             // Total questions attempted
	// 	TotalTime         float64   `json:"total_time,omitempty" bson:"total_time" db:"total_time"`                            // Total time taken in seconds
	// 	TotalSubmissions  int       `json:"total_submissions,omitempty" bson:"total_submissions" db:"total_submissions"`       // Number of submissions
	// 	CodeQualityScores []float64 `json:"code_quality_scores,omitempty" bson:"code_quality_scores" db:"code_quality_scores"` // Code quality scores for each submission
	// }
	u, err := b.storer.GetUserByUserId(ctx, payload.UserId, payload.Language)
	if err != nil {
		b.log.Errorc(ctx, "error in getting user from db", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	var ranks []Rank
	err = b.storer.Get(ctx, "ranks", &ranks)
	if err != nil {
		b.log.Errorc(ctx, "error while getting ranks redis!! going to set from mongo", map[string]interface{}{
			"error": err.Error(),
		})
		ranks, err = b.storer.GetAllRanks(ctx)
		if err != nil {
			b.log.Errorc(ctx, "error while getting ranks from DB! aborting...", map[string]interface{}{
				"error": err.Error(),
			})
			return err
		}
		res, err := b.storer.Set(ctx, "ranks", &ranks, 0)
		if err != nil {
			b.log.Errorc(ctx, "error while setting ranks to Redis!", map[string]interface{}{
				"error": err.Error(),
				"res":   res,
			})
			return err
		}
	}
	rank := getRankById(u.Rank, ranks)
	pts := float64(rank.PointsPerQuestion())
	payload.Score = int64(pts)
	quest, err := b.storer.GetQuestionByID(ctx, payload.ChallengeId, payload.QuestionId)
	if err != nil {
		b.log.Errorc(ctx, "error in GetQuestionByID", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	differenceInNanoseconds := time.Now().UnixNano() - quest.StartedAt
	differenceInSeconds := differenceInNanoseconds / 1_000_000_000
	payload.TimeTaken = differenceInSeconds
	err = b.storer.UpdateUserMetrics(ctx, payload.UserId, payload.Language, payload.IsCorrect, float64(payload.TimeTaken), payload.CodeQuality, payload.Score)
	if err != nil {
		b.log.Errorc(ctx, "error in UpdateUserMetrics", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (b *Business) UpdateUserStats(ctx context.Context, payload UpdatePayload) error {
	// type OverAllUser struct {
	// 	AttemptedQuestions []string  `json:"attempted_questions,omitempty" db:"attempted_questions" bson:"attempted_questions"` // List of question IDs user has faced
	// 	NoAttempted        int64     `json:"no_attempted,omitempty" db:"no_attempted" bson:"no_attempted"`
	// 	TotalCorrect       int64     `json:"total_correct,omitempty" db:"total_correct" bson:"total_correct"`
	// 	TotalWrong         int64     `json:"total_wrong,omitempty" db:"total_wrong" bson:"total_wrong"`
	// }
	err := b.storer.UpdateUserStats(ctx, payload.UserId, payload.QuestionId, payload.Language, payload.IsCorrect)
	if err != nil {
		b.log.Errorc(ctx, "error in UpdateUserStats", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}
	return nil
}
