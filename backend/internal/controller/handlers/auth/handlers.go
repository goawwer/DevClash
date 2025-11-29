package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/domain"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/goawwer/devclash/pkg/s3"
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Sign Up User
// @Description  Creates a user in database
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      dto.SignUpForm  true "Signup credentials"
// @Success      200
// @Failure      400          {object} wrapper.CustomError
// @Failure      500          {object} wrapper.CustomError
// @Router       /auth/signup/user [post]
func (h *AuthHandler) SignUpUser(w *wrapper.Wrapper) error {
	var input dto.SignUpForm

	err := json.NewDecoder(w.Request().Body).Decode(&input)
	if err != nil {
		logger.Error("failed to decode request body: ", err)
		return wrapper.NewError("invalid type", http.StatusBadRequest)
	}

	if err := h.AuthUsecase.SignUpUser(w.Request().Context(), input); err != nil {
		logger.Error("failed to signup user: ", err)
		switch {
		case errors.Is(err, domain.ErrEmailTaken):
			return wrapper.NewError("user with the same email already exists", http.StatusBadRequest)
		case errors.Is(err, domain.ErrUsernameTaken):
			return wrapper.NewError("user with the same username already exists", http.StatusBadRequest)
		default:
			return err
		}
	}

	return nil
}

// @Summary      Register a new organizer
// @Description  Create a new organizer account with logo upload
// @Tags         auth
// @Accept       multipart/form-data
// @Produce      json
// @Param        email    formData  string  true   "Organizer email"
// @Param        name     formData  string  true   "Organizer name"
// @Param        password formData  string  true   "Password"
// @Param        color    formData  string  true   "Brand color"
// @Param        logo     formData  file    false  "Organizer logo"
// @Success      200
// @Failure      400          {object} wrapper.CustomError
// @Failure      500          {object} wrapper.CustomError
// @Router       /auth/signup/organizer [post]
func (h *AuthHandler) SignUpOrganizer(w *wrapper.Wrapper) error {
	var input dto.SignUpInputOrganizerDetails

	if err := w.Request().ParseMultipartForm(10 << 20); err != nil {
		logger.Error("failed to parse multipart form ")
		return wrapper.NewError("invalid type", http.StatusBadRequest)
	}

	setters := map[string]func(string){
		"email":    func(v string) { input.Email = v },
		"name":     func(v string) { input.Name = v },
		"password": func(v string) { input.Password = v },
		"color":    func(v string) { input.Color = v },
	}

	for k, set := range setters {
		if v := w.Request().FormValue(k); v != "" {
			set(v)
		}
	}

	file, headers, err := w.Request().FormFile("logo")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			logger.Error("error with getting form file value: ", err)
			return err
		} else {
			input.LogoURL = "/defaults/logos/default.jpg"
		}
	}

	if file != nil {
		input.LogoURL, err = s3.StorePictureAtS3(w.Request().Context(), file, headers, input.Name, "logos")
		if err != nil {
			return err
		}
	}

	if err := h.AuthUsecase.SignUpOrganizer(w.Request().Context(), input); err != nil {
		s3.Delete(&s3.S3RemoveFileParameters{
			Ctx:      w.Request().Context(),
			Filename: input.LogoURL,
		})
	}

	return nil
}

// @Summary      Login
// @Description  Authenticate user and return access + refresh tokenss
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      dto.LoginForm  true  "Login credentials"
// @Success      200
// @Failure      400          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Failure      500          {object} wrapper.CustomError
// @Router       /auth/login [post]
func (h *AuthHandler) Login(w *wrapper.Wrapper) error {
	var input dto.LoginForm

	err := json.NewDecoder(w.Request().Body).Decode(&input)
	if err != nil {
		logger.Error("failed to decode request body: ", err)
		return wrapper.NewError("invalid type", http.StatusBadRequest)
	}

	userID, role, err := h.AuthUsecase.Login(w.Request().Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			logger.Error("user not found")
			return wrapper.NewError("user not found", http.StatusBadRequest)
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			logger.Error("invalid credentials: ", err)
			return wrapper.NewError("invalid credentials", http.StatusBadRequest)
		default:
			logger.Error(err)
			return err
		}
	}

	tokenPair, err := middleware.GenerateTokenPair(w.Request().Context(), userID, role)
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
