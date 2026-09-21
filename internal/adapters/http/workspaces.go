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
	createWorkspaceUC          *workspace.CreateWorkspaceUC
	getWorkspaceUC             *workspace.GetWorkspaceUC
	getAllWorkspaceDocumentsUC *workspace.GetAllWorkspaceDocumentsUC
}

// NewWorkspaces returns the workspaces service implementation.
func NewWorkspaces(
	createWorkspaceUC *workspace.CreateWorkspaceUC,
	getWorkspaceUC *workspace.GetWorkspaceUC,
	getAllWorkspaceDocumentsUC *workspace.GetAllWorkspaceDocumentsUC,
) workspaces.Service {
	return &workspacessrvc{
		createWorkspaceUC:          createWorkspaceUC,
		getWorkspaceUC:             getWorkspaceUC,
		getAllWorkspaceDocumentsUC: getAllWorkspaceDocumentsUC,
	}
}

// GetWorkspace implements getWorkspace.
func (s *workspacessrvc) GetWorkspace(ctx context.Context, p *workspaces.GetWorkspacePayload) (res *workspaces.Workspace, err error) {
	w, err := s.getWorkspaceUC.Execute(ctx, p.ID)
	if err != nil {
		log.Errorf(ctx, err, "failed to execute get workspace use case")
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
	log.Printf(ctx, "workspaces.createWorkspace")
	return
}

// GetAllWorkspaceDocuments implements getAllWorkspaceDocuments.
func (s *workspacessrvc) GetAllWorkspaceDocuments(ctx context.Context, p *workspaces.GetAllWorkspaceDocumentsPayload) (res *workspaces.GetAllWorkspaceDocumentsResult, err error) {
	w, err := s.getAllWorkspaceDocumentsUC.Execute(ctx, p.ID)
	if err != nil {
		log.Errorf(ctx, err, "failed to execute get all workspace documents use case")
		return nil, err
	}
	res = &workspaces.GetAllWorkspaceDocumentsResult{}
	for _, d := range w {
		res.Documents = append(res.Documents, &workspaces.Document{
			ID:          d.ID.String(),
			Title:       d.Title,
			Version:     d.Version,
			Content:     d.Content,
			WorkspaceID: d.WorkspaceID.String(),
			CreatedAt:   d.CreatedAt.Format(time.RFC3339),
		})
	}
	log.Printf(ctx, "workspaces.getAllWorkspaceDocuments")
	return
}
