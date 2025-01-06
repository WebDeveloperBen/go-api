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

type MyLogger struct {
	zerolog.Logger
}

// global variable for easy access everywhere
var Logger MyLogger

// create a new global logger
func NewLogger() MyLogger {
	// create output configuration
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Format level: fatal, error, debug, info, warn
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	// format error
	output.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}

	zerolog := zerolog.New(output).With().Caller().Timestamp().Logger()
	Logger = MyLogger{zerolog}
	return Logger
}

// this is the custom logger for the echo logger middleware
func CustomLogFunc(c echo.Context, v middleware.RequestLoggerValues) error {
	Logger.Info().
		Int("status", v.Status).
		Str("method", v.Method).
		Str("uri", v.URI).
		Str("remote_ip", v.RemoteIP).
		Str("user_agent", v.UserAgent).
		Dur("latency", v.Latency).
		Err(v.Error).
		Msg("request completed")

	return nil
}
