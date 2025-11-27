package user

import (
	"context"

	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserProfileByID(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
}

type UserUsecase struct {
	r UserRepository
}

func NewUserUsecase(repository UserRepository) *UserUsecase {
	return &UserUsecase{
		r: repository,
	}
}
