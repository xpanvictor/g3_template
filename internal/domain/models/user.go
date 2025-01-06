package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	UserCollection = "users"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
