package auth

import (
	"context"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	organizermodel "github.com/goawwer/devclash/internal/domain/organizer_model"
	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/google/uuid"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, a *accountmodel.Account, u *usermodel.User) error
	CreateOrganizer(ctx context.Context, a *accountmodel.Account, org *organizermodel.OrganizerAccount) error
	GetAccountByEmail(ctx context.Context, email string) (*accountmodel.Account, error)
	UpdateLogoByCreatorID(ctx context.Context, accountID uuid.UUID, newURL string) error
}

type AuthUsecase struct {
	r AuthRepository
}

func NewAuthUsecase(r AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		r: r,
	}
}
