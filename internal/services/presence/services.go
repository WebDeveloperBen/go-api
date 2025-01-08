package presence

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	repository "github.com/webdeveloperben/go-api/internal/repository/generated"
)

// Interface for the presence service
type PresenceServiceInterface interface {
	GetPresences(ctx echo.Context) ([]repository.GetAllPresenceRow, error)
	GetPresence(ctx echo.Context, id uuid.UUID) (*repository.GetPresenceByIDRow, error)
	CreatePresence(ctx echo.Context, presence repository.InsertPresenceParams) error
	UpdatePresence(ctx echo.Context, presence repository.UpdatePresenceParams) error
	DeletePresence(ctx echo.Context, id uuid.UUID) error
}

type PresenceService struct {
	Storage PresenceStorageInterface
}

// NewPresenceService creates a new presence service
func NewService(storage PresenceStorageInterface) *PresenceService {
	return &PresenceService{
		Storage: storage,
	}
}

// GetPresences retrieves all presences
func (s *PresenceService) GetPresences(ctx echo.Context) ([]repository.GetAllPresenceRow, error) {
	return s.Storage.GetAllPresence(ctx.Request().Context())
}

// GetPresence retrieves a presence record by ID
func (s *PresenceService) GetPresence(ctx echo.Context, id uuid.UUID) (*repository.GetPresenceByIDRow, error) {
	return s.Storage.GetPresenceByID(ctx.Request().Context(), id)
}

// CreatePresence creates a new presence record
func (s *PresenceService) CreatePresence(ctx echo.Context, presence repository.InsertPresenceParams) error {
	return s.Storage.InsertPresence(ctx.Request().Context(), presence)
}

// UpdatePresence updates an existing presence record
func (s *PresenceService) UpdatePresence(ctx echo.Context, presence repository.UpdatePresenceParams) error {
	return s.Storage.UpdatePresence(ctx.Request().Context(), presence)
}

// DeletePresence deletes a presence record by ID
func (s *PresenceService) DeletePresence(ctx echo.Context, id uuid.UUID) error {
	return s.Storage.DeletePresence(ctx.Request().Context(), id)
}
