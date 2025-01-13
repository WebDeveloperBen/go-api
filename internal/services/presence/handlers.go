package presence

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webdeveloperben/go-api/internal/lib"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
)

// Interface for the presence handler
type PresenceHandlerInterface interface {
	HandleGetPresences(c echo.Context) error
	HandleGetPresence(c echo.Context) error
	HandleCreatePresence(c echo.Context) error
	HandleUpdatePresence(c echo.Context) error
	HandleDeletePresence(c echo.Context) error
}

var _ PresenceHandlerInterface = (*PresenceHandler)(nil)

// Handler struct
type PresenceHandler struct {
	Service   PresenceServiceInterface
	Validator lib.ValidatorServiceInterface
}

// Create a new presence handler
func NewHandler(s PresenceServiceInterface, v lib.ValidatorServiceInterface) *PresenceHandler {
	return &PresenceHandler{
		Service:   s,
		Validator: v,
	}
}

// Get all presences
func (h *PresenceHandler) HandleGetPresences(c echo.Context) error {
	presence, err := h.Service.GetPresences(c)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusOK, presence)
}

// Get presence by ID
func (h *PresenceHandler) HandleGetPresence(c echo.Context) error {
	parsedUUID, err := lib.GetValidUUIDParam(c, "id")
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	presence, err := h.Service.GetPresenceByID(c, parsedUUID)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusOK, presence)
}

// Create a new presence record

func (h *PresenceHandler) HandleCreatePresence(c echo.Context) error {
	var req CreatePresenceRequest

	err := lib.ValidateInputs(c, h.Validator, &req)
	if err != nil {
		return lib.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	id, err := lib.ParseUUID(req.UserID)
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	presence := repository.InsertPresenceParams{
		LastStatus: req.LastStatus,
		UserID:     id,
		LastLogin:  req.LastLogin,
		LastLogout: req.LastLogout,
	}

	err = h.Service.CreatePresence(c, presence)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}
	return lib.WriteJSON(c, http.StatusCreated, req)
}

// Update presence record
func (h *PresenceHandler) HandleUpdatePresence(c echo.Context) error {
	var req UpdatePresenceRequest

	err := lib.ValidateInputs(c, h.Validator, &req)
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.UpdatePresence(c, req); err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusOK, req)
}

// Delete presence record
func (h *PresenceHandler) HandleDeletePresence(c echo.Context) error {
	parsedUUID, err := lib.GetUUIDParam(c, "id")
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.DeletePresence(c, parsedUUID); err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}
	return lib.WriteJSON(c, http.StatusOK, "presence deleted")
}
