package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/rbraddev/shift-rota/internal/assert"
)

func TestCreateTeamsHandler(t *testing.T) {
	var TestCases = []struct {
		name    string
		id      int64
		expBody string
		expErr  error
	}{
		{
			name: "recordExists",
			id:   1,
			expBody: `{
	"team": {
		"id": 1,
		"name": "Network",
		"version": 1
	}
}
`,
			expErr: nil,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t *testing.T) {
			app := newTestApplication(t)
			ts := newTestServer(t, app.routes())
			defer ts.Close()

			code, _, body := ts.get(t, fmt.Sprintf("/api/v1/teams/%d", tc.id))

			assert.Equal(t, code, http.StatusOK)

			if tc.expBody != "" {
				assert.Equal(t, body, tc.expBody)
			}

			// if tc.expErr != nil {

			// }
		})
	}
}
