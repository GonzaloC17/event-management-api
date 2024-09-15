package model

type UserRole string

type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	UserRole UserRole `json:"user_role"`
}
