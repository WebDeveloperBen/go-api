package assets

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/webdeveloperben/go-api/internal/lib"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
)

// AssetsStorageInterface defines methods that the assets service can call
type AssetsStorageInterface interface {
	GetAllAssetsPaginated(ctx echo.Context, limit, offset int) ([]repository.Asset, error)
	GetPublicAssetsPaginated(ctx echo.Context, limit, offset int) ([]repository.Asset, error)
	GetAssetByID(ctx echo.Context, id uuid.UUID) (*repository.Asset, error)
	GetAssetByFileName(ctx echo.Context, fileName string) (*repository.Asset, error)
	InsertAsset(ctx echo.Context, asset repository.InsertAssetParams) (*repository.Asset, error)
	UpdateAsset(ctx echo.Context, id uuid.UUID, asset repository.UpdateAssetParams) (*repository.Asset, error)
	DeleteAsset(ctx echo.Context, id uuid.UUID) error
	GetAssetsCount(ctx echo.Context) (int64, error)
}

type AssetsStorage struct {
	queries *repository.Queries
}

// NewAssetsStorage creates a new instance of AssetsStorage
func NewAssetsStorage(queries *repository.Queries) *AssetsStorage {
	return &AssetsStorage{queries: queries}
}

func (s *AssetsStorage) GetAllAssetsPaginated(ctx context.Context, limit, offset int) ([]repository.Asset, error) {
	params := repository.GetAllAssetsPaginatedParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	assets, err := s.queries.GetAllAssetsPaginated(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, lib.NewPublicError("assets not found", "no asset records found")
		}
		return nil, lib.NewPublicError("failed to fetch assets", fmt.Sprintf("database error: %v", err))
	}
	return assets, nil
}

func (s *AssetsStorage) GetPublicAssetsPaginated(ctx context.Context, limit, offset int) ([]repository.Asset, error) {
	params := repository.GetPublicAssetsPaginatedParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	asset, err := s.queries.GetPublicAssetsPaginated(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, lib.NewPublicError("assets not found", "no asset records found")
		}
		return nil, lib.NewPublicError("failed to fetch assets", fmt.Sprintf("database error: %v", err))
	}
	return asset, nil
}

func (s *AssetsStorage) GetAssetByID(ctx context.Context, id uuid.UUID) (*repository.Asset, error) {
	asset, err := s.queries.GetAssetByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, lib.NewPublicError("asset not found", fmt.Sprintf("no asset record found for id: %s", id))
		}
		return nil, lib.NewPublicError("failed to fetch asset", fmt.Sprintf("database error: %v", err))
	}
	return &asset, nil
}

func (s *AssetsStorage) GetAssetByFileName(ctx context.Context, fileName string) (*repository.Asset, error) {
	asset, err := s.queries.GetAssetByFileName(ctx, fileName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, lib.NewPublicError("asset not found", fmt.Sprintf("no asset record found for filename: %s", fileName))
		}
		return nil, lib.NewPublicError("failed to fetch asset", fmt.Sprintf("database error: %v", err))
	}
	return &asset, nil
}

func (s *AssetsStorage) InsertAsset(ctx context.Context, asset repository.InsertAssetParams) (repository.Asset, error) {
	return s.queries.InsertAsset(ctx, asset)
}

func (s *AssetsStorage) UpdateAsset(ctx context.Context, id uuid.UUID, asset repository.UpdateAssetParams) (repository.Asset, error) {
	return s.queries.UpdateAsset(ctx, asset)
}

func (s *AssetsStorage) DeleteAsset(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteAsset(ctx, id)
}

func (s *AssetsStorage) GetAssetsCount(ctx context.Context) (int64, error) {
	return s.queries.GetAssetsCount(ctx)
}
