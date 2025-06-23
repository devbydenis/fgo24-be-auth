package router

import (
	c "auth/controller"

	"github.com/gin-gonic/gin"
)

func usersRouter(r *gin.RouterGroup){
	r.GET("/", c.GetUsersHandler)
}