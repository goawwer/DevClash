package event

import (
	"context"

	eventmodel "github.com/goawwer/devclash/internal/domain/event_model"
	teammodel "github.com/goawwer/devclash/internal/domain/team_model"
	"github.com/google/uuid"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, e *eventmodel.Event) error
	GetEventTypeIDByName(ctx context.Context, name string) (uuid.UUID, error)
	GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error)
	GetOrganizerIDByAccountID(ctx context.Context, accountID uuid.UUID) (uuid.UUID, error)
	UpdateEventPictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error
	GetEventByID(ctx context.Context, id uuid.UUID) (*eventmodel.Event, error)
	GetTechnologyNamesByIDs(ctx context.Context, ids []uuid.UUID) ([]string, error)
	GetEventTypeNameByID(ctx context.Context, id uuid.UUID) (string, error)
	GetTeamsByIDs(ctx context.Context, ids []uuid.UUID) ([]teammodel.Team, error)
}

type EventUsecase struct {
	r EventRepository
}

func NewEventUsecase(repository EventRepository) *EventUsecase {
	return &EventUsecase{
		r: repository,
	}
}
