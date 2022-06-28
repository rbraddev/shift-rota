package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rbraddev/shift-rota/internal/data"
	"github.com/rbraddev/shift-rota/internal/request"
	"github.com/rbraddev/shift-rota/internal/response"
	"github.com/rbraddev/shift-rota/internal/validator"
)

func (app *application) createTeamHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	team := &data.Team{
		Name: input.Name,
	}

	v := validator.New()

	if data.ValidateTeam(v, team); v.HasErrors() {
		app.failedValidation(w, r, *v)
		return
	}

	err = app.models.Teams.Insert(team)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/teams/%d", team.ID))

	err = response.JSONWithHeaders(w, http.StatusCreated, envelope{"team": team}, headers)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getTeamHandler(w http.ResponseWriter, r *http.Request) {
	id, err := request.ReadIDParam(r)
	if err != nil {
		app.notFound(w, r)
		return
	}

	team, err := app.models.Teams.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFound(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	err = response.JSONWithHeaders(w, http.StatusOK, envelope{"team": team}, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) deleteTeamHandler(w http.ResponseWriter, r *http.Request) {
	id, err := request.ReadIDParam(r)
	if err != nil {
		app.notFound(w, r)
		return
	}

	err = app.models.Teams.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFound(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	err = response.JSONWithHeaders(w, http.StatusOK, envelope{"message": "team successfully deleted"}, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) listTeamsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Name = request.ReadString(qs, "title", "")

	input.Filters.Page = request.ReadInt(qs, "page", 1, v)
	input.Filters.PageSize = request.ReadInt(qs, "page_size", 20, v)

	input.Filters.Sort = request.ReadString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

	if data.ValidateFilters(v, input.Filters); v.HasErrors() {
		app.failedValidation(w, r, *v)
		return
	}

	teams, metadata, err := app.models.Teams.GetAll(input.Name, input.Filters)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSONWithHeaders(w, http.StatusOK, envelope{"teams": teams, "metadata": metadata}, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updateTeamHandler(w http.ResponseWriter, r *http.Request) {
	id, err := request.ReadIDParam(r)
	if err != nil {
		app.notFound(w, r)
		return
	}

	team, err := app.models.Teams.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFound(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	var input struct {
		Name *string `json:"name"`
	}

	err = request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	if input.Name != nil {
		team.Name = *input.Name
	}

	v := validator.New()

	if data.ValidateTeam(v, team); v.HasErrors() {
		app.failedValidation(w, r, *v)
		return
	}

	err = app.models.Teams.Update(team)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	err = response.JSONWithHeaders(w, http.StatusOK, envelope{"team": team}, nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
