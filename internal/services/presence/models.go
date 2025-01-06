package presence

import (
	"time"

	"github.com/google/uuid"
)

type HandleGetPresenceRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
type CreatePresenceRequest struct {
	UserID     string     `json:"user_id" validate:"required,uuid"`
	LastStatus string     `json:"last_status" validate:"required"`
	LastLogin  *time.Time `json:"last_login"`
	LastLogout *time.Time `json:"last_logout"`
}
