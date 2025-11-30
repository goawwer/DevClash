package organizer

import (
	"context"

	organizermodel "github.com/goawwer/devclash/internal/domain/organizer_model"
	"github.com/google/uuid"
)

func (o *OrganizerUsecase) GetOrganizerDetailsByID(ctx context.Context, orgID uuid.UUID) (*organizermodel.Details, error) {
	return o.r.GetOrganizerDetailsByID(ctx, orgID)
}
