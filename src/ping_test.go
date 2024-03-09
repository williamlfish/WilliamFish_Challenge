package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := buildRouter(TestTemplatePath)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	code := w.Code
	body := w.Body.String()
	if w.Code != 200 {
		t.Errorf("expected statusCode: %d, got: %d ", 200, code)
	}
	if body != "{\"message\":\"pong\"}" {
		t.Errorf("expected body to read: %s, got: %s ", "{\"message\":\"pong\"}", body)
	}
}
