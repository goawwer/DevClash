package event

import (
	"context"

	eventmodel "github.com/goawwer/devclash/internal/domain/event_model"
	"github.com/google/uuid"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, e *eventmodel.Event) error
	GetEventTypeIDByName(ctx context.Context, name string) (uuid.UUID, error)
	GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error)
	GetOrganizerIDByAccountID(ctx context.Context, accountID uuid.UUID) (uuid.UUID, error)
	UpdateEventPictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error
}

type EventUsecase struct {
	r EventRepository
}

func NewEventUsecase(repository EventRepository) *EventUsecase {
	return &EventUsecase{
		r: repository,
	}
}
