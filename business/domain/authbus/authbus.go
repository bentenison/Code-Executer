package authbus

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/sdk/delegate"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type Storer interface {
	CreateUser(ctx context.Context, user *User) (string, error)
	GetUser(ctx context.Context, u string) (*User, error)
	ListUsers(ctx context.Context) ([]*User, error)
}
type Business struct {
	log      *logger.CustomLogger
	delegate *delegate.Delegate
	ds       mux.DataSource
	storer   Storer
}

func NewBusiness(log *logger.CustomLogger, delegate *delegate.Delegate, ds mux.DataSource, storer Storer) *Business {
	return &Business{
		log:      log,
		delegate: delegate,
		ds:       ds,
		storer:   storer,
	}
}

func (b *Business) CreateUser(ctx context.Context, user *User) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		b.log.Errorc(ctx, "error in generating hashed password", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	user.PasswordHash = string(hashedpassword)
	return b.storer.CreateUser(ctx, user)
}
func (b *Business) AuthenticateUser(ctx context.Context, username, password, jwtKey string) (string, error) {
	// check if user exists in db
	user, err := b.storer.GetUser(ctx, username)
	if err != nil {
		b.log.Errorc(ctx, "error in getting user", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	// check if password hash matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		b.log.Errorc(ctx, "error in comparing password", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	tkn, err := generateJWT(user.Username, user.Role, jwtKey)
	if err != nil {
		b.log.Errorc(ctx, "error in generating token", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	// add session
	// return user token with role info
	return tkn, nil
}
func (b *Business) AuthorizeUser(ctx context.Context, token, jwtKey string) (*Claims, error) {
	claims, err := validateJWT(token, jwtKey)
	if err != nil {
		b.log.Errorc(ctx, "error in validating token", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	// if token valid get userid
	// check session
	return claims, nil
}
func (b *Business) ListUsers(ctx context.Context) ([]*User, error) {
	return b.storer.ListUsers(ctx)
}
func validateJWT(token, jwtKey string) (*Claims, error) {
	tkn, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tkn.Claims.(*Claims); ok && tkn.Valid {
		return claims, nil
	} else {
		return nil, errors.New("error token invalid")
	}
}
func generateJWT(username, role, jwtKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserId: username,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "auth-service",
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			NotBefore: &jwt.NumericDate{
				Time: time.Now().Add(2 * time.Hour),
			},
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(2 * time.Hour),
			},
		},
	})
	return token.SignedString([]byte(jwtKey))
}

var adjectives = []string{
	"Swift", "Clever", "Bold", "Brave", "Quick", "Friendly", "Creative", "Lively", "Mighty", "Epic",
	"Fearless", "Vibrant", "Daring", "Witty", "Loyal", "Radiant", "Noble", "Unique", "Powerful", "Sleek",
	"Smart", "Inventive", "Cheerful", "Dynamic", "Majestic", "Charming", "Eager", "Graceful", "Energetic",
	"Adventurous", "Playful", "Confident", "Resourceful", "Inspirational", "Generous", "Wise", "Thoughtful",
	"Curious", "Serene", "Harmonious", "Diligent", "Optimistic", "Zesty", "Pioneering", "Determined", "Faithful",
	"Fearless", "Radiant", "Vigorous", "Brilliant", "Dazzling", "Luminous", "Sharp", "Resourceful", "Vigilant",
	"Vast", "Eloquent", "Inventive", "Noble", "Sincere", "True", "Brisk", "Unstoppable", "Breezy", "Unyielding",
	"Gallant", "Majestic", "Dazzling", "Humble", "Vigorous", "Imaginative", "Steadfast", "Gracious", "Jovial",
	"Zany", "Harmonious", "Vivid", "Energetic", "Vigilant", "Whimsical", "Noble", "Fierce", "Valiant", "Zealous",
}
var nouns = []string{
	"Explorer", "Coder", "Traveler", "Artist", "Dreamer", "Thinker", "Ninja", "Guru", "Wizard", "Hero",
	"Champion", "Inventor", "Seeker", "Warrior", "Innovator", "Achiever", "Philosopher", "Maverick",
	"Magician", "Pioneer", "Creator", "Adventurer", "Leader", "Scribe", "Savior", "Shaman", "Scientist",
	"Knight", "Titan", "Rebel", "Mogul", "Strategist", "Defender", "Mastermind", "Tactician", "Vanguard",
	"Prodigy", "Conqueror", "Wanderer", "Trailblazer", "Commander", "Visionary", "Pathfinder", "Savant",
	"Guardian", "Hacker", "Druid", "Alchemist", "Scholar", "Warlock", "Samurai", "Mender", "Diplomat",
	"Orator", "Architect", "Curator", "Healer", "Ambassador", "Influencer", "Outlaw", "Magister", "Ranger",
	"Strategist", "Reformer", "Catalyst", "Pundit", "Protector", "Enchanter", "Noble", "Pursuer", "Sculptor",
	"Detective", "Nomad", "Librarian", "Historian", "Jedi", "Prophet", "Pathfinder", "Champion", "Conqueror",
	"Voyager", "Tactician", "Adventurer", "Artisan", "Sculptor", "Maverick", "Sentinel", "Bard", "Genius",
	"Vanguard", "Mogul", "Philosopher", "Master", "Devotee", "Genius", "Trailblazer", "Shapeshifter",
	"Creator", "Phantom", "Scholar", "Virtuoso", "Seer", "Illusionist", "Navigator", "Strategist",
	"Collector", "Inventor", "Seeker", "Sage", "Alchemist",
}

func generateUsername(firstName, lastName string) string {
	// Normalize input to lowercase for better formatting
	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)

	// Generate a random adjective and noun from predefined lists
	rand.Seed(uint64(time.Now().UnixNano()))
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]

	// Combine elements into a meaningful username (first name, last name, adjective, noun)
	username := fmt.Sprintf("%s_%s_%s_%s", firstName, lastName, adjective, noun)

	// Ensure uniqueness by adding a random number at the end if needed
	username = appendUniqueNumber(username)
	return username
}

// Function to ensure uniqueness by appending a random number (in case username is already taken).
func appendUniqueNumber(username string) string {
	randomSuffix := rand.Intn(1000) // Generate a random number
	return fmt.Sprintf("%s%d", username, randomSuffix)
}
