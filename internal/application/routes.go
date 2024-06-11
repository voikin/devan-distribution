package application

import (
	"github.com/gin-gonic/gin"
	"github.com/voikin/devan-distribution/internal/controller"
)

func initRoutes(server *gin.Engine, ctrl *controller.Controller) {
	server.GET("/ping", ctrl.PingHandler)
}
