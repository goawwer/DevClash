package repository

import (
	"context"
	"errors"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrEmailTaken = errors.New("email already exists")
)

type UserRepository interface {
	CreateUser(ctx context.Context, a *accountmodel.Account, u *usermodel.User) error
	GetUserProfileByID(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
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

func (r *ApplicationRepository) GetUserProfileByID(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	var u usermodel.User

	return &u, r.GetContext(ctx, &u, `
		SELECT * FROM users
		WHERE account_id = $1
	`, id)
}
