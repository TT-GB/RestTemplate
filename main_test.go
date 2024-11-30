package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Create a new request to the /api/hello endpoint
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Create a handler and serve the request
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", status)
	}

	// Check the response body
	expected := `{"message":"Hello, World!"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %v, got %v", expected, rr.Body.String())
	}
}
