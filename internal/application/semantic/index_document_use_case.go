package semantic

import (
	"context"
	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
)

type IndexDocumentUC struct {
	embedder   repository.Embedder
	embeddings repository.EmbeddingRepository
	model      string
}

func NewIndexDocumentUC(embedder repository.Embedder, embeddings repository.EmbeddingRepository, model string) *IndexDocumentUC {
	return &IndexDocumentUC{embedder: embedder, embeddings: embeddings, model: model}
}

func (uc *IndexDocumentUC) Index(ctx context.Context, documentID, workspaceID, title, content string) error {
	vector, err := uc.embedder.Embed(ctx, content, map[string]any{"title": title})
	if err != nil {
		return err
	}
	docID, err := uuid.Parse(documentID)
	if err != nil {
		return err
	}
	wsID, err := uuid.Parse(workspaceID)
	if err != nil {
		return err
	}
	return uc.embeddings.Save(ctx, &domain.Embedding{DocumentID: docID, WorkspaceID: wsID, Model: uc.model, Vector: vector})
}
