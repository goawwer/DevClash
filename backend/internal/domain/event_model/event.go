package eventmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Event struct {
	ID          uuid.UUID `db:"id" json:"id"`
	OrganizerID uuid.UUID `db:"organizer_id" json:"organizer_id"`
	TypeID      uuid.UUID `db:"type_id" json:"type_id"`
	Title       string    `db:"title" json:"title"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	IsFinished  bool      `db:"is_finished" json:"is_finished"`

	OrganizerName string `db:"organizer_name" json:"organizer_name"`
	EventTypeName string `db:"event_type_name" json:"event_type_name"`

	Properties `json:"event_properties"`
	Details    `json:"event_detail"`

	TeamsIDs     pq.StringArray `db:"teams_ids" json:"event_teams"`
	Technologies pq.StringArray `db:"technologies" json:"tech_stack"`
}
