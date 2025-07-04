package main

import (
	"auth/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "FGO24-BE-Auth",
		})
	})

	router.CombineRouters(r)

	r.Run()
}
