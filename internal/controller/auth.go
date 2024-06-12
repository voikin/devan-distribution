package controller

import (
	"errors"
	"github.com/voikin/devan-distribution/internal/entity"
	"github.com/voikin/devan-distribution/internal/errs"
	"net/http"
	"time"

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
	var input entity.User

	if err := ctx.BindJSON(&input); err != nil {
		errs.NewErrorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := c.usecase.CreateUser(input)
	if err != nil {
		errs.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
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

	accessToken, refreshToken, err := c.usecase.GenerateToken(input.Username, input.Password)
	if err != nil {
		var myErr *errs.ErrorNotFound
		if errors.As(err, &myErr) {
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

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": accessToken,
		// "refreshToken": refreshToken, // Удалите эту строку, если вы не хотите отправлять refresh token в JSON ответе
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
	newAccessToken, err := c.usecase.RefreshToken(refreshToken)
	if err != nil {
		errs.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// Отправьте новый access токен обратно клиенту
	ctx.JSON(http.StatusOK, map[string]string{
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

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
