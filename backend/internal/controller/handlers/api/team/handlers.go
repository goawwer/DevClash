package team

import (
	"errors"
	"net/http"

	"github.com/goawwer/devclash/internal/controller/wrapper"
	"github.com/goawwer/devclash/internal/dto"
	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
	"github.com/goawwer/devclash/pkg/s3"
)

func (h *TeamHandler) Create(w *wrapper.Wrapper, c *middleware.CustomClaims) (any, error) {
	var input dto.TeamCreationRequest

	if err := w.Request().ParseMultipartForm(10 << 20); err != nil {
		logger.Error("failed to parse multipart form ")
		return nil, wrapper.NewError("invalid type", http.StatusBadRequest)
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
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			logger.Error("error with getting form file value: ", err)
			return nil, err
		} else {
			input.TeamPictureURL = "/defaults/picture/team.jpg"
		}
	}

	if file != nil {
		input.TeamPictureURL, err = s3.StorePictureAtS3(w.Request().Context(), file, headers, input.Name, "teams")
		if err != nil {
			return nil, err
		}
	}

	if err := h.TeamUsecase.Create(w.Request().Context(), c.AccountID, &input); err != nil {
		s3.Delete(&s3.S3RemoveFileParameters{
			Ctx:      w.Request().Context(),
			Filename: input.TeamPictureURL,
		})

		return nil, err
	}

	return nil, nil
}
