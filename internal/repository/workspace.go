package repository

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
)

var (
	WorkspaceDao WorkspaceRepository
)

type WorkspaceRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Workspace, error)
	Create(ctx context.Context, w *domain.Workspace) (*domain.Workspace, error)
}
