package assets

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
)

// Interface for the assets service
type AssetsServiceInterface interface {
	GetAllAssets(ctx echo.Context, limit, offset int) ([]repository.Asset, error)
	GetPublicAssets(ctx echo.Context, limit, offset int) ([]repository.Asset, error)
	GetAssetByID(ctx echo.Context, id uuid.UUID) (*repository.Asset, error)
	GetAssetByFileName(ctx echo.Context, fileName string) (*repository.Asset, error)
	CreateAsset(ctx echo.Context, asset InsertAssetRequest) (repository.Asset, error)
	UpdateAsset(ctx echo.Context, asset UpdateAssetRequest) (repository.Asset, error)
	DeleteAsset(ctx echo.Context, id uuid.UUID) error
	GetAssetsCount(ctx echo.Context) (int64, error)
}

var _ AssetsServiceInterface = (*AssetsService)(nil)

type AssetsService struct {
	Storage AssetsStorageInterface
}

// NewAssetsService creates a new assets service
func NewService(storage AssetsStorageInterface) *AssetsService {
	return &AssetsService{
		Storage: storage,
	}
}

// GetAllAssets retrieves all assets with pagination
func (s *AssetsService) GetAllAssets(ctx echo.Context, limit, offset int) ([]repository.Asset, error) {
	return s.Storage.GetAllAssetsPaginated(ctx.Request().Context(), limit, offset)
}

// GetPublicAssets retrieves all public assets with pagination
func (s *AssetsService) GetPublicAssets(ctx echo.Context, limit, offset int) ([]repository.Asset, error) {
	return s.Storage.GetPublicAssetsPaginated(ctx.Request().Context(), limit, offset)
}

// GetAssetByID retrieves an asset by ID
func (s *AssetsService) GetAssetByID(ctx echo.Context, id uuid.UUID) (*repository.Asset, error) {
	return s.Storage.GetAssetByID(ctx.Request().Context(), id)
}

// GetAssetByFileName retrieves an asset by file name
func (s *AssetsService) GetAssetByFileName(ctx echo.Context, fileName string) (*repository.Asset, error) {
	return s.Storage.GetAssetByFileName(ctx.Request().Context(), fileName)
}

// CreateAsset creates a new asset and returns the created record
func (s *AssetsService) CreateAsset(ctx echo.Context, req InsertAssetRequest) (repository.Asset, error) {
	// Map InsertAssetRequest to InsertAssetParams
	params := repository.InsertAssetParams{
		FileName:      req.FileName,
		ContentType:   req.ContentType,
		ETag:          req.ETag,
		ContainerName: req.ContainerName,
		Uri:           req.Uri,
		Size:          req.Size,
		Metadata:      req.Metadata,
		IsPublic:      req.IsPublic,
		Published:     req.Published,
	}

	// Call the database layer
	return s.Storage.InsertAsset(ctx.Request().Context(), params)
}

// UpdateAsset updates an existing asset and returns the updated record
func (s *AssetsService) UpdateAsset(ctx echo.Context, req UpdateAssetRequest) (repository.Asset, error) {
	parsedUUID, err := uuid.Parse(req.ID)
	if err != nil {
		return repository.Asset{}, fmt.Errorf("invalid uuid: %v", err)
	}
	// Map InsertAssetRequest to InsertAssetParams
	params := repository.UpdateAssetParams{
		ID:            parsedUUID,
		FileName:      req.FileName,
		ContentType:   req.ContentType,
		ETag:          req.ETag,
		ContainerName: req.ContainerName,
		Uri:           req.Uri,
		Size:          req.Size,
		Metadata:      req.Metadata,
		IsPublic:      req.IsPublic,
		Published:     req.Published,
	}
	return s.Storage.UpdateAsset(ctx.Request().Context(), params.ID, params)
}

// DeleteAsset deletes an asset by ID
func (s *AssetsService) DeleteAsset(ctx echo.Context, id uuid.UUID) error {
	return s.Storage.DeleteAsset(ctx.Request().Context(), id)
}

// GetAssetsCount retrieves the total count of assets
func (s *AssetsService) GetAssetsCount(ctx echo.Context) (int64, error) {
	return s.Storage.GetAssetsCount(ctx.Request().Context())
}
