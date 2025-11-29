package usermodel

import "github.com/google/uuid"

type UserStack struct {
	UserID     uuid.UUID `db:"user_id" json:"-"`
	LanguageID uuid.UUID `db:"technology_id" json:"-"`
}
