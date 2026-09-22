package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type CreateDocumentUC struct {
	repo repository.DocumentRepository
}

func NewCreateDocumentUC(repo repository.DocumentRepository) *CreateDocumentUC {
	return &CreateDocumentUC{
		repo: repo,
	}
}

func (uc *CreateDocumentUC) Execute(ctx context.Context, name string, workspaceID string, content string) (*domain.Document, error) {
	workspaceUUID, err := uuid.Parse(workspaceID)
	if err != nil {
		log.Errorf(ctx, err, "failed to parse UUID")
		return nil, err
	}
	Document := &domain.Document{
		ID:          uuid.New(),
		Title:       name,
		Version:     1,
		Content:     content,
		WorkspaceID: workspaceUUID,
	}

	return uc.repo.Create(ctx, Document)
}
