package lib

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

// MyLogger wraps zerolog.Logger for custom behavior
type MyLogger struct {
	zerolog.Logger
}

// Global logger instance
var Logger MyLogger

// NewLogger initializes and returns a global logger
func NewLogger(isProd bool) MyLogger {
	var output zerolog.ConsoleWriter

	if isProd {
		// JSON output for production
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		// Console output for development
		output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("*** %s ***", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatErrFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s: ", i)
		}
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(output).With().Timestamp().Logger()

	Logger = MyLogger{logger}
	return Logger
}

// CustomLogFunc is the logging function for Echo middleware
func CustomLogFunc(c echo.Context, v middleware.RequestLoggerValues) error {
	Logger.Info().
		Str("path", c.Request().URL.Path).
		Str("method", c.Request().Method).
		Int("status", v.Status).
		Str("remote_ip", v.RemoteIP).
		Str("user_agent", v.UserAgent).
		Dur("latency", v.Latency).
		Str("request_id", c.Response().Header().Get(echo.HeaderXRequestID)).
		Msg("request processed")

	return nil
}

// LogError is a helper function for structured error logging
func LogError(c echo.Context, err error, msg string) {
	Logger.Error().
		Str("path", c.Request().URL.Path).
		Str("method", c.Request().Method).
		Str("request_id", c.Response().Header().Get(echo.HeaderXRequestID)).
		Err(err).
		Msg(msg)
}
