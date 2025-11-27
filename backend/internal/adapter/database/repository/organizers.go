package repository

import (
	"context"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	organizermodel "github.com/goawwer/devclash/internal/domain/organizer_model"
	"github.com/jmoiron/sqlx"
)

type OrganizerRepository interface {
	CreateOrganizer(ctx context.Context, a *accountmodel.Account, org *organizermodel.OrganizerAccount) error
}

func (r *ApplicationRepository) CreateOrganizer(ctx context.Context, a *accountmodel.Account, org *organizermodel.OrganizerAccount) error {
	return r.RunInTransaction(ctx, func(tx *sqlx.Tx) error {
		err := tx.QueryRowxContext(ctx, `
			INSERT INTO accounts (email, hashed_password, role, created_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`, a.Email, a.HashedPassword, a.Role, a.CreatedAt).Scan(&a.ID)
		if err != nil {
			return mapError(err)
		}

		err = tx.QueryRowxContext(ctx, `
			INSERT INTO organizers (name)
			VALUES ($1)
			RETURNING id
		`, org.Name).Scan(&org.ID)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, `
    		INSERT INTO organizers_details (organizer_id, logo_url, brand_color)
    		VALUES ($1, $2, $3)
		`, org.ID, org.Details.LogoURL, org.Details.Color)

		return err
	})

}
