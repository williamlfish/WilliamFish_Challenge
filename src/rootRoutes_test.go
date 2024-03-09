package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRootGetRoute(t *testing.T) {
	b, err := os.ReadFile(fmt.Sprintf("%s/%s", TestTemplatePathNoWild, "/index.html"))
	if err != nil {
		t.Fatalf("error getting template index.html: %s", err)
		return
	}
	router := buildRouter(TestTemplatePath)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	code := w.Code
	body := w.Body.String()
	if code != 200 {
		t.Errorf("expected statusCode: %d, got: %d ", 200, code)
	}
	if body != string(b) {
		t.Errorf("expected body to read: %s, got: %s ", "ping", body)
	}
}
