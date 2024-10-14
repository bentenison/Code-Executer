package authbus

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key (use environment variable or config for production)
var jwtSecret = []byte("super-secret-key")

// Struct to represent claims inside the JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	FirstName    string    `json:"first_name,omitempty"`
	LastName     string    `json:"last_name,omitempty"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
