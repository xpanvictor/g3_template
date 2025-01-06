package repositories

import (
	"context"
	"github.com/xpanvictor/g3_template.git/internal/domain/models"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}
