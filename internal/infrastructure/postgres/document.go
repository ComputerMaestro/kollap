package postgres

import (
	"context"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/infrastructure/postgres/models"
	"github.com/ComputerMaestro/kollap/internal/repository"
	"gorm.io/gorm"
)

type DocumentPostgresRepository struct {
	db *gorm.DB
}

func NewDocumentPostgresRepository(db *gorm.DB) *DocumentPostgresRepository {
	return &DocumentPostgresRepository{
		db: db,
	}
}

func (r *DocumentPostgresRepository) FindByID(ctx context.Context, id string) (*domain.Document, error) {
	var document models.Document
	if err := r.db.WithContext(ctx).First(&document, id).Error; err != nil {
		return nil, err
	}
	return &domain.Document{
		ID:          document.ID,
		WorkspaceID: document.WorkspaceID,
		Title:       document.Title,
		Content:     document.Content,
		Version:     document.Version,
		CreatedAt:   document.CreatedAt,
	}, nil
}

func (r *DocumentPostgresRepository) Find(ctx context.Context, filter repository.DocumentFilter) ([]*domain.Document, error) {
	var models []models.Document
	query := r.db.WithContext(ctx)

	query = applyDocumentFilter(query, filter)

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	documents := make([]*domain.Document, 0, len(models))

	for _, model := range models {
		documents = append(documents, toDomain(model))
	}

	return documents, nil
}

func (r *DocumentPostgresRepository) Create(ctx context.Context, document *domain.Document) (*domain.Document, error) {
	model := models.Document{
		ID:          document.ID,
		WorkspaceID: document.WorkspaceID,
		Title:       document.Title,
		Version:     document.Version,
		Content:     document.Content,
		CreatedAt:   document.CreatedAt,
		UpdatedAt:   document.UpdatedAt,
	}

	if err := r.db.
		WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return nil, err
	}

	return toDomain(model), nil
}

func toDomain(model models.Document) *domain.Document {
	return &domain.Document{
		ID:          model.ID,
		WorkspaceID: model.WorkspaceID,
		Title:       model.Title,
		Content:     model.Content,
		Version:     model.Version,
		CreatedAt:   model.CreatedAt,
	}
}

func applyDocumentFilter(query *gorm.DB, filter repository.DocumentFilter) *gorm.DB {
	if filter.WorkspaceID != nil {
		query = query.Where(
			"workspace_id = ?",
			*filter.WorkspaceID,
		)
	}

	if filter.OwnerID != nil {
		query = query.Where(
			"owner_id = ?",
			*filter.OwnerID,
		)
	}

	if filter.Title != nil {
		query = query.Where(
			"title = ?",
			*filter.Title,
		)
	}

	return query
}
