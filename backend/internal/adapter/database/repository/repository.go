package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	AuthRepository
	UserRepository
}

type ApplicationRepository struct {
	*sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &ApplicationRepository{db}
}
