package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required"`
	Password  string             `bson:"password" json:"password" validate:"required"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	InsertOne(ctx context.Context, u *User) (*User, error)
	FindOne(ctx context.Context, id string) (*User, error)
	GetAllWithPage(ctx context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]User, int64, error)
	GetByCredential(ctx context.Context, username string, password string) (*User, error)
	UpdateOne(ctx context.Context, user *User, id string) (*User, error)
}

type UserUsecase interface {
	InsertOne(ctx context.Context, u *User) (*User, error)
	GetAllWithPage(ctx context.Context, rp int64, p int64, filter interface{}, setsort interface{}) ([]User, int64, error)
	FindOne(ctx context.Context, id string) (*User, error)
	UpdateOne(ctx context.Context, user *User, id string) (*User, error)
}
