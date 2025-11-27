package accountmodel

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID             uuid.UUID `db:"id" json:"id"`
	Email          string    `db:"email" json:"email"`
	HashedPassword string    `db:"hashed_password" json:"hashed_password"`
	Role           string    `db:"role" json:"role"`
	Disabled       bool      `db:"disabled" json:"disabled"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}
