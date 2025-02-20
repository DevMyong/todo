package user

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var (
	ErrorUserEmailAlreadyExists = errors.New("user: email already exists")
)

type User struct {
	ID           primitive.ObjectID
	Email        string
	Name         string
	Provider     Provider
	PasswordHash string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Provider string

const (
	ProviderLocal  Provider = "local"
	ProviderGoogle Provider = "google"
	ProviderGithub Provider = "github"
)

// NewUser is constructor to create User
func NewUser() *User {
	return &User{
		ID: primitive.NewObjectID(),
	}
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetProvider(provider Provider) {
	u.Provider = provider
}

func (u *User) SetPasswordHash(passwordHash string) {
	u.PasswordHash = passwordHash
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}
