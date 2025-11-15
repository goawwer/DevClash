package dto

import "github.com/google/uuid"

type SignUpInput struct {
	Email    string
	Username string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	ID uuid.UUID
}

type RefreshMeta struct {
	ID      uuid.UUID
	IsAdmin bool
}
