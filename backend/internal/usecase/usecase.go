package usecase

import (
	"github.com/goawwer/devclash/internal/usecase/auth"
	"github.com/goawwer/devclash/internal/usecase/user"
)

type AppRepo interface {
	auth.AuthRepository
	user.UserRepository
}

type AppUsecase struct {
	Auth *auth.AuthUsecase
	User *user.UserUsecase
}

func NewUsecase(a AppRepo) *AppUsecase {
	auth := auth.NewAuthUsecase(a)
	user := user.NewUserUsecase(a)

	return &AppUsecase{
		Auth: auth,
		User: user,
	}
}
