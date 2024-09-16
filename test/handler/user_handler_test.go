package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GonzaloC17/event-management-api/internal/domain"
	"github.com/GonzaloC17/event-management-api/internal/handler"
	"github.com/GonzaloC17/event-management-api/internal/infrastructure/repository"
	"github.com/GonzaloC17/event-management-api/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupUserRouter() *gin.Engine {
	repo := repository.NewInMemoryUserRepository()
	service := usecase.NewUserService(repo)
	userHandler := handler.NewUserHandler(service)

	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:userID", userHandler.GetUserByID)
	r.GET("/users", userHandler.GetAllUsers)
	return r
}

func TestCreateUserHandler(t *testing.T) {
	router := setupUserRouter()

	user := `{"id":1,"name":"Gonza","email":"gonzalo@example.com","user_role":"normal"}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(user))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdUser domain.User
	err := json.Unmarshal(w.Body.Bytes(), &createdUser)
	assert.NoError(t, err)
	assert.Equal(t, "Gonza", createdUser.Name)
}

func TestGetUserByIDHandler(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := usecase.NewUserService(repo)
	userHandler := handler.NewUserHandler(service)
	router := gin.Default()
	router.GET("/users/:userID", userHandler.GetUserByID)

	user := domain.User{ID: 1, Name: "Gonza", Email: "gonzalo@example.com", UserRole: domain.Admin}
	repo.Create(user)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var retrievedUser domain.User
	err := json.Unmarshal(w.Body.Bytes(), &retrievedUser)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, retrievedUser.Name)
}

func TestGetAllUsersHandler(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	service := usecase.NewUserService(repo)
	userHandler := handler.NewUserHandler(service)
	router := gin.Default()
	router.GET("/users", userHandler.GetAllUsers)

	user1 := domain.User{ID: 1, Name: "User1", Email: "user1@example.com", UserRole: domain.Normal}
	user2 := domain.User{ID: 2, Name: "User2", Email: "user2@example.com", UserRole: domain.Admin}
	repo.Create(user1)
	repo.Create(user2)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []domain.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Contains(t, users, user1)
	assert.Contains(t, users, user2)
}
