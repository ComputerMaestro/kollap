package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type CreateDocumentUC struct {
	repo    repository.DocumentRepository
	indexer repository.DocumentIndexer
}

func NewCreateDocumentUC(repo repository.DocumentRepository, indexer repository.DocumentIndexer) *CreateDocumentUC {
	return &CreateDocumentUC{
		repo:    repo,
		indexer: indexer,
	}
}

func (uc *CreateDocumentUC) Execute(ctx context.Context, name string, workspaceID string, content string) (*domain.Document, error) {
	workspaceUUID, err := uuid.Parse(workspaceID)
	if err != nil {
		log.Errorf(ctx, err, "failed to parse UUID")
		return nil, err
	}
	Document := &domain.Document{
		ID:          uuid.New(),
		Title:       name,
		Version:     1,
		Content:     content,
		WorkspaceID: workspaceUUID,
	}

	doc, err := uc.repo.Create(ctx, Document)
	if err != nil {
		log.Errorf(ctx, err, "error creating document %v", Document)
		return nil, err
	}

	if uc.indexer != nil {
		bgCtx := context.WithoutCancel(ctx)
		go func() {
			if err := uc.indexer.Index(bgCtx, doc.ID.String(), doc.WorkspaceID.String(), doc.Title, doc.Content); err != nil {
				log.Errorf(bgCtx, err, "failed to index document %v", doc.ID)
			}
		}()
	}

	return doc, nil
}
