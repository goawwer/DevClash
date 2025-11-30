package team

import (
	"context"

	teammodel "github.com/goawwer/devclash/internal/domain/team_model"
	"github.com/google/uuid"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, accountID uuid.UUID, t *teammodel.Team) error
	UpdateTeamPictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error
}

type TeamUsecase struct {
	r TeamRepository
}

func NewTeamUsecase(repository TeamRepository) *TeamUsecase {
	return &TeamUsecase{
		r: repository,
	}
}
