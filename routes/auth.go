package routes

import (
    "github.com/gin-gonic/gin"
)

func addAuthGroup(router *gin.RouterGroup) {
	auth := router.Group("auth")

	auth.POST("singin")

	auth.POST("signup")
}
