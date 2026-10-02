package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type GetDocumentUC struct {
	repo      repository.DocumentRepository
	cacheRepo repository.DocumentCacheRepository
}

func NewGetDocumentUC(repo repository.DocumentRepository, cacheRepo repository.DocumentCacheRepository) *GetDocumentUC {
	return &GetDocumentUC{
		repo:      repo,
		cacheRepo: cacheRepo,
	}
}

func (uc *GetDocumentUC) Execute(ctx context.Context, id string) (*domain.Document, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		log.Errorf(ctx, err, "failed to parse document uuid")
		return nil, err
	}

	doc, err := uc.cacheRepo.FindByID(ctx, parsedUUID)
	if err != nil {
		log.Errorf(ctx, err, "cache miss: failed to find doc id %s in cache", id)
	} else {
		log.Infof(ctx, "cache hit: found doc in cache %s", id)
		return doc, nil
	}

	doc, err = uc.repo.FindByID(ctx, parsedUUID)
	if err != nil {
		log.Errorf(ctx, err, "failed to find the doc in the db %v", id)
		return nil, err
	}

	err = uc.cacheRepo.Save(ctx, doc)
	if err != nil {
		log.Errorf(ctx, err, "failed to update cache for doc %s", id)
	}

	return doc, nil
}
