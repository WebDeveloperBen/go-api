package lib

import "github.com/labstack/echo/v4"

func DebugLogger(c echo.Context, msg *string) {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)
	Logger.Error().
		Timestamp().
		Str("path", c.Request().URL.Path).
		Str("method", c.Request().Method).
		Str("request_id", requestID).
		Str("error", *msg).
		Msg("an error occurred")
}
