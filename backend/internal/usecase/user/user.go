package user

import (
	"context"
	"time"

	"github.com/goawwer/devclash/internal/domain"
	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) GetUserProfileByID(ctx context.Context, id uuid.UUID) (*dto.UserProfile, error) {
	user, err := u.r.GetUserProfileByAccountID(ctx, id)
	if err != nil {
		return nil, err
	}

	techIDs := make([]uuid.UUID, 0, len(user.Technologies))
	for _, s := range user.Technologies {
		if s != "" {
			techIDs = append(techIDs, uuid.MustParse(s))
		}
	}

	techNames, err := u.r.GetTechnologyNamesByIDs(ctx, techIDs)
	if err != nil {
		return nil, err
	}

	return &dto.UserProfile{
		Username:            user.Username,
		ProfilePictureURL:   user.ProfilePictureURL,
		Bio:                 user.Bio,
		ProfileStatus:       user.ProfileStatus,
		ParticipationsCount: user.ParticipationsCount,
		WinsCount:           user.WinsCount,
		TechStack:           techNames,
	}, nil
}

func (u *UserUsecase) GetUserSettingsByID(ctx context.Context, id uuid.UUID) (*dto.UserGetProfileSettings, error) {
	user, err := u.r.GetUserSettingsByAccountID(ctx, id)
	if err != nil {
		return nil, err
	}

	techIDs := make([]uuid.UUID, 0, len(user.Technologies))
	for _, s := range user.Technologies {
		if s != "" {
			techIDs = append(techIDs, uuid.MustParse(s))
		}
	}

	techNames, err := u.r.GetTechnologyNamesByIDs(ctx, techIDs)
	if err != nil {
		return nil, err
	}

	return &dto.UserGetProfileSettings{
		Username:          user.Username,
		Email:             user.Email,
		Bio:               user.Bio,
		ProfileStatus:     user.ProfileStatus,
		TechStack:         techNames,
		ProfilePictureURL: user.ProfilePictureURL,
	}, nil
}

func (u *UserUsecase) UpdateCurrentUserProfileByID(ctx context.Context, accountID uuid.UUID, input *dto.UserUpdateProfileSettings) error {
	userAccount, err := u.r.GetAccountByID(ctx, accountID)
	if err != nil {
		return err
	}

	if input.NewPassword != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(userAccount.HashedPassword), []byte(input.OldPassword)); err != nil {
			return domain.ErrInvalidOldPassword
		}

		userAccount.HashedPassword, err = utils.CreateHashPassword(input.NewPassword)
		if err != nil {
			return err
		}
	}

	techIDs := make([]uuid.UUID, 0, len(input.TechStack))
	for _, name := range input.TechStack {
		techID, err := u.r.GetTechnologyIDByName(ctx, name)
		if err != nil {
			return err
		}
		techIDs = append(techIDs, techID)
	}

	techIDStrings := make([]string, len(techIDs))
	for i, id := range techIDs {
		techIDStrings[i] = id.String()
	}

	return u.r.UpdateCurrentUserProfileByID(ctx, &accountmodel.Account{
		ID:             userAccount.ID,
		Email:          userAccount.Email,
		HashedPassword: userAccount.HashedPassword,
		CreatedAt:      userAccount.CreatedAt,
		UpdatedAt:      time.Now(),
	}, &usermodel.User{
		Username:          input.Username,
		Bio:               input.Bio,
		ProfileStatus:     input.ProfileStatus,
		ProfilePictureURL: input.ProfilePictureURL,
		Technologies:      techIDStrings,
	})
}

func (u *UserUsecase) UpdatePictureByID(ctx context.Context, newURL string, accountID uuid.UUID) error {
	return u.r.UpdateProfilePictureByAccountID(ctx, newURL, accountID)
}
