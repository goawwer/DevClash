package usermodel

import "github.com/google/uuid"

type UserSkill struct {
	UserID  uuid.UUID `db:"user_id" json:"user_id"`
	StackID uuid.UUID `db:"StackID" json:"stack_id"`
}
