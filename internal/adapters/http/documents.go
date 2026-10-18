package http

import (
	"context"
	"errors"
	"time"

	documents "github.com/ComputerMaestro/kollap/gen/documents"
	"github.com/ComputerMaestro/kollap/internal/application/document"
	"github.com/ComputerMaestro/kollap/internal/domain"
	"goa.design/clue/log"
)

// documents service example implementation.
// The example methods log the requests and return zero values.
type documentssrvc struct {
	createDocumentUC      *document.CreateDocumentUC
	getDocumentUC         *document.GetDocumentUC
	updateDocumentUC      *document.UpdateDocumentUC
	getSimilarDocumentsUC *document.GetSimilarDocumentsUC
}

// NewDocuments returns the documents service implementation.
func NewDocuments(
	createDocumentUC *document.CreateDocumentUC,
	getDocumentUC *document.GetDocumentUC,
	updateDocumentUC *document.UpdateDocumentUC,
	getSimilarDocumentsUC *document.GetSimilarDocumentsUC,
) documents.Service {
	return &documentssrvc{
		createDocumentUC:      createDocumentUC,
		getDocumentUC:         getDocumentUC,
		updateDocumentUC:      updateDocumentUC,
		getSimilarDocumentsUC: getSimilarDocumentsUC,
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

func (s *documentssrvc) UpdateDocument(ctx context.Context, p *documents.UpdateDocumentPayload) (res *documents.Document, err error) {
	updates := &document.UpdateDocumentInput{}
	if p.Content != nil {
		updates.Content = p.Content
	}
	if p.Title != nil {
		updates.Title = p.Title
	}
	w, err := s.updateDocumentUC.Execute(ctx, p.ID, *updates)
	if err != nil {
		log.Errorf(ctx, err, "failed to execute create workspace use case")
		switch {
		case errors.Is(err, domain.ErrDocumentNotFound):
			return nil, documents.MakeNotFound(err)

		case errors.Is(err, domain.ErrVersionConflict):
			return nil, documents.MakeConflict(err)

		default:
			return nil, err
		}
	}

	return toEndpointDoc(w), nil
}

// GetSimilarDocuments implements getSimilarDocuments.
func (s *documentssrvc) GetSimilarDocuments(ctx context.Context, p *documents.GetSimilarDocumentsPayload) (res []*documents.Document, err error) {
	docs, err := s.getSimilarDocumentsUC.Execute(ctx, p.Query)
	if err != nil {
		log.Errorf(ctx, err, "failed execution to get similar documents")
		return nil, err
	}
	for _, d := range docs {
		res = append(res, toEndpointDoc(d))
	}
	return res, nil
}

func toEndpointDoc(doc *domain.Document) *documents.Document {
	return &documents.Document{
		ID:          doc.ID.String(),
		Title:       doc.Title,
		Version:     doc.Version,
		Content:     doc.Content,
		WorkspaceID: doc.WorkspaceID.String(),
		CreatedAt:   doc.CreatedAt.Format(time.RFC3339),
	}
}
