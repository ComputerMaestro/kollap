package repository

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/google/uuid"
)

var (
	WorkspaceDao WorkspaceRepository
)

type WorkspaceRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Workspace, error)
	Create(ctx context.Context, w *domain.Workspace) (*domain.Workspace, error)
}
