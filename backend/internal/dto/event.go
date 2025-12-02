package dto

import (
	"time"

	"github.com/google/uuid"
)

type EventCreationRequest struct {
	Title           string    `json:"title"`
	EventPictureURL string    `json:"-"`
	Type            string    `json:"type"`
	IsOnline        bool      `json:"is_online"`
	IsFree          bool      `json:"is_free"`
	NumberOfTeams   int       `json:"number_of_teams"`
	TeamSize        int       `json:"team_size"`
	TechStack       []string  `json:"tech_stack"`
	Description     string    `json:"description"`
	Prize           string    `json:"prize"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
}

type EventPage struct {
	Title           string   `json:"title"`
	EventPictureURL string   `json:"event_picture_url"`
	Type            string   `json:"event_type_name"`
	IsOnline        bool     `json:"is_online"`
	IsFree          bool     `json:"is_free"`
	NumberOfTeams   int      `json:"number_of_teams"`
	TeamSize        int      `json:"team_size"`
	TechStack       []string `json:"tech_stack"`

	TeamName       string `json:"team_name"`
	TeamPictureURL string `json:"team_picture_url"`
	TeamStatus     string `json:"team_status"`

	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Description string    `json:"description"`
}

type EventListResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`

	OrganizerName string `json:"organizer_name"`
	EventTypeName string `json:"event_type_name"`

	IsOnline bool `json:"is_online"`
	IsFree   bool `json:"is_free"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	TechStack []string `json:"tech_stack"`

	Description string `json:"description_preview,omitempty"`
}

type TeamJoinRequestToEvent struct {
	EventID uuid.UUID `json:"event_id"`
	TeamID  uuid.UUID `json:"team_id"`
}
