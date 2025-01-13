package presence

import (
	"time"

	"github.com/google/uuid"
)

type GetPresenceRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}

type CreatePresenceRequest struct {
	UserID     string     `json:"user_id" validate:"required,uuid"`
	LastStatus string     `json:"last_status" validate:"required"`
	LastLogin  *time.Time `json:"last_login"`
	LastLogout *time.Time `json:"last_logout"`
}

type UpdatePresenceRequest struct {
	ID         string     `param:"id"`
	UserID     uuid.UUID  `json:"user_id"`
	LastStatus string     `json:"last_status"`
	LastLogin  *time.Time `json:"last_login"`
	LastLogout *time.Time `json:"last_logout"`
}
