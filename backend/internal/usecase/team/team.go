package team

import (
	"context"
	"time"

	teammodel "github.com/goawwer/devclash/internal/domain/team_model"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/google/uuid"
)

func (t *TeamUsecase) Create(ctx context.Context, accountID uuid.UUID, req *dto.TeamCreationRequest) error {
	return t.r.CreateTeam(ctx, accountID, &teammodel.Team{
		Name:           req.Name,
		TeamPictureURL: &req.TeamPictureURL,
		TeamStatus:     &req.TeamStatus,
		Description:    &req.Description,
		CreatedAt:      time.Now(),
	})
}

func (t *TeamUsecase) UpdatePictureByCreatorID(ctx context.Context, newURL string, accountID uuid.UUID) error {
	return t.r.UpdateTeamPictureByCreatorID(ctx, newURL, accountID)
}
