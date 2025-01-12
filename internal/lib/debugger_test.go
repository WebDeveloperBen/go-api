package lib_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/webdeveloperben/go-api/internal/lib"
)

func TestDebugLogger(t *testing.T) {
	// Set up a buffer to capture log output
	var logBuffer bytes.Buffer
	baseLogger := zerolog.New(&logBuffer).With().Timestamp().Logger()
	testLogger := lib.MyLogger{Logger: baseLogger} // Wrap the base logger in MyLogger
	lib.Logger = testLogger                        // Replace the global logger with the test logger

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/test-path", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Response().Header().Set(echo.HeaderXRequestID, "test-request-id")

	// Define the error message
	msg := "Test error message"

	// Call the DebugLogger function
	lib.DebugLogger(c, &msg)

	// Parse the log output
	logOutput := logBuffer.String()
	assert.Contains(t, logOutput, `"path":"/test-path"`, "Log should contain the request path")
	assert.Contains(t, logOutput, `"method":"GET"`, "Log should contain the request method")
	assert.Contains(t, logOutput, `"request_id":"test-request-id"`, "Log should contain the request ID")
	assert.Contains(t, logOutput, `"error":"Test error message"`, "Log should contain the error message")
	assert.Contains(t, logOutput, `"an error occurred"`, "Log should contain the main message")
}
