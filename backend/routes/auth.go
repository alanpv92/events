package routes

import (
	"github.com/gin-gonic/gin"
)

func registerAuthRouteGroups(server *gin.Engine) {
	authRouterGroup := server.Group("/auth")

	authRouterGroup.POST("/login", func(ctx *gin.Context) {})
	authRouterGroup.POST("/register", func(ctx *gin.Context) {})

}
