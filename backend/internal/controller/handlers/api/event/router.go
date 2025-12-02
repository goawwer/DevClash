package event

import (
	"github.com/go-chi/chi/v5"
	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/usecase/event"
)

type EventHandler struct {
	*event.EventUsecase
}

func handler(e *event.EventUsecase) *EventHandler {
	return &EventHandler{e}
}

func New(usecase *event.EventUsecase) *chi.Mux {
	r := chi.NewRouter()

	h := handler(usecase)

	r.Post("/create", wrapper.AuthWrap(h.Create))
	r.Post("/join", wrapper.AuthWrap(h.TeamJoinEvent))
	r.Get("/all", wrapper.AuthWrap(h.GetAllEvents))
	r.Get("/{id}", wrapper.AuthWrap(h.GetEventPage))

	return r
}
