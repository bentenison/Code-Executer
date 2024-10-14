package authdb

import "database/sql"

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDB struct {
	ID           sql.NullString `json:"id"`
	Username     sql.NullString `json:"username"`
	Email        sql.NullString `json:"email"`
	PasswordHash sql.NullString `json:"password_hash"`
	FirstName    sql.NullString `json:"first_name"`
	LastName     sql.NullString `json:"last_name"`
	Role         sql.NullString `json:"role"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}
