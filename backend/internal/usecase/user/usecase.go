package user

import (
	"context"

	"github.com/goawwer/devclash/internal/domain/usermodel"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error)
}

type UserUsecase struct {
	r UserRepository
}

func NewUserUsecase(repository UserRepository) *UserUsecase {
	return &UserUsecase{
		r: repository,
	}
}
