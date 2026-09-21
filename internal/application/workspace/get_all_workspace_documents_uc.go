package workspace

import (
	"context"
	"log"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
)

type GetAllWorkspaceDocumentsUC struct {
	docRepo repository.DocumentRepository
}

func NewGetAllWorkspaceDocumentsUC(docRepo repository.DocumentRepository) *GetAllWorkspaceDocumentsUC {
	return &GetAllWorkspaceDocumentsUC{
		docRepo: docRepo,
	}
}

func (uc *GetAllWorkspaceDocumentsUC) Execute(ctx context.Context, id string) ([]*domain.Document, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		log.Fatalf("failed to parse UUID string: %v", err)
	}
	filter := repository.DocumentFilter{
		WorkspaceID: &parsedUUID,
	}
	return uc.docRepo.Find(ctx, filter)
}
