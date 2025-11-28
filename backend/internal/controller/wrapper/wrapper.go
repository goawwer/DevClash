package wrapper

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/goawwer/devclash/middleware"
)

type Wrapper struct {
	w http.ResponseWriter
	r *http.Request
}

func (w *Wrapper) Writer() http.ResponseWriter {
	return w.w
}

func (w *Wrapper) Request() *http.Request {
	return w.r
}

func (w *Wrapper) JSONEncode(statusCode int, v any) {
	w.Writer().WriteHeader(statusCode)
	json.NewEncoder(w.w).Encode(&v)
}

func (w *Wrapper) JSONDecode(v any) error {
	return json.NewDecoder(w.r.Body).Decode(v)
}

func (w *Wrapper) Param(key string) string {
	return chi.URLParam(w.r, key)
}

func (w *Wrapper) Query(key string) string {
	return w.r.URL.Query().Get(key)
}

func (w *Wrapper) claims() (*middleware.CustomClaims, error) {
	claims, ok := w.r.Context().Value(middleware.ClaimsKey).(*middleware.CustomClaims)
	if !ok {
		return nil, errors.New("missing claims")
	}

	return claims, nil
}
