package domain

import "errors"

// Database
var (
	// Duplicates
	ErrEmailTaken         = errors.New("email already taken")
	ErrUsernameTaken      = errors.New("username already taken")
	ErrOrganizerNameTaken = errors.New("organizer with the same company name already exists")
	ErrTeamsNameTaken     = errors.New("team with the same name already exists, please make up another")

	// Violations
	ErrForeignKeyViolation = errors.New("referenced entity does not exist")
	ErrNotNullViolation    = errors.New("missing required fields")
	ErrCheckViolation      = errors.New("data does not satisfy required constraints")
)

// Signup
var (
	ErrRequiredFields     = errors.New("these fields required and cannot be empty")
	ErrInvalidEmail       = errors.New("email is invalid")
	ErrInvalidUsername    = errors.New("username must contains only letters and numbers and be at least 5 characters")
	ErrInvalidCompanyName = errors.New("name must contains only letters")
)
