package service

import (
	"github.com/GonzaloC17/event-management-api/internal/model"
	"github.com/GonzaloC17/event-management-api/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user model.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUser(userID int) (model.User, error) {
	return s.repo.GetByID(userID)
}

func (s *UserService) GetAllUsers() []model.User {
	return s.repo.GetAll()
}
