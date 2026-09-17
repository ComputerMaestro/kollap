package workspace

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type CreateWorkspaceUC struct {
	repo repository.WorkspaceRepository
}

func NewCreateWorkspaceUC(repo repository.WorkspaceRepository) *CreateWorkspaceUC {
	return &CreateWorkspaceUC{
		repo: repo,
	}
}

func (uc *CreateWorkspaceUC) Execute(ctx context.Context, name string, ownerId string) (*domain.Workspace, error) {
	parsedUUID, err := uuid.Parse(ownerId)
	if err != nil {
		log.Errorf(ctx, err, "failed to parse UUID")
		return nil, err
	}
	workspace := &domain.Workspace{
		ID:      uuid.New(),
		Name:    name,
		OwnerID: parsedUUID,
	}

	return uc.repo.Create(ctx, workspace)
}
