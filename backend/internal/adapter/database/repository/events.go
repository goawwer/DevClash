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
	GetEventByID(ctx context.Context, id uuid.UUID) (*eventmodel.Event, error)
	GetEventTypeNameByID(ctx context.Context, id uuid.UUID) (string, error)
}

func (r *ApplicationRepository) CreateEvent(ctx context.Context, e *eventmodel.Event) error {
	return r.RunInTransaction(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO events (id, organizer_id, type_id, title, created_at)
			VALUES ($1, $2, $3, $4, $5)
		`, e.ID, e.OrganizerID, e.TypeID, e.Title, e.CreatedAt)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, `
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

		_, err = tx.ExecContext(ctx, `
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
			_, err = tx.ExecContext(ctx, `
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

func (r *ApplicationRepository) GetEventTypeIDByName(ctx context.Context, name string) (uuid.UUID, error) {
	var id uuid.UUID

	return id, r.GetContext(ctx, &id, `
		SELECT id FROM event_types 
		WHERE name = $1
	`, name)
}

func (r *ApplicationRepository) GetEventTypeNameByID(ctx context.Context, id uuid.UUID) (string, error) {
	var name string

	return name, r.GetContext(ctx, &name, `
		SELECT name FROM event_types 
		WHERE id = $1
	`, id)
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

func (r *ApplicationRepository) GetEventByID(ctx context.Context, id uuid.UUID) (*eventmodel.Event, error) {
	var e eventmodel.Event

	return &e, r.GetContext(ctx, &e, `
		SELECT
            e.id, e.organizer_id, e.type_id, e.title, e.created_at, e.updated_at, e.is_finished,
            
            ed.event_picture_url, ed.start_time, ed.end_time, ed.description, ed.prize,
            
            ep.is_online, ep.is_free, ep.number_of_teams, ep.team_size,
            
            COALESCE(
                (
                    SELECT ARRAY_AGG(et.technology_id::text)
                    FROM events_technologies et
                    WHERE et.event_id = e.id
                ),
                '{}'
            ) AS technologies,
            
            COALESCE(
                (
                    SELECT ARRAY_AGG(etm.team_id::text)
                    FROM events_teams etm
                    WHERE etm.event_id = e.id
                ),
                '{}'
            ) AS teams_ids
        FROM events e
        	LEFT JOIN events_details ed ON e.id = ed.event_id
        	LEFT JOIN event_properties ep ON e.id = ep.event_id
        WHERE e.id = $1
	`, id)
}
