package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type WorkspaceModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string
	OwnerID   uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
