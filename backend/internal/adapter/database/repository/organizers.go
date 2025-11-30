package repository

import (
	"context"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	organizermodel "github.com/goawwer/devclash/internal/domain/organizer_model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OrganizerRepository interface {
	CreateOrganizer(ctx context.Context, a *accountmodel.Account, org *organizermodel.OrganizerAccount) error
	GetOrganizerDetailsByID(ctx context.Context, orgID uuid.UUID) (*organizermodel.Details, error)
	GetOrganizerIDByAccountID(ctx context.Context, accountID uuid.UUID) (uuid.UUID, error)
	UpdateLogoByCreatorID(ctx context.Context, accountID uuid.UUID, newURL string) error
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
			INSERT INTO organizers (account_id, name)
			VALUES ($1, $2)
			RETURNING id
		`, a.ID, org.Name).Scan(&org.ID)
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

func (r *ApplicationRepository) GetOrganizerDetailsByID(ctx context.Context, orgID uuid.UUID) (*organizermodel.Details, error) {
	var details organizermodel.Details

	return &details, r.GetContext(ctx, &details, `
		SELECT org_d.company_description, org_d.logo_url, org_d.brand_color FROM organizers_details org_d
		WHERE organizer_id = $1 
	`, orgID)
}

func (r *ApplicationRepository) GetOrganizerIDByAccountID(ctx context.Context, accountID uuid.UUID) (uuid.UUID, error) {
	var id uuid.UUID

	return id, r.GetContext(ctx, &id, `
		SELECT id FROM organizers 
		WHERE account_id = $1
	`, accountID)
}

func (r *ApplicationRepository) UpdateLogoByCreatorID(ctx context.Context, organizerID uuid.UUID, newURL string) error {
	_, err := r.ExecContext(ctx, `
		UPDATE organizers_details od
		SET logo_url = $1
		FROM organizers o
		WHERE od.organizer_id = o.id
		  AND o.id = $2
	`, newURL, organizerID)

	return err
}
