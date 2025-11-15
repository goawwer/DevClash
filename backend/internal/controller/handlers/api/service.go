package api

import (
	"github.com/go-chi/chi/v5"
	userHandler "github.com/goawwer/devclash/internal/controller/handlers/api/user"
	"github.com/goawwer/devclash/internal/usecase"
	"github.com/goawwer/devclash/middleware"
)

func New(u *usecase.AppUsecase) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Middleware)

	userHandler := userHandler.New(u.User)

	r.Mount("/users", userHandler)

	return r
}
