package domain

import "github.com/google/uuid"

type Embedding struct {
	DocumentID  uuid.UUID
	WorkspaceID uuid.UUID
	Model       string
	Vector      []float32
}
