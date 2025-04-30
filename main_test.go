package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetResourcesHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/resources", nil)
	w := httptest.NewRecorder()

	getResourcesHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestGetResourcesWithPostHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/resources", nil)
	w := httptest.NewRecorder()

	getResourcesHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", resp.StatusCode)
	}
}
