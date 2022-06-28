package main

import (
	"fmt"
	"net/http"

	"github.com/alexedwards/flow"
)

func (app *application) routes() http.Handler {
	mux := flow.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.Use(app.recoverPanic)

	prefix := "/api/v1"

	mux.HandleFunc(fmt.Sprintf("%s/healthcheck", prefix), app.healthcheckHandler, "GET")

	mux.HandleFunc(fmt.Sprintf("%s/teams", prefix), app.createTeamHandler, "POST")
	mux.HandleFunc(fmt.Sprintf("%s/teams", prefix), app.listTeamsHandler, "GET")
	mux.HandleFunc(fmt.Sprintf("%s/teams/:id", prefix), app.getTeamHandler, "GET")
	mux.HandleFunc(fmt.Sprintf("%s/teams/:id", prefix), app.updateTeamHandler, "PATCH")
	mux.HandleFunc(fmt.Sprintf("%s/teams/:id", prefix), app.deleteTeamHandler, "DELETE")

	return mux
}
