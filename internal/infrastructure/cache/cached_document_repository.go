package cache

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type CachedDocumentRepository struct {
	db    repository.DocumentRepository
	cache repository.DocumentCacheRepository
}

func NewCachedDocumentRepository(db repository.DocumentRepository, cache repository.DocumentCacheRepository) *CachedDocumentRepository {
	return &CachedDocumentRepository{
		db:    db,
		cache: cache,
	}
}

func (r *CachedDocumentRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	doc, err := r.cache.Get(ctx, id)
	if err != nil {
		log.Errorf(ctx, err, "cache miss: failed to find doc id %s in cache", id)
	} else {
		return doc, nil
	}

	doc, err = r.db.FindByID(ctx, id)
	if err != nil {
		log.Errorf(ctx, err, "failed to find the doc in the db %v", id)
		return nil, err
	}

	err = r.cache.Set(ctx, doc)
	if err != nil {
		log.Errorf(ctx, err, "failed to update cache for doc %s", id)
	}

	return &domain.Document{
		ID:          doc.ID,
		WorkspaceID: doc.WorkspaceID,
		Title:       doc.Title,
		Content:     doc.Content,
		Version:     doc.Version,
		CreatedAt:   doc.CreatedAt,
		Embedding:   doc.Embedding,
	}, nil
}

func (r *CachedDocumentRepository) Find(ctx context.Context, filter repository.DocumentFilter) ([]*domain.Document, error) {
	return r.db.Find(ctx, filter)
}

func (r *CachedDocumentRepository) Create(ctx context.Context, document *domain.Document) (*domain.Document, error) {
	return r.db.Create(ctx, document)
}

func (r *CachedDocumentRepository) Update(ctx context.Context, document *domain.Document) (*domain.Document, error) {
	doc, err := r.db.Update(ctx, document)
	if err != nil {
		log.Errorf(ctx, err, "failed to update document in db")
		return nil, err
	}

	err = r.cache.Unlink(ctx, doc.ID)
	if err != nil {
		log.Errorf(ctx, err, "failed cache invalidation")
	}
	return doc, nil
}
