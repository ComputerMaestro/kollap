package models

import (
	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"time"
)

type Embedding struct {
	DocumentID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	WorkspaceID uuid.UUID `gorm:"type:uuid;index"`
	Model       string
	Vector      pgvector.Vector `gorm:"type:vector(1024)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
