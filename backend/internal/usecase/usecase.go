package usecase

import (
	"github.com/goawwer/devclash/internal/usecase/auth"
	"github.com/goawwer/devclash/internal/usecase/organizer"
	"github.com/goawwer/devclash/internal/usecase/team"
	"github.com/goawwer/devclash/internal/usecase/user"
)

type AppRepo interface {
	auth.AuthRepository
	user.UserRepository
	organizer.OrganizerRepository
	team.TeamRepository
}

type AppUsecase struct {
	Auth *auth.AuthUsecase
	User *user.UserUsecase
	Org  *organizer.OrganizerUsecase
	Team *team.TeamUsecase
}

func NewUsecase(a AppRepo) *AppUsecase {
	auth := auth.NewAuthUsecase(a)
	user := user.NewUserUsecase(a)
	org := organizer.NewOrgUsecase(a)
	team := team.NewTeamUsecase(a)

	return &AppUsecase{
		Auth: auth,
		User: user,
		Org:  org,
		Team: team,
	}
}
