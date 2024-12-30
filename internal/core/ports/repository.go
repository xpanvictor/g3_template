package ports

import "go_graphql_gin/internal/core/domain"

type UserRepository interface {
	FindByID(id string) (*domain.User, error)
	Create(user *domain.User) error
	// Other methods...
}
