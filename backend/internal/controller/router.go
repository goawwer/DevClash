package controller

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	_ "github.com/goawwer/devclash/docs"
	"github.com/goawwer/devclash/internal/controller/handlers/api"
	"github.com/goawwer/devclash/internal/controller/handlers/auth"
	"github.com/goawwer/devclash/internal/usecase"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router(u *usecase.AppUsecase) http.Handler {
	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(c.Handler)
	r.Use(chiMiddleware.Timeout(time.Second * 60))
	r.Use(logger.ChiLogger)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Mount("/auth", auth.New(u))
	r.Mount("/api", api.New(u))

	return r
}
