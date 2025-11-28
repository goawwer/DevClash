package organizer

import (
	"github.com/go-chi/chi/v5"
	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/usecase/organizer"
)

type OrganizerHandler struct {
	*organizer.OrganizerUsecase
}

func handler(o *organizer.OrganizerUsecase) *OrganizerHandler {
	return &OrganizerHandler{o}
}

func New(usecase *organizer.OrganizerUsecase) *chi.Mux {
	r := chi.NewRouter()

	h := handler(usecase)

	r.Get("/{id}", wrapper.AuthWrap(h.GetOrganizerByID))

	return r
}
