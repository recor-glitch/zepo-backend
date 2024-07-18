package user

import (
	"github.com/recor-glitch/zepo-backend/internal/domain/user"
	"github.com/recor-glitch/zepo-backend/internal/infrastructure/repository"
)

func GetUserByEmail(email string, repo *repository.UserRepository) (*user.User, error) {
	return repo.GetUserByEmail(email)
}