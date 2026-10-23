package document

import (
	"context"
	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
)

const similarDocumentsLimit = 5

type GetSimilarDocumentsUC struct {
	embedder   repository.Embedder
	embeddings repository.EmbeddingRepository
	documents  repository.DocumentRepository
}

func NewGetSimilarDocumentsUC(embedder repository.Embedder, embeddings repository.EmbeddingRepository, documents repository.DocumentRepository) *GetSimilarDocumentsUC {
	return &GetSimilarDocumentsUC{embedder: embedder, embeddings: embeddings, documents: documents}
}

func (uc *GetSimilarDocumentsUC) Execute(ctx context.Context, workspaceID, query string) ([]*domain.Document, error) {
	if uc.embedder == nil || uc.embeddings == nil || uc.documents == nil {
		return nil, domain.ErrSemanticSearchUnavailable
	}
	wsID, err := uuid.Parse(workspaceID)
	if err != nil {
		return nil, err
	}
	vector, err := uc.embedder.Embed(ctx, query, map[string]any{})
	if err != nil {
		return nil, err
	}
	ids, err := uc.embeddings.FindSimilarDocumentIDs(ctx, wsID, vector, similarDocumentsLimit)
	if err != nil {
		return nil, err
	}
	docs := make([]*domain.Document, 0, len(ids))
	for _, id := range ids {
		doc, err := uc.documents.FindByID(ctx, id)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}
