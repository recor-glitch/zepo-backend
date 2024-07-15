package repository

import (
	"github.com/recor-glitch/zepo-backend/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRespository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByID(id string) (*user.User, error) {
	var u user.User
	err := r.DB.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Create(u *user.User) error {
	return r.DB.Create(u).Error
}
