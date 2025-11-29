package team

import (
	"context"

	teammodel "github.com/goawwer/devclash/internal/domain/team_model"
	"github.com/google/uuid"
)

type TeamRepository interface {
	Create(ctx context.Context, accountID uuid.UUID, t *teammodel.Team) error
}

type TeamUsecase struct {
	r TeamRepository
}

func NewTeamUsecase(repository TeamRepository) *TeamUsecase {
	return &TeamUsecase{
		r: repository,
	}
}
