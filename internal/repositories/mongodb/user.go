package mongodb

import (
	"context"
	"errors"
	"github.com/xpanvictor/g3_template.git/internal/domain/models"
	"github.com/xpanvictor/g3_template.git/internal/domain/repositories"
	"github.com/xpanvictor/g3_template.git/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type userRepository struct {
	collection mongo.Collection
}

func NewUserRepository(db mongo.Database, collection string) repositories.UserRepository {
	return &userRepository{
		collection: db.Collection(collection),
	}
}

func (ur userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := ur.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongoDriver.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (ur userRepository) List(ctx context.Context) ([]*models.User, error) {
	cursor, err := ur.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur userRepository) Create(ctx context.Context, user *models.User) error {
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}
	_, err := ur.collection.InsertOne(ctx, user)
	return err
}

func (ur userRepository) Update(ctx context.Context, user *models.User) error {
	_, err := ur.collection.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		}},
	)
	return err
}

func (ur userRepository) Delete(ctx context.Context, id string) error {
	_, err := ur.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
