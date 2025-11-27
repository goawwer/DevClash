package middleware

import (
	"context"
	"time"

	"github.com/goawwer/devclash/internal/dto"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type databaseIntegration interface {
	StoreRefreshToken(ctx context.Context, rec dto.RefreshTokenRecord) error
	ConsumeRefreshToken(ctx context.Context, id uuid.UUID, accountID uuid.UUID, providedHash string) (bool, uuid.UUID, error)
	CleanupExpiredRefreshTokens(ctx context.Context) error
}

type Config struct {
	Secret     string `env:"JWT_SECRET_CODE"`
	Repository databaseIntegration
}

type TokenPair struct {
	AccessToken     string
	AccessTokenExp  time.Time
	RefreshToken    string
	RefreshTokenExp time.Time
}

type CustomClaims struct {
	AccountID uuid.UUID `json:"account_id"`
	Role      string    `json:"role"`
	TokenType string    `json:"token_type"`
	jwt.RegisteredClaims
}

type contextKey string

var ClaimsKey contextKey = "claims"

type CustomCookie struct {
	Name  string
	Value string
	Exp   time.Duration
}
