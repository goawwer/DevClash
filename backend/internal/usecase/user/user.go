package user

import (
	"context"

	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/google/uuid"
)

func (u *UserUsecase) GetUserProfileByID(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	return u.r.GetUserProfileByID(ctx, id)
}
