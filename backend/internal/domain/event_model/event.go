package eventmodel

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `db:"id" json:"id"`
	OrganizerID uuid.UUID `db:"organizer_id" json:"organizer_id"`
	TypeID      uuid.UUID `db:"type_id" json:"type_id"`
	Title       string    `db:"title" json:"title"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	IsFinished  bool      `db:"is_finished" json:"is_finished"`

	Properties   *Properties `json:"event_properties"`
	Details      *Details    `json:"event_detail"`
	Technologies []uuid.UUID `json:"tech_stack"`
}
