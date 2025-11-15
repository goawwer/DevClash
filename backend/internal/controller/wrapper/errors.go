package wrapper

import (
	"errors"
	"net/http"
)

type CustomError struct {
	Err        error
	StatusCode int
}

func (c *CustomError) Error() string {
	return c.Err.Error()
}

func (w *Wrapper) Error(err error) {
	var ce *CustomError

	if ok := errors.As(err, &ce); ok {
		w.JSONEncode(
			ce.StatusCode,
			map[string]string{"error": ce.Error()},
		)
	} else {
		w.JSONEncode(
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
	}
}

func NewError(msg string, status int) *CustomError {
	return &CustomError{
		Err:        errors.New(msg),
		StatusCode: status,
	}
}
