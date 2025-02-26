package repository

import (
	"database/sql"
	"github.com/devmyong/todo/internal/adapter/db/model"
	"github.com/devmyong/todo/internal/domain/user"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Create(u *user.User) error {
	um := model.User{
		ID:           u.ID,
		Email:        u.Email,
		Name:         u.Name,
		Provider:     string(u.Provider),
		PasswordHash: u.PasswordHash,
	}

	_, err := r.db.Exec(
		"INSERT INTO users (id, email, name, provider, password_hash) VALUES ($1, $2, $3, $4, $5)",
		um.ID, um.Email, um.Name, um.Provider, um.PasswordHash,
	)
	return err
}

func (r *PostgresUserRepository) GetByEmail(email string) (*user.User, error) {
	var um model.User
	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).
		Scan(&um.ID, &um.Email, &um.Name, &um.Provider, &um.PasswordHash, &um.CreatedAt, &um.UpdatedAt)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		ID:           um.ID,
		Email:        um.Email,
		Name:         um.Name,
		Provider:     user.Provider(um.Provider),
		PasswordHash: um.PasswordHash,
		CreatedAt:    um.CreatedAt,
		UpdatedAt:    um.UpdatedAt,
	}
	return u, nil
}
