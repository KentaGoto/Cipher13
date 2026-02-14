package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRot13(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "lowercase", input: "hello", want: "uryyb"},
		{name: "uppercase", input: "HELLO", want: "URYYB"},
		{name: "mixed and symbols", input: "test-strings", want: "grfg-fgevatf"},
		{name: "empty", input: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rot13(tt.input); got != tt.want {
				t.Fatalf("rot13(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestHandleRot13(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := newRouter()

	req := httptest.NewRequest(http.MethodGet, "/api/rot13?s=test-strings", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusOK)
	}

	var body rot13Response
	if err := json.Unmarshal(res.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if body.Original != "test-strings" {
		t.Fatalf("original = %q, want %q", body.Original, "test-strings")
	}
	if body.Rot13 != "grfg-fgevatf" {
		t.Fatalf("rot13 = %q, want %q", body.Rot13, "grfg-fgevatf")
	}
}
