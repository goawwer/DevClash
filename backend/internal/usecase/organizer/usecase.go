package organizer

import (
	"context"

	organizermodel "github.com/goawwer/devclash/internal/domain/organizer_model"
	"github.com/google/uuid"
)

type OrganizerRepository interface {
	GetOrganizerDetailsByID(ctx context.Context, orgID uuid.UUID) (*organizermodel.Details, error)
}

type OrganizerUsecase struct {
	r OrganizerRepository
}

func NewOrgUsecase(repository OrganizerRepository) *OrganizerUsecase {
	return &OrganizerUsecase{
		r: repository,
	}
}
