package service

import "github.com/GonzaloC17/event-management-api/internal/model"

var (
	users = make(map[string]model.User)
)

/*func CreateUser(user model.User) error{
	if _, exists := users[user.ID]; exists{
		return errors.New()
	}
}*/
