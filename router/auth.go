package router

import (
	c "auth/controller"

	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	r.POST("/register", c.RegisterHandler)
	r.POST("/login", c.LoginHandler)
	r.POST("/reset-password", c.ResetPasswordHandler)
	
	r.GET("/forgot-password", c.ForgotPasswordHandler)
}