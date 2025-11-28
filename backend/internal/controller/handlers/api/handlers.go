package api

import (
	"net/http"
	"time"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/s3"
)

// @Summary      Logout
// @Description  Refreshes TokenPair and check refresh token is consumed or not
// @Tags         api
// @Accept       json
// @Produce      json
// @Success      200
// @Security     CookieAuth
// @Failure      401          {object} wrapper.CustomError
// @Router       /api/logout [post]
func Logout(w *wrapper.Wrapper, _ *middleware.CustomClaims) (any, error) {
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

// @Summary      Get S3 Image URL
// @Description  Generates a presigned URL for accessing an image stored in S3
// @Tags         api
// @Accept       json
// @Produce      json
// @Success      200
// @Security     CookieAuth
// @Failure      500          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Router       /api/logout [post]
func GetS3Url(w *wrapper.Wrapper, _ *middleware.CustomClaims) (any, error) {
	return s3.PresignKey(&s3.S3GetFileParameters{
		Ctx:      w.Request().Context(),
		FileName: w.Query("file"),
		Expires:  time.Duration(30) * time.Minute,
	})
}
