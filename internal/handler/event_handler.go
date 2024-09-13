package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "lista de eventos",
	})
}

func CreateEvents(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "evento creado",
	})
}
