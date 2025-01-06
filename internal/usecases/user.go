package usecases

import (
	"context"
	"github.com/xpanvictor/g3_template.git/internal/domain/models"
	"github.com/xpanvictor/g3_template.git/internal/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*models.User, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *UserUseCase) CreateUser(ctx context.Context, name string, email string) (*models.User, error) {
	user := &models.User{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Time{},
	}
	err := uc.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) ListUsers(ctx context.Context) ([]*models.User, error) {
	return uc.repo.List(ctx)
}
