package user

import (
	"context"

	"github.com/goawwer/devclash/internal/dto"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserProfileByID(ctx context.Context, id uuid.UUID) (*dto.UserProfile, error)
}

type UserUsecase struct {
	r UserRepository
}

func NewUserUsecase(repository UserRepository) *UserUsecase {
	return &UserUsecase{
		r: repository,
	}
}
