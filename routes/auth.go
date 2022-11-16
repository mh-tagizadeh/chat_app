package routes

import (
    "github.com/gin-gonic/gin"
)

func AddAuthGroup(router *gin.RouterGroup) {
	auth := router.Group("auth")

	auth.POST("singin")

	auth.POST("signup")
}
