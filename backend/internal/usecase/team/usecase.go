package team

import (
	"context"
	"time"

	teammodel "github.com/goawwer/devclash/internal/domain/team_model"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/google/uuid"
)

func (t *TeamUsecase) Create(ctx context.Context, accountID uuid.UUID, req *dto.TeamCreationRequest) error {
	return t.r.Create(ctx, accountID, &teammodel.Team{
		Name:           req.Name,
		TeamPictureURL: &req.TeamPictureURL,
		TeamStatus:     &req.TeamStatus,
		Description:    &req.Description,
		CreatedAt:      time.Now(),
	})
}
