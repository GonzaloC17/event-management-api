package repository_test

import (
	"testing"

	"github.com/GonzaloC17/event-management-api/internal/domain"
	"github.com/GonzaloC17/event-management-api/internal/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()

	user := domain.User{ID: 1, Name: "Gonza", Email: "gonzalo@example.com", UserRole: domain.Normal}
	err := repo.Create(user)
	assert.NoError(t, err)

	// Test crear usuario duplicado
	err = repo.Create(user)
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
}

func TestGetUserByID(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()

	user := domain.User{ID: 1, Name: "Gonza", Email: "gonzalo@example.com", UserRole: domain.Admin}
	repo.Create(user)

	retrievedUser, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)

	_, err = repo.GetByID(2)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}

func TestGetAllUsers(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()

	user1 := domain.User{ID: 1, Name: "User1", Email: "user1@example.com", UserRole: domain.Normal}
	user2 := domain.User{ID: 2, Name: "User2", Email: "user2@example.com", UserRole: domain.Admin}
	repo.Create(user1)
	repo.Create(user2)

	users := repo.GetAll()
	assert.Len(t, users, 2)
	assert.Contains(t, users, user1)
	assert.Contains(t, users, user2)
}
