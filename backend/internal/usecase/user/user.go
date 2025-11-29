package user

import (
	"context"

	"github.com/goawwer/devclash/internal/dto"
	"github.com/google/uuid"
)

func (u *UserUsecase) GetUserProfileByID(ctx context.Context, id uuid.UUID) (*dto.UserProfile, error) {
	return u.r.GetUserProfileByID(ctx, id)
}
