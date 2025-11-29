package repository

import (
	"context"
	"errors"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrEmailTaken = errors.New("email already exists")
)

type UserRepository interface {
	CreateUser(ctx context.Context, a *accountmodel.Account, u *usermodel.User) error
	GetUserProfileByID(ctx context.Context, id uuid.UUID) (*dto.UserProfile, error)
}

func (r *ApplicationRepository) CreateUser(ctx context.Context, a *accountmodel.Account, u *usermodel.User) error {
	return r.RunInTransaction(ctx, func(tx *sqlx.Tx) error {
		err := tx.QueryRowxContext(ctx, `
			INSERT INTO accounts (email, hashed_password, role, created_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`, a.Email, a.HashedPassword, a.Role, a.CreatedAt).Scan(&a.ID)
		if err != nil {
			return mapError(err)
		}

		_, err = tx.ExecContext(ctx, `
		INSERT INTO users (account_id, username)
			VALUES ($1, $2)
		`, a.ID, u.Username)

		return err
	})
}

func (r *ApplicationRepository) GetUserProfileByID(ctx context.Context, id uuid.UUID) (*dto.UserProfile, error) {
	var p dto.UserProfile

	return &p, r.GetContext(ctx, &p, `
		SELECT 
			u.username, u.profile_picture_url, u.bio, u.profile_status,
			u.participations_count, u.wins_count,
			ARRAY_AGG(t.name) FILTER (WHERE t.name IS NOT NULL) AS tech_stack
		FROM users u
			LEFT JOIN users_skills us ON u.id = us.user_id
			LEFT JOIN technologies t ON us.technology_id = t.id
		WHERE account_id = $1
		GROUP BY u.id
	`, id)
}
