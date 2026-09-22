package http

import (
	"context"
	"time"

	documents "github.com/ComputerMaestro/kollap/gen/documents"
	"github.com/ComputerMaestro/kollap/internal/application/document"
	"goa.design/clue/log"
)

// documents service example implementation.
// The example methods log the requests and return zero values.
type documentssrvc struct {
	createDocumentUC *document.CreateDocumentUC
	getDocumentUC    *document.GetDocumentUC
}

// NewDocuments returns the documents service implementation.
func NewDocuments(
	createDocumentUC *document.CreateDocumentUC,
	getDocumentUC *document.GetDocumentUC,
) documents.Service {
	return &documentssrvc{
		createDocumentUC: createDocumentUC,
		getDocumentUC:    getDocumentUC,
	}
}

// GetDocument implements getDocument.
func (s *documentssrvc) GetDocument(ctx context.Context, p *documents.GetDocumentPayload) (res *documents.Document, err error) {
	doc, err := s.getDocumentUC.Execute(ctx, p.ID)
	if err != nil {
		log.Errorf(ctx, err, "error fetching document details")
		return
	}
	res = &documents.Document{
		ID:        doc.ID.String(),
		Title:     doc.Title,
		Version:   doc.Version,
		Content:   doc.Content,
		CreatedAt: doc.CreatedAt.Format(time.RFC3339),
	}
	return
}

func (s *documentssrvc) CreateDocument(ctx context.Context, p *documents.CreateDocumentPayload) (res *documents.Document, err error) {
	content := ""
	if p.Content != nil {
		content = *p.Content
	}
	w, err := s.createDocumentUC.Execute(ctx, p.Title, p.WorkspaceID, content)
	if err != nil {
		log.Errorf(ctx, err, "failed to execute create workspace use case")
		return nil, err
	}
	res = &documents.Document{
		ID:          w.ID.String(),
		Title:       w.Title,
		Version:     w.Version,
		Content:     w.Content,
		WorkspaceID: w.WorkspaceID.String(),
		CreatedAt:   w.CreatedAt.Format(time.RFC3339),
	}
	return
}
