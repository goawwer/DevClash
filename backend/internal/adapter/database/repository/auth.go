package repository

import (
	"context"
	"crypto/hmac"
	"database/sql"
	"errors"
	"time"

	"github.com/goawwer/devclash/internal/dto"
	"github.com/google/uuid"
)

type AuthRepository interface {
	StoreRefreshToken(ctx context.Context, rec dto.RefreshTokenRecord) error
	ConsumeRefreshToken(ctx context.Context, id uuid.UUID, accountID uuid.UUID, providedHash string) (bool, uuid.UUID, error)
	CleanupExpiredRefreshTokens(ctx context.Context) error
}

func (r *ApplicationRepository) existsAccountByID(ctx context.Context, id uuid.UUID) (bool, error) {
	var exists bool
	err := r.GetContext(ctx, &exists, `SELECT EXISTS(SELECT 1 FROM accounts WHERE id = $1)`, id)
	return exists, err
}

func (r *ApplicationRepository) StoreRefreshToken(ctx context.Context, rec dto.RefreshTokenRecord) error {
	_, err := r.ExecContext(ctx, `
		INSERT INTO refresh_tokens (id, account_id, token_hash, expires_at)
		VALUES ($1, $2, $3, $4)
	`, rec.ID, rec.AccountID, string(rec.TokenHash), rec.ExpiresAt)

	return err
}

func (r *ApplicationRepository) ConsumeRefreshToken(ctx context.Context, id uuid.UUID, accountID uuid.UUID, providedHash string) (bool, uuid.UUID, error) {
	var (
		storedHash []byte
		consumedAt *time.Time
		expiresAt  time.Time
	)

	exists, err := r.existsAccountByID(ctx, accountID)
	if err != nil {
		return false, uuid.Nil, err
	}

	if !exists {
		return false, uuid.Nil, err
	}

	tx, err := r.BeginTx(ctx, nil)
	if err != nil {
		return false, uuid.Nil, err
	}
	defer r.MustBegin().Rollback()

	err = tx.QueryRowContext(ctx, `
		SELECT account_id, token_hash, consumed_at, expires_at FROM refresh_tokens
		WHERE id = $1 FOR UPDATE
	`, id).Scan(&accountID, &storedHash, &consumedAt, &expiresAt)

	if errors.Is(err, sql.ErrNoRows) {
		return false, uuid.Nil, nil
	}

	if err != nil {
		return false, uuid.Nil, err
	}

	// if mismatch -> invalid / stolen token
	if !hmac.Equal(storedHash, []byte(providedHash)) {
		return false, uuid.Nil, errors.New("invalid refresh token")
	}

	// if refresh token has been already used / expired
	if consumedAt != nil || time.Now().After(expiresAt) {
		return false, accountID, nil
	}

	// mark consumed
	_, err = tx.ExecContext(ctx, `
		UPDATE refresh_tokens SET consumed_at = $1 WHERE id = $2
	`, time.Now(), id)

	if err != nil {
		return false, uuid.Nil, err
	}

	if err := tx.Commit(); err != nil {
		return false, uuid.Nil, err
	}

	return true, accountID, nil
}

func (r *ApplicationRepository) CleanupExpiredRefreshTokens(ctx context.Context) error {
	q := `
		DELETE FROM refresh_tokens
		WHERE expires_at < NOW() OR consumed_at IS NOT NULL
	`

	_, err := r.ExecContext(ctx, q)

	return err
}
