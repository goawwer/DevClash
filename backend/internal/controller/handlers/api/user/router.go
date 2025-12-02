package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/usecase/user"
)

type UserHandler struct {
	*user.UserUsecase
}

func handler(u *user.UserUsecase) *UserHandler {
	return &UserHandler{u}
}

func New(usecase *user.UserUsecase) *chi.Mux {
	r := chi.NewRouter()

	h := handler(usecase)

	r.Group(func(u chi.Router) {
		u.Get("/me/profile", wrapper.AuthWrap(h.GetCurrentUserProfile))
		u.Get("/me/settings", wrapper.AuthWrap(h.GetCurrentUserSettings))
		u.Put("/me", wrapper.AuthWrap(h.UpdateUserProfile))
		r.Get("/me/team_id", wrapper.AuthWrap(h.GetTeamIDForCurrentUser))
	})

	return r
}
