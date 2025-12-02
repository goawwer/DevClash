package repository

import (
	"context"
	"errors"
	"time"

	teammodel "github.com/goawwer/devclash/internal/domain/team_model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var (
	ErrTeamNameTaken = errors.New("team with that name already exists")
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, accountID uuid.UUID, t *teammodel.Team) error
	UpdateTeamPictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error
	GetTeamsByIDs(ctx context.Context, ids []uuid.UUID) ([]teammodel.Team, error)
}

func (r *ApplicationRepository) CreateTeam(ctx context.Context, accountID uuid.UUID, t *teammodel.Team) error {
	return r.RunInTransaction(ctx, func(tx *sqlx.Tx) error {
		err := tx.GetContext(ctx, &t.CreatorID, `
			SELECT id FROM users 
			WHERE account_id = $1
		`, &accountID)
		if err != nil {
			return err
		}

		t.LeaderID = t.CreatorID

		if err := tx.QueryRowxContext(ctx, `
			INSERT INTO teams (creator_id, leader_id, name, team_status, description, team_picture_url, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			RETURNING id
		`, t.CreatorID, t.LeaderID, t.Name, t.TeamStatus, t.Description, t.TeamPictureURL, t.CreatedAt).Scan(&t.ID); err != nil {
			return mapError(err)
		}

		var leadRoleId uuid.UUID
		if err = tx.GetContext(ctx, &leadRoleId, `
			SELECT id FROM teams_roles
			WHERE name = 'Lead' 
		`); err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, `
			INSERT INTO teams_members (user_id, team_id, role, joined_at)
			VALUES ($1, $2, $3, $4)
		`, t.CreatorID, t.ID, &leadRoleId, time.Now())

		return err
	})
}

func (r *ApplicationRepository) UpdateTeamPictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error {
	_, err := r.ExecContext(ctx, `
		UPDATE teams t
        SET team_picture_url = $1
        FROM users u
        WHERE t.creator_id = u.id
          AND u.account_id = $2
	`, newURL, accountID)

	return err
}

func (r *ApplicationRepository) GetTeamsByIDs(ctx context.Context, ids []uuid.UUID) ([]teammodel.Team, error) {
	var teams []teammodel.Team

	return teams, r.SelectContext(ctx, &teams, ` 
        SELECT id, name, team_status, team_picture_url FROM teams
        WHERE id = ANY($1)
    `, pq.Array(ids))
}
