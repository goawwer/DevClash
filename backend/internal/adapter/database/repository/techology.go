package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type TechnologyRepository interface {
	GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error)
	GetTechnologyNamesByIDs(ctx context.Context, ids []uuid.UUID) ([]string, error)
}

func (r *ApplicationRepository) GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error) {
	var id uuid.UUID

	return id, r.GetContext(ctx, &id, `
		SELECT id FROM technologies
		WHERE name = $1
	`, name)
}

func (r *ApplicationRepository) GetTechnologyNamesByIDs(ctx context.Context, ids []uuid.UUID) ([]string, error) {
	var names []string
	return names, r.SelectContext(ctx, &names, `
		SELECT name FROM technologies WHERE id = ANY($1)
	`, pq.Array(ids))
}
