package eventmodel

import (
	"time"

	"github.com/google/uuid"
)

type Details struct {
	EventID         uuid.UUID `db:"event_id" json:"event_id"`
	EventPictureURL string    `db:"event_picture_url" json:"event_picture_url"`
	Description     string    `db:"description" json:"description"`
	Prize           string    `db:"prize" json:"prize"`
	StartTime       time.Time `db:"start_time" json:"start_time"`
	EndTime         time.Time `db:"end_time" json:"end_time"`
}
