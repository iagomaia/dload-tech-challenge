package repositories

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UserCollection = "users"
)

type UserDbe struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty"`
	Name      string              `bson:"name"`
	Email     string              `bson:"email"`
	Password  string              `bson:"password"`
	CreatedAt time.Time           `bson:"createdAt"`
	UpdatedAt *time.Time          `bson:"updatedAt"`
	DeletedAt *time.Time          `bson:"deletedAt"`
}

type UpdateUserDBPayload struct {
	Name      string     `bson:"name,omitempty"`
	UpdatedAt *time.Time `bson:"updatedAt,omitempty"`
}
