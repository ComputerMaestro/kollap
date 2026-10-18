package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
)

type GetSimilarDocumentsUC struct {
	searcher repository.SemanticSearcher
}

func NewGetSimilarDocumentsUC(searcher repository.SemanticSearcher) *GetSimilarDocumentsUC {
	return &GetSimilarDocumentsUC{
		searcher: searcher,
	}
}

func (uc *GetSimilarDocumentsUC) Execute(ctx context.Context, query string) ([]*domain.Document, error) {
	return uc.searcher.Search(ctx, query)
}
