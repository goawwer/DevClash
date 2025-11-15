package auth

import (
	"context"

	"github.com/goawwer/devclash/internal/domain/usermodel"
)

type AuthRepository interface {
	Create(ctx context.Context, Auth *usermodel.User) error
	GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error)
}

type AuthUsecase struct {
	r AuthRepository
}

func NewAuthUsecase(r AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		r: r,
	}
}
