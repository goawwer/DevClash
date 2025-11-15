package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/goawwer/devclash/internal/domain/usermodel"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthUsecase) SignUp(ctx context.Context, input dto.SignUpInput) error {
	hashedPassword, err := utils.CreateHashPassword(input.Password)
	if err != nil {
		return err
	}

	user := &usermodel.User{
		Email:          input.Email,
		Username:       input.Email,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
	}

	if err := usermodel.Validate(user); err != nil {
		return err
	}

	return a.r.Create(ctx, user)
}

func (a *AuthUsecase) Login(ctx context.Context, input dto.LoginInput) (uuid.UUID, bool, error) {
	u, err := a.r.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return uuid.Nil, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(input.Password))
	if err != nil {
		return uuid.Nil, false, fmt.Errorf("invalid credentials: %w", err)
	}

	return u.ID, u.IsAdmin, nil
}
