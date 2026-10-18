package repository

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
)

type EmbeddingRepository interface {
	FindSimilarDocuments(ctx context.Context, embedding []float32) ([]*domain.Document, error)
}
