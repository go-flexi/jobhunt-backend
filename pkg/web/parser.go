package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-flexi/jobhunt-backend/pkg/errors"
)

var (
	ErrEmpty = fmt.Errorf("empty")
)

// ParseStrQueryParam parses a string query parameter from a http.Request
func ParseStrQueryParam(key string, r *http.Request) (string, bool) {
	val := r.URL.Query().Get(key)
	return val, val != ""
}

// ParseStrChiURLParams parses a string URL parameter from a http.Request
func ParseChiURLParams(key string, r *http.Request) (string, bool) {
	value := chi.URLParam(r, key)
	return value, value != ""
}

// ParseInt parses an integer from the URL param or the query param
func ParseInt(key string, r *http.Request, parseStrFunc func(key string, r *http.Request) (string, bool)) (int, error) {
	val, ok := parseStrFunc(key, r)
	if !ok {
		return 0, ErrEmpty
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.NewValidationsErr(errors.NewValidation(key, fmt.Errorf("value must be an integer")))
	}

	return i, nil
}
