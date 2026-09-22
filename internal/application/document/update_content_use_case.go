package document

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/repository"
)

type UpdateContentUC struct {
	repo repository.DocumentRepository
}

func NewUpdateContentUC(repo repository.DocumentRepository) *UpdateContentUC {
	return &UpdateContentUC{
		repo: repo,
	}
}

func (uc *UpdateContentUC) Execute(ctx context.Context, content string) (*domain.Document, error) {
	// workspaceUUID, err := uuid.Parse(workspaceID)
	// if err != nil {
	// 	log.Errorf(ctx, err, "failed to parse UUID")
	// 	return nil, err
	// }
	// Document := &domain.Document{
	// 	ID:          uuid.New(),
	// 	Title:       name,
	// 	Version:     1,
	// 	Content:     content,
	// 	WorkspaceID: workspaceUUID,
	// }
	// return uc.repo.Create(ctx, Document)
	return nil, nil
}
