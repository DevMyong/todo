package user

import (
	"github.com/devmyong/todo/internal/adapter/db/model"
)

// Repository is an interface for user repository
type Repository interface {
	Create(user *User) error
	FindAll() ([]model.User, error)
	FindByID(id string) (model.User, error)
	GetByEmail(email string) (*User, error)
	Update(id string, user model.User) error
	Delete(id string) error
}

// EventPublisher is an interface for user event publisher
type EventPublisher interface {
	PublishUserCreated(id string) error
}

// OAuthClient is an interface for OAuth client
type OAuthClient interface {
	VerifyToken(token string) (string, error)
}

type OAuthUserInfo struct {
	Email string
	Name  string
}

func NewOAuthUserInfo(email, name string) *OAuthUserInfo {
	return &OAuthUserInfo{
		Email: email,
		Name:  name,
	}
}

func (u *OAuthUserInfo) VerifyToken(token string) (string, error) {
	return u.Email, nil
}
