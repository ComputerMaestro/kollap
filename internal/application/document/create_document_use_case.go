package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"github.com/google/uuid"
	"goa.design/clue/log"
)

type CreateDocumentUC struct {
	repo     repository.DocumentRepository
	embedder repository.Embedder
}

func NewCreateDocumentUC(repo repository.DocumentRepository, embedder repository.Embedder) *CreateDocumentUC {
	return &CreateDocumentUC{
		repo:     repo,
		embedder: embedder,
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

	bgCtx := context.WithoutCancel(ctx)
	go embedDocument(bgCtx, uc.repo, uc.embedder, doc)

	return doc, nil
}

func embedDocument(ctx context.Context, documentRepo repository.DocumentRepository, embedder repository.Embedder, doc *domain.Document) {
	embedding, err := embedder.Embed(ctx, doc)
	if err != nil {
		log.Errorf(ctx, err, "failed to embed document %v", doc.ID)
		return
	}
	doc.Embedding = embedding
	_, err = documentRepo.Update(ctx, doc)
	if err != nil {
		log.Errorf(ctx, err, "failed to update embedding in the doc %v", doc.ID)
	}
}
