package document

import (
	"context"
	"errors"
	"testing"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
)

type testEmbedder struct {
	vector []float32
	err    error
}

func (e testEmbedder) Embed(context.Context, string, map[string]any) ([]float32, error) {
	return e.vector, e.err
}

type testEmbeddingRepository struct {
	workspaceID uuid.UUID
	vector      []float32
	ids         []uuid.UUID
}

func (r *testEmbeddingRepository) Save(context.Context, *domain.Embedding) error { return nil }
func (r *testEmbeddingRepository) FindSimilarDocumentIDs(_ context.Context, workspaceID uuid.UUID, vector []float32, _ int) ([]uuid.UUID, error) {
	r.workspaceID, r.vector = workspaceID, vector
	return r.ids, nil
}

type testDocumentRepository struct {
	docs map[uuid.UUID]*domain.Document
}

func (r *testDocumentRepository) FindByID(_ context.Context, id uuid.UUID) (*domain.Document, error) {
	return r.docs[id], nil
}
func (r *testDocumentRepository) Find(context.Context, repository.DocumentFilter) ([]*domain.Document, error) {
	return nil, nil
}
func (r *testDocumentRepository) Create(context.Context, *domain.Document) (*domain.Document, error) {
	return nil, nil
}
func (r *testDocumentRepository) Update(context.Context, *domain.Document) (*domain.Document, error) {
	return nil, nil
}

func TestGetSimilarDocumentsScopesToWorkspaceAndReturnsDocuments(t *testing.T) {
	workspaceID, docID := uuid.New(), uuid.New()
	embeddings := &testEmbeddingRepository{ids: []uuid.UUID{docID}}
	documents := &testDocumentRepository{docs: map[uuid.UUID]*domain.Document{docID: {ID: docID, WorkspaceID: workspaceID}}}
	uc := NewGetSimilarDocumentsUC(testEmbedder{vector: []float32{0.1, 0.2}}, embeddings, documents)

	got, err := uc.Execute(context.Background(), workspaceID.String(), "query")
	if err != nil {
		t.Fatal(err)
	}
	if embeddings.workspaceID != workspaceID {
		t.Fatalf("workspace = %v, want %v", embeddings.workspaceID, workspaceID)
	}
	if len(embeddings.vector) != 2 || embeddings.vector[0] != 0.1 {
		t.Fatalf("unexpected query vector: %v", embeddings.vector)
	}
	if len(got) != 1 || got[0].ID != docID {
		t.Fatalf("unexpected search results: %#v", got)
	}
}

func TestGetSimilarDocumentsDisabled(t *testing.T) {
	uc := NewGetSimilarDocumentsUC(nil, nil, nil)
	_, err := uc.Execute(context.Background(), uuid.NewString(), "query")
	if !errors.Is(err, domain.ErrSemanticSearchUnavailable) {
		t.Fatalf("error = %v, want unavailable", err)
	}
}
