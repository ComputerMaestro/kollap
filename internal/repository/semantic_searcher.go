package repository

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
)

type SemanticSearcher interface {
	Search(ctx context.Context, query string) ([]*domain.Document, error)
}
