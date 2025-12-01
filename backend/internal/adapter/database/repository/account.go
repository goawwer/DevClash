package repository

import (
	"context"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	"github.com/google/uuid"
)

type AccountRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (*accountmodel.Account, error)
	GetAccountByID(ctx context.Context, id uuid.UUID) (*accountmodel.Account, error)
}

func (r *ApplicationRepository) GetAccountByEmail(ctx context.Context, email string) (*accountmodel.Account, error) {
	var a accountmodel.Account

	return &a, r.GetContext(ctx, &a, `
		SELECT * FROM accounts
		WHERE email = $1
	`, email)
}

func (r *ApplicationRepository) GetAccountByID(ctx context.Context, id uuid.UUID) (*accountmodel.Account, error) {
	var a accountmodel.Account

	return &a, r.GetContext(ctx, &a, `
		SELECT * FROM accounts 
		WHERE id = $1
	`, id)
}
