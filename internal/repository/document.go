package repository

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/google/uuid"
)

type DocumentRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Document, error)
	Find(ctx context.Context, filter DocumentFilter) ([]*domain.Document, error)
	Create(ctx context.Context, document *domain.Document) (*domain.Document, error)
}

type DocumentFilter struct {
	WorkspaceID *uuid.UUID
	OwnerID     *uuid.UUID
	Title       *string
}
