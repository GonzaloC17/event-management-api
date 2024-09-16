package usecase

import "github.com/GonzaloC17/event-management-api/internal/domain"

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user domain.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUser(userID int) (domain.User, error) {
	return s.repo.GetByID(userID)
}

func (s *UserService) GetAllUsers() []domain.User {
	return s.repo.GetAll()
}
