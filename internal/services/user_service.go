package services

import (
	"github.com/cbitbaly/internal/models"
	"github.com/cbitbaly/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers(page, pageSize int) ([]models.User, int64, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(int(id))
}

func (s *UserService) GetUserWithPosts(id uint) (*models.User, error) {
	return s.repo.FindWithPosts(id)
}

func (s *UserService) UpdateUser(id uint, data map[string]any) error {
	return s.repo.Update(id, data)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
