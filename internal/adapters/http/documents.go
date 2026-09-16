package http

import (
	"context"

	documents "github.com/ComputerMaestro/kollap/gen/documents"
	"goa.design/clue/log"
)

// documents service example implementation.
// The example methods log the requests and return zero values.
type documentssrvc struct{}

// NewDocuments returns the documents service implementation.
func NewDocuments() documents.Service {
	return &documentssrvc{}
}

// GetDocument implements getDocument.
func (s *documentssrvc) GetDocument(ctx context.Context, p *documents.GetDocumentPayload) (res *documents.Document, err error) {
	res = &documents.Document{}
	log.Printf(ctx, "documents.getDocument")
	return
}
