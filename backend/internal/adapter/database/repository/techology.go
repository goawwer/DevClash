package repository

import (
	"context"

	"github.com/google/uuid"
)

type TechnologyRepository interface {
	GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error)
}

func (r *ApplicationRepository) GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error) {
	var id uuid.UUID

	return id, r.GetContext(ctx, &id, `
		SELECT id FROM technologies
		WHERE name = $1
	`, name)
}
