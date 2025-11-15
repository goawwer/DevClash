package user

import (
	"net/http"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/middleware"
)

// @Summary      Authorized
// @Description  Check current state of authorization
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200
// @Security     CookieAuth
// @Failure      401          {object} wrapper.CustomError
// @Router       /api/users/check [get]
func (h *UserHandler) Check(_ *wrapper.Wrapper, _ *middleware.CustomClaims) (any, error) {
	return nil, nil
}

// @Summary      Logout
// @Description  Refreshes TokenPair and check refresh token is consumed or not
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200
// @Security     CookieAuth
// @Failure      401          {object} wrapper.CustomError
// @Router       /api/users/logout [post]
func (h *UserHandler) Logout(w *wrapper.Wrapper, _ *middleware.CustomClaims) (any, error) {
	http.SetCookie(w.Writer(), &http.Cookie{
		Name:     "access",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	http.SetCookie(w.Writer(), &http.Cookie{
		Name:     "refresh",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	return nil, nil
}
