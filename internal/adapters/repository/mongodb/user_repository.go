package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go_graphql_gin/internal/adapters/repository/models"
	"go_graphql_gin/internal/core/domain"
)

type UserRepository struct {
	collection *mongo.Collection
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	var model models.UserModel
	objectID, _ := primitive.ObjectIDFromHex(id)

	err := r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&model)
	if err != nil {
		return nil, err
	}

	return model.ToDomain(), nil
}
