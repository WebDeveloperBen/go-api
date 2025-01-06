package presence

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/webdeveloperben/go-api/internal/lib"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
)

type PresenceStorageInterface interface {
	GetPresenceByID(ctx echo.Context, id uuid.UUID) (*repository.GetPresenceByIDRow, error)
	InsertPresence(ctx echo.Context, presence repository.InsertPresenceParams) error
	UpdatePresence(ctx echo.Context, presence repository.UpdatePresenceParams) error
	DeletePresence(ctx echo.Context, id uuid.UUID) error
	GetAllPresence(ctx echo.Context) ([]repository.GetAllPresenceRow, error)
	UpdateLogoutTime(ctx echo.Context, params repository.UpdateLogoutTimeParams) error
}

type PresenceStorage struct {
	Queries *repository.Queries
}

// NewPresenceStorage creates a new storage layer for presence
func NewPresenceStorage(queries *repository.Queries) *PresenceStorage {
	return &PresenceStorage{
		Queries: queries,
	}
}

// GetPresenceByID retrieves a presence record by ID
func (s *PresenceStorage) GetPresenceByID(ctx echo.Context, id uuid.UUID) (*repository.GetPresenceByIDRow, error) {
	row, err := s.Queries.GetPresenceByID(ctx.Request().Context(), id)
	if err != nil {
		return nil, lib.WriteError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to fetch presence: %w", err))
	}
	return &row, nil
}

// InsertPresence inserts a new presence record
func (s *PresenceStorage) InsertPresence(ctx echo.Context, presence repository.InsertPresenceParams) error {
	if err := s.Queries.InsertPresence(ctx.Request().Context(), presence); err != nil {
		return lib.WriteError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to insert presence: %w", err))
	}
	return nil
}

// UpdatePresence updates an existing presence record
func (s *PresenceStorage) UpdatePresence(ctx echo.Context, presence repository.UpdatePresenceParams) error {
	if err := s.Queries.UpdatePresence(ctx.Request().Context(), presence); err != nil {
		return lib.WriteError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to update presence: %w", err))
	}
	return nil
}

// DeletePresence deletes a presence record by ID
func (s *PresenceStorage) DeletePresence(ctx echo.Context, id uuid.UUID) error {
	if err := s.Queries.DeletePresence(ctx.Request().Context(), id); err != nil {
		return lib.WriteError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to delete presence: %w", err))
	}
	return nil
}

// GetAllPresence retrieves all presence records
func (s *PresenceStorage) GetAllPresence(ctx echo.Context) ([]repository.GetAllPresenceRow, error) {
	rows, err := s.Queries.GetAllPresence(ctx.Request().Context())
	if err != nil {
		// Check if the error is due to no rows being found
		if err == sql.ErrNoRows {
			return []repository.GetAllPresenceRow{}, nil
		}
		return nil, fmt.Errorf("failed to fetch presences: %w", err)
	}
	return rows, nil
}

// UpdateLogoutTime updates the logout time for a presence record
func (s *PresenceStorage) UpdateLogoutTime(ctx echo.Context, params repository.UpdateLogoutTimeParams) error {
	if err := s.Queries.UpdateLogoutTime(ctx.Request().Context(), params); err != nil {
		return lib.WriteError(ctx, http.StatusInternalServerError, fmt.Errorf("failed to update logout time: %w", err))
	}
	return nil
}
