package user

import (
	"context"

	"github.com/goawwer/devclash/internal/dto"
	"github.com/google/uuid"
)

func (u *UserUsecase) GetUserProfileByID(ctx context.Context, id uuid.UUID) (*dto.UserProfile, error) {
	return u.r.GetUserProfileByID(ctx, id)
}

func (u *UserUsecase) GetUserSettingsByID(ctx context.Context, id uuid.UUID) (*dto.UserProfileSettings, error) {
	return u.r.GetUserSettingsByID(ctx, id)
}
