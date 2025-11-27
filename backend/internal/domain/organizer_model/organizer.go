package organizermodel

import (
	"github.com/google/uuid"
)

type OrganizerAccount struct {
	ID        uuid.UUID `db:"id" json:"id"`
	AccountID uuid.UUID `db:"account_id" json:"account_id"`
	Name      string    `db:"name" json:"organizer_name"`
	Details   *Details  `json:"organizer_details"`
}
