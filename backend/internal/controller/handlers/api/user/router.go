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
		u.Get("/current/profile", wrapper.AuthWrap(h.GetCurrentUserProfile))
	})

	return r
}
