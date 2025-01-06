package lib

import (
	"errors"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetUUIDParam extracts and validates a UUID parameter from the Echo context.
func GetUUIDParam(c echo.Context, param string) (uuid.UUID, error) {
	// Get the parameter from the request path
	paramValue := c.Param(param)
	if paramValue == "" {
		return uuid.UUID{}, errors.New("parameter is required")
	}

	// Parse and validate the UUID
	parsedUUID, err := uuid.Parse(paramValue)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid UUID format")
	}

	return parsedUUID, nil
}
