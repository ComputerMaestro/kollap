package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
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
	return uc.repo.FindByID(ctx, id)
}
