package repository

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
)

type Embedder interface {
	Embed(ctx context.Context, doc *domain.Document) ([]float32, error)
}
