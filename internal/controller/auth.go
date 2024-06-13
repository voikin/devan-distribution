package controller

import (
	"errors"
	"github.com/voikin/devan-distribution/internal/DTO"
	"net/http"
	"time"

	"github.com/voikin/devan-distribution/internal/errs"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body entity.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errs.errorResponse
// @Failure 500 {object} errs.errorResponse
// @Failure default {object} errs.errorResponse
// @Router /auth/sign-up [post]
func (c *Controller) signUp(ctx *gin.Context) {
	var input DTO.CreateUser

	if err := ctx.BindJSON(&input); err != nil {
		errs.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := c.usecase.CreateUser(ctx.Request.Context(), input)
	if err != nil {
		errs.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errs.errorResponse
// @Failure 500 {object} errs.errorResponse
// @Failure default {object} errs.errorResponse
// @Router /auth/sign-in [post]
func (c *Controller) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		errs.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := c.usecase.GenerateToken(ctx.Request.Context(), input.Username, input.Password)
	if err != nil {
		var errNotFound *errs.ErrorNotFound
		if errors.As(err, &errNotFound) {
			errs.NewErrorResponse(ctx, http.StatusNotFound, err.Error())
			return
		}
		errs.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Создание HTTP-only cookie для refresh token
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true, // использовать только при передаче через HTTPS
		Path:     "/",
	})

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
		"role":        "admin", //TODO: think about it
	})
}

func (c *Controller) refreshToken(ctx *gin.Context) {
	// Получение refresh token из cookie
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		errs.NewErrorResponse(ctx, http.StatusUnauthorized, "refresh token cookie is missing")
		return
	}

	if refreshToken == "" {
		errs.NewErrorResponse(ctx, http.StatusUnauthorized, "refresh token is empty")
		return
	}

	// Используйте новый метод RefreshToken для обновления access токена
	newAccessToken, newRefreshToken, err := c.usecase.RefreshToken(ctx.Request.Context(), refreshToken)
	if err != nil {
		errs.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    newRefreshToken,
		HttpOnly: true,
		Secure:   true, // использовать только при передаче через HTTPS
		Path:     "/",
	})

	// Отправьте новый access токен обратно клиенту
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessToken,
	})
}

// @Summary Logout
// @Tags auth
// @Description logout
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Failure 400,404 {object} errs.errorResponse
// @Failure 500 {object} errs.errorResponse
// @Failure default {object} errs.errorResponse
// @Router /auth/logout [post]
func (c *Controller) logout(ctx *gin.Context) {
	// Удаление HTTP-only cookie, установив срок его действия в прошлом
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true, // использовать только при передаче через HTTPS
		Path:     "/",
	})

	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// @Summary Roles
// @Tags auth
// @Description Get all roles
// @ID roles
// @Accept  json
// @Produce  json
// @Success 200 {array} DTO.Role
// @Failure 400,404 {object} errs.errorResponse
// @Failure 500 {object} errs.errorResponse
// @Failure default {object} errs.errorResponse
// @Router /auth/roles [get]
func (c *Controller) roles(ctx *gin.Context) {
	roles, err := c.usecase.GetRoles(ctx)
	if err != nil {
		errs.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}
