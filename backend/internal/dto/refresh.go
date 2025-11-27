package dto

import (
	"time"

	"github.com/google/uuid"
)

type RefreshTokenRecord struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	TokenHash  string
	ExpiresAt  time.Time
	ConsumedAt *time.Time
}
