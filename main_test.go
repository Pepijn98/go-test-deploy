package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	Index(w, req)

	res := w.Result()

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if string(data) != "Hello World!" {
		t.Errorf("Expected 'Hello World!', got '%s'", string(data))
	}
}

func TestGreet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet?name=Pepijn", nil)
	w := httptest.NewRecorder()

	Greet(w, req)

	res := w.Result()

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	if string(data) != "Hello Pepijn" {
		t.Errorf("Expected 'Hello Pepijn', got '%s'", string(data))
	}
}
