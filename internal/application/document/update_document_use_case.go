package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type UpdateDocumentUC struct {
	repo     repository.DocumentRepository
	embedder repository.Embedder
}

type UpdateDocumentInput struct {
	Title   *string
	Content *string
}

func NewUpdateDocumentUC(repo repository.DocumentRepository, embedder repository.Embedder) *UpdateDocumentUC {
	return &UpdateDocumentUC{
		repo:     repo,
		embedder: embedder,
	}
}

func (uc *UpdateDocumentUC) Execute(ctx context.Context, documentId string, updateInput UpdateDocumentInput) (*domain.Document, error) {
	docUUID, err := uuid.Parse(documentId)
	if err != nil {
		log.Errorf(ctx, err, "failed to parse document UUID")
		return nil, err
	}

	doc, err := uc.repo.FindByID(ctx, docUUID)
	if err != nil {
		log.Errorf(ctx, err, "failed to find the doc in the db")
		return nil, err
	}

	if updateInput.Title != nil {
		doc.Title = *updateInput.Title
	}

	if updateInput.Content != nil {
		doc.Content = *updateInput.Content
	}

	doc, err = uc.repo.Update(ctx, doc)
	if err != nil {
		log.Errorf(ctx, err, "error udpating document %v", documentId)
		return nil, err
	}

	bgCtx := context.WithoutCancel(ctx)
	go embedDocument(bgCtx, uc.repo, uc.embedder, doc)

	return doc, nil
}
