package user

import (
	"context"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserProfileByAccountID(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
	GetUserSettingsByAccountID(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
	GetTechnologyNamesByIDs(ctx context.Context, ids []uuid.UUID) ([]string, error)
	UpdateCurrentUserProfileByID(ctx context.Context, a *accountmodel.Account, u *usermodel.User) error
	GetAccountByID(ctx context.Context, id uuid.UUID) (*accountmodel.Account, error)
	UpdateProfilePictureByAccountID(ctx context.Context, newURL string, accountID uuid.UUID) error
	GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error)
}

type UserUsecase struct {
	r UserRepository
}

func NewUserUsecase(repository UserRepository) *UserUsecase {
	return &UserUsecase{
		r: repository,
	}
}
