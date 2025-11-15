package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/goawwer/devclash/internal/controller/handlers/api"
	"github.com/goawwer/devclash/internal/controller/handlers/auth"
	"github.com/goawwer/devclash/internal/usecase"
)

func Router(u *usecase.AppUsecase) http.Handler {
	r := chi.NewRouter()

	authHandler := auth.New(u)
	apiHander := api.New(u)

	r.Mount("/auth", authHandler)
	r.Mount("/api", apiHander)

	return r
}
