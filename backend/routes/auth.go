package routes

import (
	"github.com/alanpv92/events/controller"
	"github.com/alanpv92/events/middlewares"
	"github.com/gin-gonic/gin"
)

func registerAuthRouteGroups(server *gin.Engine) {
	authRouterGroup := server.Group("/auth")

	authRouterGroup.POST("/login", controller.Login)
	authRouterGroup.POST("/register", controller.Register)
	authRouterGroup.POST("/refresh", middlewares.Authenticate, controller.RefreshToken)
	authRouterGroup.POST("/send-mail", middlewares.Authenticate, controller.SendVerificationMail)
    authRouterGroup.POST("/verify-email",middlewares.Authenticate,controller.VerifyOtp)
}
