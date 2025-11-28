package wrapper

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/goawwer/devclash/middleware"
	"github.com/goawwer/devclash/pkg/logger"
)

type PublicHandler func(w *Wrapper) error

func PublicWrap(handler PublicHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		wrapper := &Wrapper{
			w: w,
			r: r,
		}

		err := handler(wrapper)
		if err != nil {
			wrapper.Error(err)
			return
		}

		w.WriteHeader(http.StatusOK)

	}
}

type AuthHandler func(w *Wrapper, c *middleware.CustomClaims) (any, error)

func AuthWrap(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		wrapper := &Wrapper{
			w: w,
			r: r,
		}

		claims, err := wrapper.claims()
		if err != nil {
			logger.Error("failed to get claims: ", err)
			wrapper.Error(NewError("invalid or missing token", http.StatusUnauthorized))
			return
		}

		res, err := handler(wrapper, claims)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				wrapper.Error(NewError("not found", http.StatusBadRequest))
				return
			}

			wrapper.Error(err)
			return
		}

		wrapper.JSONEncode(http.StatusOK, res)
	}
}
