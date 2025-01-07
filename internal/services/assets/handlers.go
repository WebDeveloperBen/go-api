package assets

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webdeveloperben/go-api/internal/lib"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
)

// Interface for the assets handler
type AssetsHandlerInterface interface {
	HandleGetAllAssets(c echo.Context) error
	HandleGetPublicAssets(c echo.Context) error
	HandleGetAssetByID(c echo.Context) error
	HandleGetAssetByFileName(c echo.Context) error
	HandleCreateAsset(c echo.Context) error
	HandleUpdateAsset(c echo.Context) error
	HandleDeleteAsset(c echo.Context) error
	HandleGetAssetsCount(c echo.Context) error
}

// AssetsHandler struct
type AssetsHandler struct {
	Service   AssetsServiceInterface
	Validator lib.ValidatorServiceInterface
}

// Create a new assets handler
func NewAssetsHandler(service AssetsServiceInterface, validator lib.ValidatorServiceInterface) *AssetsHandler {
	return &AssetsHandler{
		Service:   service,
		Validator: validator,
	}
}

// HandleGetAllAssets retrieves all assets with pagination
func (h *AssetsHandler) HandleGetAllAssets(c echo.Context) error {
	limit, offset, err := lib.ParsePaginationParams(c)
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	assets, err := h.Service.GetAllAssets(c, limit, offset)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusOK, assets)
}

// HandleGetPublicAssets retrieves public assets with pagination
func (h *AssetsHandler) HandleGetPublicAssets(c echo.Context) error {
	limit, offset, err := lib.ParsePaginationParams(c)
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	assets, err := h.Service.GetPublicAssets(c, limit, offset)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusOK, assets)
}

// HandleGetAssetByID retrieves an asset by ID
func (h *AssetsHandler) HandleGetAssetByID(c echo.Context) error {
	id, err := lib.GetUUIDParam(c, "id")
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	asset, err := h.Service.GetAssetByID(c, id)
	if err != nil {
		return lib.WriteError(c, http.StatusNotFound, err)
	}

	return lib.WriteJSON(c, http.StatusOK, asset)
}

// HandleGetAssetByFileName retrieves an asset by file name
func (h *AssetsHandler) HandleGetAssetByFileName(c echo.Context) error {
	fileName, err := lib.GetValidParam(c, "fileName")
	if err != nil {
		return lib.WriteError(c, http.StatusUnprocessableEntity, err)
	}

	asset, err := h.Service.GetAssetByFileName(c, fileName)
	if err != nil {
		return lib.WriteError(c, http.StatusNotFound, err)
	}

	return lib.WriteJSON(c, http.StatusOK, asset)
}

// HandleCreateAsset creates a new asset
func (h *AssetsHandler) HandleCreateAsset(c echo.Context) error {
	var req repository.InsertAssetParams
	if err := c.Bind(&req); err != nil {
		return lib.InvalidJSON(c)
	}

	if err := lib.ValidateRequest(h.Validator, req); err != nil {
		return lib.InvalidRequest(c, err)
	}

	asset, err := h.Service.CreateAsset(c, req)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusCreated, asset)
}

// HandleUpdateAsset updates an existing asset
func (h *AssetsHandler) HandleUpdateAsset(c echo.Context) error {
	id, err := lib.GetUUIDParam(c, "id")
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	var req repository.UpdateAssetParams
	if err := c.Bind(&req); err != nil {
		return lib.InvalidJSON(c)
	}

	if err := lib.ValidateRequest(h.Validator, req); err != nil {
		return lib.InvalidRequest(c, err)
	}

	asset, err := h.Service.UpdateAsset(c, id, req)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusOK, asset)
}

// HandleDeleteAsset deletes an asset by ID
func (h *AssetsHandler) HandleDeleteAsset(c echo.Context) error {
	id, err := lib.GetUUIDParam(c, "id")
	if err != nil {
		return lib.WriteError(c, http.StatusBadRequest, err)
	}

	if err := h.Service.DeleteAsset(c, id); err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusOK, "asset deleted")
}

// HandleGetAssetsCount retrieves the total count of assets
func (h *AssetsHandler) HandleGetAssetsCount(c echo.Context) error {
	count, err := h.Service.GetAssetsCount(c)
	if err != nil {
		return lib.WriteError(c, http.StatusInternalServerError, err)
	}

	return lib.WriteJSON(c, http.StatusOK, map[string]int64{"count": count})
}
