package repository

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/google/uuid"
)

type DocumentCacheRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Document, error)
}
