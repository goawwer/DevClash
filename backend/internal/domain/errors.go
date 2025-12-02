package domain

import (
	"errors"
	"net/http"
	"strings"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/pkg/logger"
)

type ApiError struct{}

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

// User Profile
var (
	ErrInvalidOldPassword = errors.New("invalid old password, please try again")
)

// Permission
var (
	ErrNotForUserRole = errors.New("sorry, user cannot make an event")
)

// Events
var (
	ErrNotLeaderJoining            = errors.New("only the team leader can join the event")
	ErrTeamAlreadyJoined           = errors.New("team is already registered for this event")
	ErrTeamCountGreaterThanAllowed = errors.New("event is full, maximum teams reached")
	ErrTeamMembersCountIsNotValid  = errors.New("current team members count is not valid for event team members count properties")
)

func (a ApiError) RequestFileError(maxSise string, err error) error {
	if strings.Contains(err.Error(), "request body too large") {
		return wrapper.NewError(
			"picture is too large, maximum allowed size is"+maxSise,
			http.StatusRequestEntityTooLarge,
		)
	}

	logger.Error("failed to parse multipart form: ", err)
	return wrapper.NewError(
		"invalid file upload format",
		http.StatusBadRequest,
	)
}
