package main

import (
	"net/http"

	"github.com/alexedwards/flow"
)

func (app *application) routes() http.Handler {
	mux := flow.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.Use(app.recoverPanic)

	mux.HandleFunc("/healthcheck", app.healthcheckHandler, "GET")

	mux.HandleFunc("/teams", app.createTeamHandler, "POST")
	mux.HandleFunc("/teams", app.listTeamsHandler, "GET")
	mux.HandleFunc("/teams/:id", app.getTeamHandler, "GET")
	mux.HandleFunc("/teams/:id", app.updateTeamHandler, "PATCH")
	mux.HandleFunc("/teams/:id", app.deleteTeamHandler, "DELETE")

	return mux
}
