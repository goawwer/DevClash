package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
)

// @Summary      Sign Up
// @Description  Creates a user in database
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      dto.SignUpInput  true  "Login credentials"
// @Success      200
// @Failure      400          {object} wrapper.CustomError
// @Failure      500          {object} wrapper.CustomError
// @Router       /auth/signup [post]
func (h *AuthHandler) SignUp(w *wrapper.Wrapper) error {
	var input dto.SignUpInput

	err := json.NewDecoder(w.Request().Body).Decode(&input)
	if err != nil {
		logger.Error("failed to decode request body: ", err)
		return wrapper.NewError("invalid type", http.StatusBadRequest)
	}

	if err := h.AuthUsecase.SignUp(w.Request().Context(), input); err != nil {
		logger.Error("failed to signup user", err)
		return err
	}

	return nil
}

// @Summary      Login
// @Description  Authenticate user and return access + refresh tokenss
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      dto.LoginInput  true  "Login credentials"
// @Success      200
// @Failure      400          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Router       /auth/login [post]
func (h *AuthHandler) Login(w *wrapper.Wrapper) error {
	var input dto.LoginInput

	err := json.NewDecoder(w.Request().Body).Decode(&input)
	if err != nil {
		logger.Error("failed to decode request body: ", err)
		return wrapper.NewError("invalid type", http.StatusBadRequest)
	}

	userID, isAdmin, err := h.AuthUsecase.Login(w.Request().Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			logger.Error("user not found")
			return wrapper.NewError("user not found", http.StatusBadRequest)
		default:
			logger.Error(err)
			return err
		}
	}

	tokenPair, err := middleware.GenerateTokenPair(w.Request().Context(), userID, isAdmin)
	if err != nil {
		logger.Error("login token pair generation: ", err)
		return err
	}

	middleware.SetAuthCookie(w.Writer(), middleware.CustomCookie{
		Name:  "access",
		Value: tokenPair.AccessToken,
		Exp:   time.Until(tokenPair.AccessTokenExp),
	})

	middleware.SetAuthCookie(w.Writer(), middleware.CustomCookie{
		Name:  "refresh",
		Value: tokenPair.RefreshToken,
		Exp:   time.Until(tokenPair.RefreshTokenExp),
	})

	return nil
}

// @Summary      Refresh
// @Description  Refreshes TokenPair and check refresh token is consumed or not
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      401          {object} wrapper.CustomError
// @Router       /auth/refresh [post]
func (h *AuthHandler) Refresh(w *wrapper.Wrapper) error {
	tokenPair, err := middleware.RefreshToken(w.Request())
	if err != nil {
		logger.Error("refresh token handler: ", err)
		return wrapper.NewError("Unauthorized", http.StatusUnauthorized)
	}

	middleware.SetAuthCookie(w.Writer(), middleware.CustomCookie{
		Name:  "access",
		Value: tokenPair.AccessToken,
		Exp:   time.Until(tokenPair.AccessTokenExp),
	})

	middleware.SetAuthCookie(w.Writer(), middleware.CustomCookie{
		Name:  "refresh",
		Value: tokenPair.RefreshToken,
		Exp:   time.Until(tokenPair.RefreshTokenExp),
	})

	return nil
}
