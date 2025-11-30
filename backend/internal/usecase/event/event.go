package event

import (
	"context"
	"time"

	"github.com/goawwer/devclash/internal/domain"
	eventmodel "github.com/goawwer/devclash/internal/domain/event_model"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/google/uuid"
)

func (e *EventUsecase) Create(ctx context.Context, id uuid.UUID, role string, input *dto.EventCreationRequest) error {
	if role == "user" {
		return domain.ErrNotForUserRole
	}

	eventID := uuid.New()

	orgID, err := e.r.GetOrganizerIDByAccountID(ctx, id)
	if err != nil {
		return err
	}

	eventTypeID, err := e.r.GetEventTypeIDByName(ctx, input.Type)
	if err != nil {
		return err
	}

	techIDs := make([]uuid.UUID, 0, len(input.TechStack))
	for _, name := range input.TechStack {
		techID, err := e.r.GetTechnologyIDByName(ctx, name)
		if err != nil {
			return err
		}
		techIDs = append(techIDs, techID)
	}

	return e.r.CreateEvent(ctx, &eventmodel.Event{
		ID:          eventID,
		OrganizerID: orgID,
		TypeID:      eventTypeID,
		Title:       input.Title,
		CreatedAt:   time.Now(),

		Properties: &eventmodel.Properties{
			EventID:       eventID,
			IsOnline:      input.IsOnline,
			IsFree:        input.IsFree,
			NumberOfTeams: input.NumberOfTeams,
			TeamSize:      input.TeamSize,
		},

		Details: &eventmodel.Details{
			EventID:         eventID,
			EventPictureURL: input.EventPictureURL,
			StartTime:       input.StartTime,
			EndTime:         input.EndTime,
			Description:     input.Description,
			Prize:           input.Prize,
		},

		Technologies: techIDs,
	})
}

func (e *EventUsecase) UpdatePictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error {
	return e.r.UpdateEventPictureByCreatorID(ctx, newURL, accountID)
}
