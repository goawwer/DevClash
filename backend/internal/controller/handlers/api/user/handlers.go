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
// @Router       /api/users/me/profile [get]
func (h *UserHandler) GetCurrentUserProfile(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	return h.GetUserProfileByID(w.Request().Context(), c.AccountID)
}

// @Summary      Profile settings
// @Description  Get current user profile settings
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200
// @Security     CookieAuth
// @Failure      500          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Router        /api/users/me/settings [get]
func (h *UserHandler) GetCurrentUserSettings(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	return h.GetUserSettingsByID(w.Request().Context(), c.AccountID)
}
