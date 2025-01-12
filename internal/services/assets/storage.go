package assets

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/webdeveloperben/go-api/internal/lib"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
)

// AssetsStorageInterface defines methods that the assets service can call
type AssetsStorageInterface interface {
	GetAllAssetsPaginated(ctx context.Context, limit, offset int) ([]repository.Asset, error)
	GetPublicAssetsPaginated(ctx context.Context, limit, offset int) ([]repository.Asset, error)
	GetAssetByID(ctx context.Context, id uuid.UUID) (*repository.Asset, error)
	GetAssetByFileName(ctx context.Context, fileName string) (*repository.Asset, error)
	InsertAsset(ctx context.Context, asset repository.InsertAssetParams) (repository.Asset, error)
	UpdateAsset(ctx context.Context, id uuid.UUID, asset repository.UpdateAssetParams) (repository.Asset, error)
	DeleteAsset(ctx context.Context, id uuid.UUID) error
	GetAssetsCount(ctx context.Context) (int64, error)
}

var _ AssetsStorageInterface = (*AssetsStorage)(nil)

type AssetsStorage struct {
	queries *repository.Queries
}

// NewAssetsStorage creates a new instance of AssetsStorage
func NewStorage(queries *repository.Queries) *AssetsStorage {
	return &AssetsStorage{queries: queries}
}

func (s *AssetsStorage) GetAllAssetsPaginated(ctx context.Context, limit, offset int) ([]repository.Asset, error) {
	params := repository.GetAllAssetsPaginatedParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	assets, err := s.queries.GetAllAssetsPaginated(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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

	lib.Logger.Info().Msg("HERE")
	asset, err := s.queries.GetPublicAssetsPaginated(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, lib.NewPublicError("assets not found", "no asset records found")
		}
		return nil, lib.NewPublicError("failed to fetch assets", fmt.Sprintf("database error: %v", err))
	}

	lib.Logger.Info().Msg("asset")
	return asset, nil
}

func (s *AssetsStorage) GetAssetByID(ctx context.Context, id uuid.UUID) (*repository.Asset, error) {
	asset, err := s.queries.GetAssetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Check for wrapped errors
			return nil, lib.NewPublicError("asset not found", fmt.Sprintf("no asset record found for id: %s", id))
		}
		return nil, lib.NewPublicError("asset not found", fmt.Sprintf("database error: %v", err))
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
	updatedAsset, err := s.queries.UpdateAsset(ctx, asset)
	if err != nil {
		fmt.Printf("Error type: %T, Error value: %v\n", err, err)

		// Handle `sql.ErrNoRows` by string matching as a fallback
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "no rows in result set" {
			return repository.Asset{}, lib.NewPublicError("asset not found", fmt.Sprintf("no asset record found for ID: %s", id.String()))
		}

		// Return other errors as internal server errors
		return repository.Asset{}, lib.NewPublicError("failed to update asset", fmt.Sprintf("database error: %v", err))
	}

	return updatedAsset, nil
}

func (s *AssetsStorage) DeleteAsset(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteAsset(ctx, id)
}

func (s *AssetsStorage) GetAssetsCount(ctx context.Context) (int64, error) {
	return s.queries.GetAssetsCount(ctx)
}
