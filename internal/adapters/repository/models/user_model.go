package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_graphql_gin/internal/core/domain"
	"time"
)

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
}

// ToDomain Convert between domain and database models
func (m *UserModel) ToDomain() *domain.User {
	return &domain.User{
		ID:        m.ID.Hex(),
		Username:  m.Username,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	}
}
