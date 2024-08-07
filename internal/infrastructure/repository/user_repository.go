package repository

import (
	"database/sql"

	"github.com/recor-glitch/zepo-backend/internal/domain/user"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRespository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByID(id string) (*user.User, error) {
	var u user.User
	query := `SELECT id, name, email, createdat, image, role FROM users WHERE id=$1`
	err := r.DB.QueryRow(query, id).Scan(&u.Id, &u.Name, &u.Email, &u.CreatedAt, &u.Image, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*user.User, error) {
	var u user.User
	query := `SELECT id, name, email, created_at, image, role FROM users WHERE email=$1`
	err := r.DB.QueryRow(query, email).Scan(&u.Id, &u.Name, &u.Email, &u.CreatedAt, &u.Image, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Create(u *user.User) error {
	// CREATE USER
	create_query := `INSERT INTO users (id, name, email, image) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(create_query, u.Id, u.Name, u.Email, u.Image)
	return err
}
