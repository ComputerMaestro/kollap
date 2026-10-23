package domain

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID          uuid.UUID
	Title       string
	WorkspaceID uuid.UUID
	Content     string
	Version     int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
