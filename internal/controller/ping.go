package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
