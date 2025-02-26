package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User struct is model for user
type User struct {
	ID           primitive.ObjectID `db:"id"`
	Email        string             `db:"email"`
	Name         string             `db:"name"`
	Provider     string             `db:"provider"`
	PasswordHash string             `db:"password_hash"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
