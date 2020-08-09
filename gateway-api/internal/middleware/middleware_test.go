package middleware_test

import (
	"bytes"
	"github.com/bgoldovsky/dutyer/gateway-api/internal/logger"
	."github.com/bgoldovsky/dutyer/gateway-api/internal/middleware"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestPanicMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic!")
	})

	handlerToTest := PanicMiddleware(nextHandler)

	req := httptest.NewRequest("GET", "http://localhost", nil)
	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}

func TestLogMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	handlerToTest := LogMiddleware(nextHandler)

	var buf bytes.Buffer
	logger.Log.SetOutput(&buf)

	req := httptest.NewRequest("GET", "http://localhost/abc", nil)
	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)

	template := "/abc"
	if !strings.Contains(buf.String(), template) {
		t.Errorf("log not contains %q", template)
	}

	logger.Log.SetOutput(os.Stderr)
}
