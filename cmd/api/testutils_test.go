package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rbraddev/shift-rota/internal/data/mocks"
	"github.com/rbraddev/shift-rota/internal/log"
)

func newTestApplication(t *testing.T) *application {
	logger := log.New(io.Discard, log.LevelAll)
	models := mocks.NewMockModels()
	config := config{
		port:    4000,
		env:     "testing",
		version: false,
	}
	return &application{
		config: config,
		models: models,
		logger: logger,
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

func (ts *testServer) delete(t *testing.T, urlPath string) (int, http.Header, string) {
	req, err := http.NewRequest("DELETE", ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}
