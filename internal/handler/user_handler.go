package handler

import (
	"net/http"
	"strconv"

	"github.com/GonzaloC17/event-management-api/internal/model"
	"github.com/GonzaloC17/event-management-api/internal/service"
	"github.com/GonzaloC17/event-management-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := service.CreateUser(newUser)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := service.GetUser(userID)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	users := service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
