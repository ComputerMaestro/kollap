package workspace

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type GetWorkspaceUC struct {
	repo repository.WorkspaceRepository
}

func NewGetWorkspaceUC(repo repository.WorkspaceRepository) *GetWorkspaceUC {
	return &GetWorkspaceUC{
		repo: repo,
	}
}

func (uc *GetWorkspaceUC) Execute(ctx context.Context, id string) (*domain.Workspace, error) {
	parseUUID, err := uuid.Parse(id)
	if err != nil {
		log.Errorf(ctx, err, "failed to parse uuid")
		return nil, err
	}
	return uc.repo.FindByID(ctx, parseUUID)
}
