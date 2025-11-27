package user

import (
	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/middleware"
)

// @Summary      Profile
// @Description  Get current user profile
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200
// @Security     CookieAuth
// @Failure      500          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Router       /api/users/current/profile [post]
func (h *UserHandler) GetCurrentUserProfile(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	u, err := h.GetUserProfileByID(w.Request().Context(), c.AccountID)

	return UserProfile{
		Username:            u.Username,
		ProfilePictureURL:   u.ProfilePictureURL,
		Bio:                 u.Bio,
		ProfileStatus:       u.ProfileStatus,
		ParticipationsCount: u.ParticipationsCount,
		WinsCount:           u.WinsCount,
	}, err
}
