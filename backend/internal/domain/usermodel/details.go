package usermodel

import "github.com/google/uuid"

type UserDetail struct {
	UserID        uuid.UUID `db:"user_id" json:"user_id"`
	ImageUrl      string    `db:"image_url" json:"image_url"`
	Bio           string    `db:"bio" json:"bio"`
	ProfileStatus string    `db:"profile_status" json:"profile_status"`
}
