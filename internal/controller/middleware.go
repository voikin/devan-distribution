package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/voikin/devan-distribution/internal/errs"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user"
)

func (c *Controller) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		errs.NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		errs.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		errs.NewErrorResponse(ctx, http.StatusUnauthorized, "token is empty")
		return
	}

	user, err := c.usecase.VerifyToken(ctx.Request.Context(), headerParts[1])
	if err != nil {
		errs.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, user)
}
