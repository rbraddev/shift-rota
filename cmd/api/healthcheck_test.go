package main

import (
	"net/http"
	"testing"

	"github.com/rbraddev/shift-rota/internal/assert"
)

func TestHealthCheck(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/api/v1/healthcheck")

	assert.Equal(t, code, http.StatusOK)

	assert.Equal(t, body, `{
	"status": "available",
	"system_info": {
		"environment": "testing",
		"version": "unavailable"
	}
}
`)
}
