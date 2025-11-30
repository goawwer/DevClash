package api

import (
	"github.com/go-chi/chi/v5"
	eventHandler "github.com/goawwer/devclash/internal/controller/handlers/api/event"
	orgHandler "github.com/goawwer/devclash/internal/controller/handlers/api/organizer"
	teamHandler "github.com/goawwer/devclash/internal/controller/handlers/api/team"
	userHandler "github.com/goawwer/devclash/internal/controller/handlers/api/user"
	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/usecase"
	"github.com/goawwer/devclash/middleware"
)

func New(u *usecase.AppUsecase) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Middleware)

	// api generic logic
	r.Group(func(r chi.Router) {
		r.Post("/logout", wrapper.AuthWrap(Logout))
		r.Get("/image", wrapper.AuthWrap(GetS3Url))
	})

	r.Mount("/users", userHandler.New(u.User))
	r.Mount("/organizers", orgHandler.New(u.Org))
	r.Mount("/teams", teamHandler.New(u.Team))
	r.Mount("/events", eventHandler.New(u.Event))

	return r
}
