package semanticsearch

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"goa.design/clue/log"
)

type SemanticSearcher struct {
	embedder repository.Embedder
	repo     repository.EmbeddingRepository
}

func NewSemanticSearcher(embedder repository.Embedder, repo repository.EmbeddingRepository) *SemanticSearcher {
	return &SemanticSearcher{
		embedder: embedder,
		repo:     repo,
	}
}

func (s *SemanticSearcher) Search(ctx context.Context, query string) ([]*domain.Document, error) {
	embedding, err := s.embedder.Embed(ctx, query, map[string]any{})
	if err != nil {
		log.Errorf(ctx, err, "failed to embed the query")
		return nil, err
	}
	log.Infof(ctx, "embedded query %v", embedding)
	docs, err := s.repo.FindSimilarDocuments(ctx, embedding)
	if err != nil {
		log.Errorf(ctx, err, "failed to fetch documnts for query")
		return nil, err
	}
	return docs, nil
}
