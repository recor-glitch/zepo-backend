package user

import (
	"github.com/recor-glitch/zepo-backend/internal/domain/user"
	"github.com/recor-glitch/zepo-backend/internal/infrastructure/repository"
)

func CreateUser(u *user.User, repo *repository.UserRepository) error {
	return repo.Create(u)
}