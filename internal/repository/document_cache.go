package repository

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/google/uuid"
)

type DocumentCacheRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Document, error)
	Set(ctx context.Context, doc *domain.Document) error
	Unlink(ctx context.Context, id uuid.UUID) error
}
