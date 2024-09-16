package handler

import (
	"net/http"
	"strconv"

	"github.com/GonzaloC17/event-management-api/internal/domain"
	"github.com/GonzaloC17/event-management-api/internal/usecase"
	"github.com/GonzaloC17/event-management-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *usecase.UserService
}

// Constructor para UserHandler
func NewUserHandler(userService *usecase.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := h.userService.CreateUser(newUser)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.userService.GetUser(userID)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users := h.userService.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
