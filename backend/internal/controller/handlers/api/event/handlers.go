package event

import (
	"encoding/json"
	"errors"
	"net/http"
	"path"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/domain"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/helpers"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/goawwer/devclash/pkg/s3"
	"github.com/google/uuid"
)

// @Summary      Create event
// @Description  Create event (users can't do it)
// @Tags         event
// @Accept       multipart/form-data
// @Produce      json
// @Param        payload     formData  dto.EventCreationRequest  true   "Event preoperties"
// @Param        picture     formData  file    true  "Event picture"
// @Success      200
// @Security     CookieAuth
// @Failure      400          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Failure      413          {object} wrapper.CustomError
// @Failure      500          {object} wrapper.CustomError
// @Router       /api/events/create [post]
func (h *EventHandler) Create(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	var input dto.EventCreationRequest

	if err := w.Request().ParseMultipartForm(10 << 20); err != nil {
		return nil, domain.ApiError.RequestFileError(domain.ApiError{}, "10 Mb", err)
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

	input.EventPictureURL = "/defaults/picture/event.jpg"

	if err := h.EventUsecase.Create(w.Request().Context(), c.AccountID, c.Role, &input); err != nil {
		if !errors.Is(err, domain.ErrNotForUserRole) {
			logger.Error("failed to create event")
			return nil, err
		}

		return nil, wrapper.NewError(err.Error(), http.StatusBadRequest)
	}

	if hasPicture {
		newURL := path.Join("events", uuid.NewString()+headers.Filename)
		if err := s3.StorePictureAtS3(w.Request().Context(), file, headers, newURL); err != nil {
			logger.Error("failed to upload team picture: ", err)
		} else {
			if err := h.EventUsecase.UpdatePictureByCreatorID(w.Request().Context(), newURL, c.AccountID); err != nil {
				return nil, err
			}
		}
	}

	return nil, nil
}

func (h *EventHandler) GetEventPage(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	id := w.Param("id")
	uuidID := uuid.MustParse(id)

	return h.EventUsecase.GetEventPageByID(w.Request().Context(), uuidID)
}

func (h *EventHandler) GetAllEvents(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	params := helpers.GetQueryWithFilterParameters(w.Request())
	return h.EventUsecase.GetAllEvents(w.Request().Context(), params)
}
