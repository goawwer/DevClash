package usecase

import (
	"github.com/goawwer/devclash/internal/usecase/auth"
	"github.com/goawwer/devclash/internal/usecase/event"
	"github.com/goawwer/devclash/internal/usecase/organizer"
	"github.com/goawwer/devclash/internal/usecase/team"
	"github.com/goawwer/devclash/internal/usecase/user"
)

type AppRepo interface {
	auth.AuthRepository
	user.UserRepository
	organizer.OrganizerRepository
	team.TeamRepository
	event.EventRepository
}

type AppUsecase struct {
	Auth  *auth.AuthUsecase
	User  *user.UserUsecase
	Org   *organizer.OrganizerUsecase
	Team  *team.TeamUsecase
	Event *event.EventUsecase
}

func NewUsecase(a AppRepo) *AppUsecase {
	auth := auth.NewAuthUsecase(a)
	user := user.NewUserUsecase(a)
	org := organizer.NewOrgUsecase(a)
	team := team.NewTeamUsecase(a)
	event := event.NewEventUsecase(a)

	return &AppUsecase{
		Auth:  auth,
		User:  user,
		Org:   org,
		Team:  team,
		Event: event,
	}
}
