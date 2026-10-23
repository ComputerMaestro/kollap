package repository

import (
	"context"
	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/google/uuid"
)

type EmbeddingRepository interface {
	Save(ctx context.Context, embedding *domain.Embedding) error
	FindSimilarDocumentIDs(ctx context.Context, workspaceID uuid.UUID, vector []float32, limit int) ([]uuid.UUID, error)
}
