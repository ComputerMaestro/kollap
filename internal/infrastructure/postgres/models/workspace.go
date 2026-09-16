package models

import (
	"database/sql"
	"time"
)

type WorkspaceModel struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string
	OwnerID   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
