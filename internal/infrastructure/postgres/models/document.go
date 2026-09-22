package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string
	WorkspaceID uuid.UUID `gorm:"type:uuid"`
	Content     string
	Version     int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
