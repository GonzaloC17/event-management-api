package repository

import "github.com/GonzaloC17/event-management-api/internal/model"

type UserRepository interface {
	Create(user model.User) error
	GetByID(id int) (model.User, error)
	GetAll() []model.User
}
