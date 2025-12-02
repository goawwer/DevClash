package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type TechnologyRepository interface {
	GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error)
	GetTechnologyNamesByIDs(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]string, error)
}

type TechNameResult struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (r *ApplicationRepository) GetTechnologyIDByName(ctx context.Context, name string) (uuid.UUID, error) {
	var id uuid.UUID

	return id, r.GetContext(ctx, &id, `
		SELECT id FROM technologies
		WHERE name = $1
	`, name)
}

func (r *ApplicationRepository) GetTechnologyNamesByIDs(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]string, error) {
	if len(ids) == 0 {
		return make(map[uuid.UUID]string), nil
	}

	var results []TechNameResult

	query := `SELECT id, name FROM technologies WHERE id = ANY($1)`

	err := r.SelectContext(ctx, &results, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}

	idToNameMap := make(map[uuid.UUID]string, len(results))
	for _, res := range results {
		idToNameMap[res.ID] = res.Name
	}

	return idToNameMap, nil
}
