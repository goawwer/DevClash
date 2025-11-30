package auth

import (
	"context"
	"time"

	"github.com/goawwer/devclash/internal/domain"
	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	organizermodel "github.com/goawwer/devclash/internal/domain/organizer_model"
	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/goawwer/devclash/pkg/mail"
	"github.com/goawwer/devclash/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthUsecase) SignUpUser(ctx context.Context, input dto.SignUpForm) error {
	if err := domain.Validate(dto.SignUpForm{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}); err != nil {
		return err
	}

	hashedPassword, err := utils.CreateHashPassword(input.Password)
	if err != nil {
		return err
	}

	account := &accountmodel.Account{
		Email:          input.Email,
		Role:           "user",
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
	}

	user := &usermodel.User{
		Username: input.Username,
	}

	return a.r.CreateUser(ctx, account, user)
}

func (a *AuthUsecase) Login(ctx context.Context, input dto.LoginForm) (uuid.UUID, string, error) {
	account, err := a.r.GetAccountByEmail(ctx, input.Email)
	if err != nil {
		return uuid.Nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.HashedPassword), []byte(input.Password))
	if err != nil {
		return uuid.Nil, "", err
	}

	return account.ID, account.Role, nil
}

func (a *AuthUsecase) SignUpOrganizer(ctx context.Context, input dto.SignUpInputOrganizerDetails) (uuid.UUID, error) {
	hashedPassword, err := utils.CreateHashPassword(input.Password)
	if err != nil {
		return uuid.Nil, err
	}

	account := &accountmodel.Account{
		Email:          input.Email,
		HashedPassword: hashedPassword,
		Role:           "organizer",
		CreatedAt:      time.Now(),
	}

	organizerID := uuid.New()

	organizer := &organizermodel.OrganizerAccount{
		ID:   organizerID,
		Name: input.Name,
		Details: &organizermodel.Details{
			OrganizerID: organizerID,
			LogoURL:     &input.LogoURL,
			Color:       &input.Color,
		},
	}

	if err := a.r.CreateOrganizer(ctx, account, organizer); err != nil {
		logger.Error("failed to create organizer: ", err)
		return uuid.Nil, err
	}

	return organizerID, mail.OrganizerSignupInfo(account, organizer)
}

func (a *AuthUsecase) UpdateLogoByCreatorID(ctx context.Context, accountID uuid.UUID, newURL string) error {
	return a.r.UpdateLogoByCreatorID(ctx, accountID, newURL)
}
