package postgres

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/infrastructure/postgres/models"
	"gorm.io/gorm"
)

type WorkspacePostgresRepository struct {
	db *gorm.DB
}

func NewWorkspacePostgresRepository(db *gorm.DB) *WorkspacePostgresRepository {
	return &WorkspacePostgresRepository{
		db: db,
	}
}

func (r *WorkspacePostgresRepository) FindByID(ctx context.Context, id string) (*domain.Workspace, error) {
	var workspace models.WorkspaceModel
	if err := r.db.WithContext(ctx).First(&workspace, id).Error; err != nil {
		return nil, err
	}
	return &domain.Workspace{
		ID:        workspace.ID,
		OwnerID:   workspace.OwnerID,
		Name:      workspace.Name,
		CreatedAt: workspace.CreatedAt,
	}, nil
}

func (r *WorkspacePostgresRepository) Create(ctx context.Context, w *domain.Workspace) (*domain.Workspace, error) {
	res := r.db.WithContext(ctx).Create(w)
	if res.Error != nil {
		return nil, res.Error
	}
	return w, nil
}
