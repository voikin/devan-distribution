package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine, ctrl *Controller) {
	server.GET("/ping", ctrl.PingHandler)

	auth := server.Group("/auth")
	{
		auth.POST("/sign-up", ctrl.signUp)
		auth.POST("/sign-in", ctrl.signIn)
		auth.GET("/refresh", ctrl.refreshToken)
		auth.POST("/logout", ctrl.logout)
		auth.GET("/roles", ctrl.roles)
	}

	api := server.Group("/api", ctrl.userIdentity)
	{
		api.POST("/orders")
	}
}
