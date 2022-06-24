package request

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/alexedwards/flow"
	"github.com/rbraddev/shift-rota/internal/validator"
)

func ReadIDParam(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(flow.Param(r.Context(), "id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func ReadString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

func ReadCSV(qs url.Values, key string, defaultValue []string) []string {
	csv := qs.Get(key)

	if csv == "" {
		return defaultValue
	}

	return strings.Split(csv, ",")
}

func ReadInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddFieldError(key, "must be an integer value")
		return defaultValue
	}

	return i
}
