package repository

import "context"

type DocumentRepository interface {
	GetDocument(ctx context.Context)
}
