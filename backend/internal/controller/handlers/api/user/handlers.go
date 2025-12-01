package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"path"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/domain"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/goawwer/devclash/pkg/s3"
	"github.com/google/uuid"
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

// @Summary      Update user
// @Description  Update current user profile
// @Tags         user
// @Param        payload     formData  dto.EventCreationRequest  true   "Event preoperties"
// @Param        picture     formData  file    true  "Event picture"
// @Success      200
// @Security     CookieAuth
// @Failure      400          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Failure      413          {object} wrapper.CustomError
// @Failure      500          {object} wrapper.CustomError
// @Router        /api/users/me/update [put]
func (h *UserHandler) UpdateUserProfile(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	var input dto.UserUpdateProfileSettings

	if err := w.Request().ParseMultipartForm(3 << 20); err != nil {
		return nil, domain.ApiError.RequestFileError(domain.ApiError{}, "3 Mb", err)
	}

	payloadRaw := w.Request().FormValue("payload")
	if payloadRaw == "" {
		return nil, wrapper.NewError("missing payload", http.StatusBadRequest)
	}

	if err := json.Unmarshal([]byte(payloadRaw), &input); err != nil {
		return nil, wrapper.NewError("invalid payload", http.StatusBadRequest)
	}

	file, headers, err := w.Request().FormFile("picture")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		logger.Error("error with getting form file value: ", err)
		return nil, err
	}

	hasPicture := err == nil

	if err := h.UserUsecase.UpdateCurrentUserProfileByID(w.Request().Context(), c.AccountID, &input); err != nil {
		logger.Error("failed to update current user profile: ", err)
		return nil, err
	}

	if hasPicture {
		newURL := path.Join("users", uuid.NewString()+headers.Filename)

		if err := s3.StorePictureAtS3(w.Request().Context(), file, headers, newURL); err != nil {
			logger.Error("failed to upload user picture: ", err)
		} else {
			if *input.ProfilePictureURL != "" {
				if err := s3.Delete(&s3.S3RemoveFileParameters{
					Ctx:      w.Request().Context(),
					Filename: *input.ProfilePictureURL,
				}); err != nil {
					logger.Error("failed delete picture from s3: ", err)
					return nil, err
				}
			}

			if err := h.UserUsecase.UpdatePictureByID(w.Request().Context(), newURL, c.AccountID); err != nil {
				logger.Error("failed to update picture url in database after uploading to s3: ", err)
				return nil, err
			}
		}
	}

	return nil, nil
}
