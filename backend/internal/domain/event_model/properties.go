package eventmodel

import "github.com/google/uuid"

type Properties struct {
	EventID       uuid.UUID `db:"event_id" json:"event_id"`
	IsOnline      bool      `db:"is_online" json:"is_online"`
	IsFree        bool      `db:"is_free" json:"is_free"`
	NumberOfTeams int       `db:"number_of_teams" json:"number_of_teams"`
	TeamSize      int       `db:"team_size" json:"team_size"`
}
