package dto

import (
	"time"
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
	Title           string    `db:"title" json:"title"`
	EventPictureURL string    `db:"event_picture_url" json:"event_picture_url"`
	Type            string    `db:"event_type_name" json:"event_type_name"`
	Properties      *[]string `db:"event_properties" json:"event_properties"`
	TechStack       *[]string `db:"tech_stack" json:"tech_stack"`
	Teams           *[]string `db:"teams" json:"teams"`
	StartDate       time.Time `db:"start_time" json:"start_time"`
	EndDate         time.Time `db:"end_time" json:"end_time"`
	Description     string    `db:"description" json:"description"`
}
