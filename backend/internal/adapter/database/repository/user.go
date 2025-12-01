package repository

import (
	"context"
	"errors"

	accountmodel "github.com/goawwer/devclash/internal/domain/account_model"
	usermodel "github.com/goawwer/devclash/internal/domain/user_model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrEmailTaken = errors.New("email already exists")
)

type UserRepository interface {
	CreateUser(ctx context.Context, a *accountmodel.Account, u *usermodel.User) error
	GetUserProfileByAccountID(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
	GetUserSettingsByAccountID(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
	UpdateCurrentUserProfileByID(ctx context.Context, a *accountmodel.Account, u *usermodel.User) error
	UpdateProfilePictureByAccountID(ctx context.Context, newURL string, accountID uuid.UUID) error
}

func (r *ApplicationRepository) CreateUser(ctx context.Context, a *accountmodel.Account, u *usermodel.User) error {
	return r.RunInTransaction(ctx, func(tx *sqlx.Tx) error {
		err := tx.QueryRowxContext(ctx, `
			INSERT INTO accounts (email, hashed_password, role, created_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`, a.Email, a.HashedPassword, a.Role, a.CreatedAt).Scan(&a.ID)
		if err != nil {
			return mapError(err)
		}

		_, err = tx.ExecContext(ctx, `
		INSERT INTO users (account_id, username)
			VALUES ($1, $2)
		`, a.ID, u.Username)

		return err
	})
}

func (r *ApplicationRepository) GetUserProfileByAccountID(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	var u usermodel.User

	err := r.GetContext(ctx, &u, `
		SELECT 
			u.username, 
			u.profile_picture_url, 
			u.bio, 
			u.profile_status,
			u.participations_count, 
			u.wins_count,
			COALESCE(
				ARRAY_AGG(us.technology_id::text) FILTER (WHERE us.technology_id IS NOT NULL),
				'{}'
			) AS technologies
		FROM users u
			LEFT JOIN users_skills us ON u.id = us.user_id
		WHERE u.account_id = $1
		GROUP BY u.id
	`, id)

	return &u, err
}

func (r *ApplicationRepository) GetUserSettingsByAccountID(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	var u usermodel.User

	err := r.GetContext(ctx, &u, `
		SELECT 
			a.email, 
			u.username, 
			u.profile_picture_url, 
			u.bio, 
			u.profile_status,
			COALESCE(
				ARRAY_AGG(us.technology_id::text) FILTER (WHERE us.technology_id IS NOT NULL),
				'{}'
			) AS technologies
		FROM accounts a
			JOIN users u ON a.id = u.account_id
			LEFT JOIN users_skills us ON u.id = us.user_id
		WHERE a.id = $1
		GROUP BY u.id, a.email
	`, id)

	return &u, err
}

func (r *ApplicationRepository) UpdateCurrentUserProfileByID(
	ctx context.Context, a *accountmodel.Account, u *usermodel.User,
) error {
	return r.RunInTransaction(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, `
			UPDATE accounts 
				SET email = $1, hashed_password = $2, updated_at = $3
			WHERE id = $4
		`, a.Email, a.HashedPassword, a.UpdatedAt, a.ID)
		if err != nil {
			return err
		}

		if err := tx.QueryRowxContext(ctx, `
			UPDATE users 
				SET username = $1, bio = $2, profile_status = $3
			WHERE account_id = $4
			RETURNING id
		`, u.Username, u.Bio, u.ProfileStatus, a.ID).Scan(&u.ID); err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx,
			`DELETE FROM users_skills WHERE user_id = $1`, u.ID)
		if err != nil {
			return err
		}

		for _, techID := range u.Technologies {
			_, err = tx.ExecContext(ctx, `
                INSERT INTO users_skills (user_id, technology_id)
                VALUES ($1, $2)
            `, u.ID, techID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *ApplicationRepository) UpdateProfilePictureByAccountID(ctx context.Context, newURL string, accountID uuid.UUID) error {
	_, err := r.ExecContext(ctx, `
		UPDATE users u
			SET profile_picture_url = $1
		FROM accounts a
		WHERE u.account_id = a.id
			  AND a.id = $2
	`, newURL, accountID)

	return err
}
