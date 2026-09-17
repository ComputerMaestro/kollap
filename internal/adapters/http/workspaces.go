package http

import (
	"context"
	"time"

	workspaces "github.com/ComputerMaestro/kollap/gen/workspaces"
	"github.com/ComputerMaestro/kollap/internal/application/workspace"
	"goa.design/clue/log"
)

// workspaces service example implementation.
// The example methods log the requests and return zero values.
type workspacessrvc struct {
	createWorkspaceUC *workspace.CreateWorkspaceUC
}

// NewWorkspaces returns the workspaces service implementation.
func NewWorkspaces(createWorkspaceUC *workspace.CreateWorkspaceUC) workspaces.Service {
	return &workspacessrvc{
		createWorkspaceUC: createWorkspaceUC,
	}
}

// GetWorkspace implements getWorkspace.
func (s *workspacessrvc) CreateWorkspace(ctx context.Context, p *workspaces.CreateWorkspacePayload) (res *workspaces.Workspace, err error) {
	w, err := s.createWorkspaceUC.Execute(ctx, p.Name, p.OwnerID)
	if err != nil {
		log.Errorf(ctx, err, "failed to execute create workspace use case")
		return nil, err
	}
	res = &workspaces.Workspace{
		ID:        w.ID.String(),
		Name:      w.Name,
		OwnerID:   w.OwnerID.String(),
		CreatedAt: w.CreatedAt.Format(time.RFC3339),
	}
	log.Printf(ctx, "workspaces.getWorkspace")
	return
}
