package repository

import (
	"errors"

	"github.com/GonzaloC17/event-management-api/internal/domain"
)

type InMemoryUserRepository struct {
	users map[int]domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[int]domain.User)}
}

func (r *InMemoryUserRepository) Create(user domain.User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetByID(userID int) (domain.User, error) {
	user, exists := r.users[userID]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) GetAll() []domain.User {
	var userList []domain.User
	for _, user := range r.users {
		userList = append(userList, user)
	}
	return userList
}
