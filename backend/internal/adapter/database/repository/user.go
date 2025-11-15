package repository

import (
	"context"

	"github.com/goawwer/devclash/internal/domain/usermodel"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error)
	Create(ctx context.Context, user *usermodel.User) error
}

func (r *ApplicationRepository) Create(ctx context.Context, user *usermodel.User) error {
	query := `
		INSERT INTO users (email, username, hashed_password, created_at)
			VALUES ($1, $2, $3, $4)
	`

	_, err := r.ExecContext(ctx, query, user.Email, user.Username, user.HashedPassword, user.CreatedAt)

	return err
}

func (r *ApplicationRepository) GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	var u usermodel.User

	query := `
		SELECT * FROM users
		WHERE email = $1
	`

	if err := r.GetContext(ctx, &u, query, email); err != nil {
		return nil, err
	}

	return &u, nil
}
