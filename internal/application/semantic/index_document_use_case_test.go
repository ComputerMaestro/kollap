package semantic

import (
	"context"
	"testing"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/google/uuid"
)

type indexTestEmbedder struct{}

func (indexTestEmbedder) Embed(context.Context, string, map[string]any) ([]float32, error) {
	return []float32{1, 2, 3}, nil
}

type indexTestRepository struct{ saved *domain.Embedding }

func (r *indexTestRepository) Save(_ context.Context, embedding *domain.Embedding) error {
	r.saved = embedding
	return nil
}
func (r *indexTestRepository) FindSimilarDocumentIDs(context.Context, uuid.UUID, []float32, int) ([]uuid.UUID, error) {
	return nil, nil
}

func TestIndexDocumentStoresSeparateEmbedding(t *testing.T) {
	docID, workspaceID := uuid.New(), uuid.New()
	repo := &indexTestRepository{}
	uc := NewIndexDocumentUC(indexTestEmbedder{}, repo, "test-model")
	if err := uc.Index(context.Background(), docID.String(), workspaceID.String(), "title", "content"); err != nil {
		t.Fatal(err)
	}
	if repo.saved == nil || repo.saved.DocumentID != docID || repo.saved.WorkspaceID != workspaceID || repo.saved.Model != "test-model" {
		t.Fatalf("unexpected stored embedding: %#v", repo.saved)
	}
	if len(repo.saved.Vector) != 3 {
		t.Fatalf("vector = %v", repo.saved.Vector)
	}
}
