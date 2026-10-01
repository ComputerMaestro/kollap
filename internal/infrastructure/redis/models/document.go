package models

import (
	"time"

	"github.com/google/uuid"
)

type CacheDocumentModel struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	WorkspaceID uuid.UUID `json:workspace_id"`
	Content     string    `json:"content"`
	Version     int64     `json:"version"`
	CreatedAt   time.Time `json:created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
