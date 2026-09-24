package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type UpdateDocumentUC struct {
	repo repository.DocumentRepository
}

type UpdateDocumentInput struct {
	Title   *string
	Content *string
}

func NewUpdateDocumentUC(repo repository.DocumentRepository) *UpdateDocumentUC {
	return &UpdateDocumentUC{
		repo: repo,
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
		log.Errorf(ctx, err, "failed to update document")
		return nil, err
	}

	return doc, nil
}
