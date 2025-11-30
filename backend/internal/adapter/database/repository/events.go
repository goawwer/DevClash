package repository

import (
	"context"

	eventmodel "github.com/goawwer/devclash/internal/domain/event_model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, e *eventmodel.Event) error
	GetEventTypeIDByName(ctx context.Context, name string) (uuid.UUID, error)
	UpdateEventPictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error
}

func (r *ApplicationRepository) CreateEvent(ctx context.Context, e *eventmodel.Event) error {
	return r.RunInTransaction(ctx, func(tx *sqlx.Tx) error {
		_, err := r.ExecContext(ctx, `
			INSERT INTO events (id, organizer_id, type_id, title, created_at)
			VALUES ($1, $2, $3, $4, $5)
		`, e.ID, e.OrganizerID, e.TypeID, e.Title, e.CreatedAt)
		if err != nil {
			return err
		}

		_, err = r.ExecContext(ctx, `
			INSERT INTO event_properties (event_id, is_online, is_free, number_of_teams, team_size)
			VALUES ($1, $2, $3, $4, $5)
		`, e.Properties.EventID,
			e.Properties.IsOnline,
			e.Properties.IsFree,
			e.Properties.NumberOfTeams,
			e.Properties.TeamSize,
		)
		if err != nil {
			return err
		}

		_, err = r.ExecContext(ctx, `
				INSERT INTO events_details (event_id, event_picture_url, start_time, end_time, description, prize)
				VALUES ($1, $2, $3, $4, $5, $6)
			`,
			e.Details.EventID,
			e.Details.EventPictureURL,
			e.Details.StartTime,
			e.Details.EndTime,
			e.Details.Description,
			e.Details.Prize,
		)
		if err != nil {
			return err

		}

		for _, techID := range e.Technologies {
			_, err = r.ExecContext(ctx, `
				INSERT INTO events_technologies (event_id, technology_id)
				VALUES ($1, $2)		
			`, e.ID, techID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// func (r *ApplicationRepository) GetByID(ctx context.Context, id uuid.UUID) ()

func (r *ApplicationRepository) GetEventTypeIDByName(ctx context.Context, name string) (uuid.UUID, error) {
	var id uuid.UUID

	return id, r.GetContext(ctx, &id, `
		SELECT id FROM event_types 
		WHERE name = $1
	`, name)
}

func (r *ApplicationRepository) UpdateEventPictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error {
	_, err := r.ExecContext(ctx, `
		UPDATE events_details ed
        	SET event_picture_url = $1
        FROM events e
        JOIN organizers o ON e.organizer_id = o.id
        	WHERE ed.event_id = e.id
         	AND o.account_id = $2
	`, newURL, accountID)

	return err
}
