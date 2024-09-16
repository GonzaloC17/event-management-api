package repository

import (
	"errors"

	"github.com/GonzaloC17/event-management-api/internal/model"
)

type InMemoryUserRepository struct {
	users map[int]model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[int]model.User)}
}

func (r *InMemoryUserRepository) Create(user model.User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetByID(userID int) (model.User, error) {
	user, exists := r.users[userID]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) GetAll() []model.User {
	var userList []model.User
	for _, user := range r.users {
		userList = append(userList, user)
	}
	return userList
}
