package workspace

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
)

type CreateWorkspaceUC struct {
	repo repository.WorkspaceRepository
}

func NewCreateWorkspaceUC(repo repository.WorkspaceRepository) *CreateWorkspaceUC {
	return &CreateWorkspaceUC{
		repo: repo,
	}
}

func (uc *CreateWorkspaceUC) Execute(ctx context.Context, name string) (*domain.Workspace, error) {

	workspace := &domain.Workspace{
		Name: name,
	}

	return uc.repo.Create(ctx, workspace)
}
