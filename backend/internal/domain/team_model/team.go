package teammodel

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID                  uuid.UUID `db:"id" json:"id"`
	CreatorID           uuid.UUID `db:"creator_id" json:"creator_id"`
	LeaderID            uuid.UUID `db:"leader_id" json:"leader_id"`
	Name                string    `db:"name" json:"name"`
	TeamStatus          *string   `db:"team_status" json:"team_status"`
	Description         *string   `db:"desciption" json:"description"`
	TeamPictureURL      *string   `db:"team_picture_url" json:"team_picture_url"`
	ParticipationsCount int       `db:"participations_count" json:"participations_count"`
	WinsCount           int       `db:"wins_count" json:"wins_count"`
	CreatedAt           time.Time `db:"created_at" json:"created_at"`
	Disabled            bool      `db:"disabled" json:"disabled"`
}

type JoinTeamsToEventValidationData struct {
	MaxTeams    int `db:"number_of_teams"`
	MaxTeamSize int `db:"team_size"`

	CurrentTeamCount       int `db:"current_team_count"`
	CurrentTeamMemberCount int `db:"current_member_count"`

	TeamLeaderID uuid.UUID `db:"leader_id"`

	IsTeamAlreadyJoined bool `db:"is_joined"`
}
