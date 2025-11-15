package usermodel

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID   `db:"id" json:"id"`
	Email          string      `db:"email" json:"email"`
	Username       string      `db:"username" json:"username"`
	HashedPassword string      `db:"hashed_password" json:"-"`
	CreatedAt      time.Time   `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time   `db:"updated_at" json:"updated_at"`
	IsAdmin        bool        `db:"is_admin" json:"is_admin"`
	Disabled       bool        `db:"disabled" json:"disabled"`
	Details        *UserDetail `json:"user_details"`
}

type UserEntity struct {
	ID             uuid.UUID `db:"id"`
	Email          string    `db:"email"`
	Username       string    `db:"username"`
	HashedPassword string    `db:"hashed_password"`
	IsAdmin        bool      `db:"is_admin"`
}
