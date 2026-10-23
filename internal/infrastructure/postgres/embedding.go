package postgres

import (
	"context"
	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/infrastructure/postgres/models"
	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmbeddingPostgresRepository struct{ db *gorm.DB }

func NewEmbeddingPostgresRepository(db *gorm.DB) *EmbeddingPostgresRepository {
	return &EmbeddingPostgresRepository{db: db}
}

func (r *EmbeddingPostgresRepository) Save(ctx context.Context, embedding *domain.Embedding) error {
	model := models.Embedding{DocumentID: embedding.DocumentID, WorkspaceID: embedding.WorkspaceID, Model: embedding.Model, Vector: pgvector.NewVector(embedding.Vector)}
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "document_id"}}, DoUpdates: clause.AssignmentColumns([]string{"workspace_id", "model", "vector", "updated_at"})}).Create(&model).Error
}

func (r *EmbeddingPostgresRepository) FindSimilarDocumentIDs(ctx context.Context, workspaceID uuid.UUID, vector []float32, limit int) ([]uuid.UUID, error) {
	var rows []models.Embedding
	err := r.db.WithContext(ctx).Where("workspace_id = ?", workspaceID).Order(clause.Expr{SQL: "vector <-> ?", Vars: []any{pgvector.NewVector(vector)}}).Limit(limit).Find(&rows).Error
	if err != nil {
		return nil, err
	}
	ids := make([]uuid.UUID, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.DocumentID)
	}
	return ids, nil
}
