package team

import (
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

// @Summary      Create team
// @Description  Create team
// @Tags         team
// @Accept       multipart/form-data
// @Produce      json
// @Param        name     formData  string  true   "Team name"
// @Param        team_status formData  string  false   "Team status"
// @Param        description    formData  string  false   "Team description"
// @Param        picture     formData  file    false  "Team picture"
// @Success      200
// @Security     CookieAuth
// @Failure      400          {object} wrapper.CustomError
// @Failure      401          {object} wrapper.CustomError
// @Failure      413          {object} wrapper.CustomError
// @Failure      500          {object} wrapper.CustomError
// @Router       /api/teams/create [post]
func (h *TeamHandler) Create(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	var input dto.TeamCreationRequest

	if err := w.Request().ParseMultipartForm(1 << 20); err != nil {
		return nil, domain.ApiError.RequestFileError(domain.ApiError{}, "1Mb", err)
	}

	setters := map[string]func(string){
		"name":        func(s string) { input.Name = s },
		"description": func(s string) { input.Description = s },
		"team_status": func(s string) { input.TeamStatus = s },
	}

	for k, set := range setters {
		if v := w.Request().FormValue(k); v != "" {
			set(v)
		}
	}

	file, headers, err := w.Request().FormFile("picture")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		logger.Error("error with getting form file value: ", err)
		return nil, err
	}
	hasPicture := err == nil

	input.TeamPictureURL = "/defaults/picture/team.jpg"

	if err := h.TeamUsecase.Create(w.Request().Context(), c.AccountID, &input); err != nil {
		if !errors.Is(err, domain.ErrTeamsNameTaken) {
			logger.Error("failed to create team")
			return nil, err
		}

		return nil, wrapper.NewError(err.Error(), http.StatusBadRequest)
	}

	if hasPicture {
		newURL := path.Join("teams", uuid.NewString()+headers.Filename)
		if err := s3.StorePictureAtS3(w.Request().Context(), file, headers, newURL); err != nil {
			logger.Error("failed to upload team picture: ", err)
		} else {
			if err := h.TeamUsecase.UpdatePictureByCreatorID(w.Request().Context(), newURL, c.AccountID); err != nil {
				return nil, err
			}
		}
	}

	return nil, nil
}
