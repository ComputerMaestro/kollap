package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ComputerMaestro/kollap/internal/domain"
	"github.com/ComputerMaestro/kollap/internal/infrastructure/redis/models"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type DocumentRedisRepository struct {
	client *redis.Client
}

func NewDocumentRedisRepository(client *redis.Client) *DocumentRedisRepository {
	return &DocumentRedisRepository{
		client: client,
	}
}

func (r *DocumentRedisRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	res, err := r.client.Get(ctx, id.String()).Result()
	if err != nil {
		return nil, err
	}
	doc := &models.CacheDocumentModel{}
	err = json.Unmarshal([]byte(res), doc)
	return toDomain(doc), nil
}

func (r *DocumentRedisRepository) Set(ctx context.Context, doc *domain.Document) error {
	jsonData, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, doc.ID.String(), jsonData, 30*time.Second).Err()
}

func (r *DocumentRedisRepository) Unlink(ctx context.Context, id uuid.UUID) error {
	err := r.client.Unlink(ctx, id.String()).Err()
	if err != nil {
		return err
	}
	return nil
}

func toDomain(doc *models.CacheDocumentModel) *domain.Document {
	return &domain.Document{
		ID:          doc.ID,
		Title:       doc.Title,
		WorkspaceID: doc.WorkspaceID,
		Content:     doc.Content,
		Version:     doc.Version,
		CreatedAt:   doc.CreatedAt,
		UpdatedAt:   doc.UpdatedAt,
	}
}
