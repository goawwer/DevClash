package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/goawwer/devclash/docs"
	"github.com/goawwer/devclash/internal/controller/handlers/api"
	"github.com/goawwer/devclash/internal/controller/handlers/auth"
	"github.com/goawwer/devclash/internal/usecase"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router(u *usecase.AppUsecase) http.Handler {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	authHandler := auth.New(u)
	apiHander := api.New(u)

	r.Mount("/auth", authHandler)
	r.Mount("/api", apiHander)

	return r
}
