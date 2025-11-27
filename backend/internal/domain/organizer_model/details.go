package organizermodel

import "github.com/google/uuid"

type Details struct {
	OrganizerID uuid.UUID `db:"organizer_id" json:"organizer_id"`
	Description *string   `db:"company_description" json:"company_description"`
	LogoURL     *string   `db:"logo_url" json:"logo_url"`
	Color       *string   `db:"brand_color" json:"brand_color"`
}
