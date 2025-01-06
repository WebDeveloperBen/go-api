package presence

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webdeveloperben/ise-go-api/internal/lib"
	repository "github.com/webdeveloperben/ise-go-api/internal/repository/postgres"
)

// Interface for the presence handler
type PresenceHandlerInterface interface {
	HandleGetPresence(c echo.Context) error
	HandleCreatePresence(c echo.Context) error
	HandleUpdatePresence(c echo.Context) error
	HandleDeletePresence(c echo.Context) error
}

// Handler struct
type PresenceHandler struct {
	Service   PresenceServiceInterface
	Validator lib.ValidatorServiceInterface
}

// Create a new presence handler
func NewPresenceHandler(s PresenceServiceInterface, v lib.ValidatorServiceInterface) *PresenceHandler {
	return &PresenceHandler{
		Service:   s,
		Validator: v,
	}
}

// Get presence by ID
func (h *PresenceHandler) HandleGetPresence(c echo.Context) error {
	parsedUUID, err := lib.GetUUIDParam(c, "id")
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	req := HandleGetPresenceRequest{ID: parsedUUID}
	if err := lib.ValidateRequest(h.Validator, req); err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	presence, err := h.Service.GetPresence(c.Request().Context(), parsedUUID)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, presence)
}

// Create a new presence record
func (h *PresenceHandler) HandleCreatePresence(c echo.Context) error {
	var presence repository.InsertPresenceParams
	if err := c.Bind(&presence); err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.CreatePresence(c.Request().Context(), presence); err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, presence)
}

// Update presence record
func (h *PresenceHandler) HandleUpdatePresence(c echo.Context) error {
	var presence repository.InsertPresenceParams
	if err := c.Bind(&presence); err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.UpdatePresence(c.Request().Context(), presence); err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, presence)
}

// Delete presence record
func (h *PresenceHandler) HandleDeletePresence(c echo.Context) error {
	parsedUUID, err := lib.GetUUIDParam(c, "id")
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}
	if err := h.Service.DeletePresence(c.Request().Context(), parsedUUID); err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}
	return lib.WriteJSON(c, http.StatusOK, "presence deleted")
}
