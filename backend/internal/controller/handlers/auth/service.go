package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/usecase"
	"github.com/goawwer/devclash/internal/usecase/auth"
)

type AuthHandler struct {
	*auth.AuthUsecase
}

func handler(a *auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{a}
}

func New(usecase *usecase.AppUsecase) *chi.Mux {
	r := chi.NewRouter()

	h := handler(usecase.Auth)

	r.Post("/signup", wrapper.PublicWrap(h.SignUp))
	r.Post("/login", wrapper.PublicWrap(h.Login))
	r.Post("/refresh", wrapper.PublicWrap(h.Refresh))

	return r
}
