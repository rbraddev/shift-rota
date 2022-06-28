package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/rbraddev/shift-rota/internal/assert"
)

func TestCreateTeamsHandler(t *testing.T) {
	var TestCases = []struct {
		name     string
		id       int64
		wantCode int
		wantBody string
	}{
		{
			name:     "recordExists",
			id:       1,
			wantCode: http.StatusOK,
			wantBody: `{
	"team": {
		"id": 1,
		"name": "Network",
		"version": 1
	}
}
`,
		}, {
			name:     "recordDoesNotExist",
			id:       2,
			wantCode: http.StatusNotFound,
			wantBody: `{
	"Error": "The requested resource could not be found"
}
`,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			app := newTestApplication(t)
			ts := newTestServer(t, app.routes())
			defer ts.Close()

			code, _, body := ts.get(t, fmt.Sprintf("/api/v1/teams/%d", tc.id))

			assert.Equal(t, code, tc.wantCode)
			assert.Equal(t, body, tc.wantBody)

		})
	}
}

func TestDeleteTeamHandler(t *testing.T) {
	t.Run("deleteTeam", func(t *testing.T) {
		app := newTestApplication(t)
		ts := newTestServer(t, app.routes())
		defer ts.Close()

		code, _, body := ts.delete(t, fmt.Sprintf("/api/v1/teams/%d", 1))

		fmt.Println(code)
		fmt.Println(body)

	})
}
