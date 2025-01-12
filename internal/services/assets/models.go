package assets

import "github.com/google/uuid"

type GetAssetByFileNameRequest struct {
	FileName string `json:"file_name" validate:"required"`
}

type GetAllAssetsPaginatedRequest struct {
	Limit  int32 `query:"limit" validate:"omitempty,gt=0"`
	Offset int32 `query:"offset" validate:"omitempty,gte=0"`
}

// ApplyDefaults sets default values for the request fields if they are not provided
func (r *GetAllAssetsPaginatedRequest) ApplyDefaults() {
	if r.Limit == 0 {
		r.Limit = 100 // Default limit
	}
	if r.Offset < 0 {
		r.Offset = 0 // Default offset
	}
}

type DeleteAssetRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}

type InsertAssetRequest struct {
	FileName      string `json:"file_name" validate:"required"`
	ContentType   string `json:"content_type" validate:"required"`
	ETag          string `json:"e_tag" validate:"omitempty"`
	ContainerName string `json:"container_name" validate:"required"`
	Uri           string `json:"uri" validate:"required"`
	Size          int32  `json:"size" validate:"required"`
	Metadata      []byte `json:"metadata" validate:"omitempty"`
	IsPublic      bool   `json:"is_public"`
	Published     bool   `json:"published"`
}

type UpdateAssetRequest struct {
	ID            string `param:"id"` // NOTE: must be of type string
	FileName      string `json:"file_name" validate:"omitempty"`
	ContentType   string `json:"content_type" validate:"omitempty"`
	ETag          string `json:"e_tag" validate:"omitempty"`
	ContainerName string `json:"container_name" validate:"omitempty"`
	Uri           string `json:"uri" validate:"omitempty"`
	Size          int32  `json:"size" validate:"omitempty"`
	Metadata      []byte `json:"metadata" validate:"omitempty"`
	IsPublic      bool   `json:"is_public" validate:"omitempty"`
	Published     bool   `json:"published" validate:"omitempty"`
}
