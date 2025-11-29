package organizer

import (
	"net/http"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/google/uuid"
)

// @Summary      Organizer Profile
// @Description  Get current user profile settings
// @Tags         organizer
// @Accept       json
// @Produce      json
// @Success      200
// @Security     CookieAuth
// @Failure      500          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Router       /api/organizers/{id} [get]
func (h *OrganizerHandler) GetOrganizerByID(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	idStr := w.Param("id")

	uuid, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("failed to parse uuid: ", err)
		return nil, wrapper.NewError("invalid id type", http.StatusBadRequest)
	}

	return h.OrganizerUsecase.GetOrganizerDetailsByID(w.Request().Context(), uuid)
}
