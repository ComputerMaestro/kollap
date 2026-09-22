package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type GetDocumentUC struct {
	repo repository.DocumentRepository
}

func NewGetDocumentUC(repo repository.DocumentRepository) *GetDocumentUC {
	return &GetDocumentUC{
		repo: repo,
	}
}

func (uc *GetDocumentUC) Execute(ctx context.Context, id string) (*domain.Document, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		log.Errorf(ctx, err, "failed to parse document uuid")
		return nil, err
	}
	return uc.repo.FindByID(ctx, parsedUUID)
}
