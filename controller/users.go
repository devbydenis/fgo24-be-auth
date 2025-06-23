package controller

import (
	m "auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(c *gin.Context) {
	var userList []gin.H
	
	for _, user := range m.Users {
		userList = append(userList, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data users",
		"users":   userList,
	})
}