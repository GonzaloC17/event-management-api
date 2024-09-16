package domain

type UserRepository interface {
	Create(user User) error
	GetByID(id int) (User, error)
	GetAll() []User
}
