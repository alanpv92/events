package routes

import (
	"github.com/alanpv92/events/controller"
	"github.com/gin-gonic/gin"
)

func registerAuthRouteGroups(server *gin.Engine) {
	authRouterGroup := server.Group("/auth")

	authRouterGroup.POST("/login", controller.Login)
	authRouterGroup.POST("/register", controller.Register)

}
