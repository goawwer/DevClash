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
