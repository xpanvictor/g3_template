package ports

import (
	"context"
	"go_graphql_gin/internal/core/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, id string) (*domain.User, error)
}

type NotificationService interface {
	NotifyUserCreation(ctx context.Context, userID string) error
}

type ProfileService interface {
	CreateProfile(ctx context.Context, profile *domain.Profile) error
}
