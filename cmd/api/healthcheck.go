package main

import (
	"net/http"

	"github.com/rbraddev/shift-rota/internal/response"
	"github.com/rbraddev/shift-rota/internal/version"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version.Get(),
		},
	}

	err := response.JSONWithHeaders(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
