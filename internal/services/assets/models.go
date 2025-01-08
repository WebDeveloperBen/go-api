package assets

import "github.com/google/uuid"

type GetAssetByFileNameRequest struct {
	FileName string `json:"file_name" validate:"required"`
}

type DeleteAssetRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
