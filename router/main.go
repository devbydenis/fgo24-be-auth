package router

import "github.com/gin-gonic/gin"

func CombineRouters(r *gin.Engine){
	authRouter(r.Group("/auth"))
	usersRouter(r.Group("/users"))
}