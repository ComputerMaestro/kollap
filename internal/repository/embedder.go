package repository

import (
	"context"
)

type Embedder interface {
	Embed(ctx context.Context, content string, metadata map[string]any) ([]float32, error)
}

type DocumentIndexer interface {
	Index(ctx context.Context, documentID, workspaceID string, title, content string) error
}
