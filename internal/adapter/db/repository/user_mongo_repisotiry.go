package repository

import (
	"context"
	"github.com/devmyong/todo/backend/backend/internal/adapter/db/model"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepositoryMongo is struct to connect db
type UserRepositoryMongo struct {
	collection *mongo.Collection
}

// NewUserRepositoryMongo is method to create new user repository
func NewUserRepositoryMongo(client *mongo.Client) *UserRepositoryMongo {
	collection := client.Database("todo").Collection("users") // todo: replace to mongo db name
	return &UserRepositoryMongo{collection: collection}
}

// Create is method to create user
func (u *UserRepositoryMongo) Create(user *model.User) error {
	_, err := u.collection.InsertOne(context.Background(), user)
	return err
}

// FindAll is method to find all user
func (u *UserRepositoryMongo) FindAll() ([]model.User, error) {
	return []model.User{}, nil
}

// FindByID is method to find user by id
func (u *UserRepositoryMongo) FindByID(id string) (model.User, error) {
	return model.User{}, nil
}

// Update is method to update user
func (u *UserRepositoryMongo) Update(id string, user model.User) error {
	return nil
}

// Delete is method to delete user
func (u *UserRepositoryMongo) Delete(id string) error {
	return nil
}
