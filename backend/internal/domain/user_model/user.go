package usermodel

import (
	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID    `db:"id" json:"id"`
	AccountID           uuid.UUID    `db:"account_id" json:"account_id"`
	Username            string       `db:"username" json:"username"`
	ProfilePictureURL   *string      `db:"profile_picture_url" json:"profile_picture_url"`
	Bio                 *string      `db:"bio" json:"bio"`
	ProfileStatus       *string      `db:"profile_status" json:"profile_status"`
	ParticipationsCount int          `db:"participations_count" json:"participations_count"`
	WinsCount           int          `db:"wins_count" json:"wins_count"`
	Stack               *[]UserStack `db:"-" json:"-"`
}
