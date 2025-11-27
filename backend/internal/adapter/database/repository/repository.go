package repository

import (
	"context"

	"github.com/goawwer/devclash/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	AuthRepository
	UserRepository
	OrganizerRepository
	AccountRepository
}

type ApplicationRepository struct {
	*sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &ApplicationRepository{db}
}

func (r *ApplicationRepository) RunInTransaction(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := r.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			if err := tx.Commit(); err != nil {
				logger.Error("transaction failed on commit: ", err)
			}
		}
	}()

	return fn(tx)
}
