package service

import (
	"errors"

	"github.com/GonzaloC17/event-management-api/internal/model"
)

var (
	userStore = make(map[int]model.User)
)

func CreateUser(user model.User) error {
	if _, exists := userStore[user.ID]; exists {
		return errors.New("user already exists")
	}
	userStore[user.ID] = user
	return nil
}

func GetUser(userID int) (model.User, error) {
	user, exists := userStore[userID]
	if !exists {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func GetAllUsers() []model.User {
	var userList []model.User
	for _, user := range userStore {
		userList = append(userList, user)
	}
	return userList
}
