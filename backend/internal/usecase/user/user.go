package user

import (
	"context"

	"github.com/goawwer/devclash/internal/domain/usermodel"
)

func (u *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	return u.r.GetUserByEmail(ctx, email)
}
