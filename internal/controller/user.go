package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateUserHandler(ctx *gin.Context) {
	goCtx := ctx.Request.Context()
	user := c.usecase.CreateUser(goCtx)
	ctx.JSON(http.StatusCreated, user)
}
