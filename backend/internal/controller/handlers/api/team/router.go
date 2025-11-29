package team

import (
	"github.com/go-chi/chi/v5"
	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/usecase/team"
)

type TeamHandler struct {
	*team.TeamUsecase
}

func handler(o *team.TeamUsecase) *TeamHandler {
	return &TeamHandler{o}
}

func New(usecase *team.TeamUsecase) *chi.Mux {
	r := chi.NewRouter()

	h := handler(usecase)

	r.Post("/create", wrapper.AuthWrap(h.Create))

	return r
}
