package controller

import (
	"context"
	"github.com/voikin/devan-distribution/internal/entity"
)

// UseCase TODO are we really need this ?
type UseCase interface {
	CreateUser(ctx context.Context, input entity.User) (int64, error)
	GenerateToken(username, password string) (string, string, error)
	RefreshToken(refreshToken string) (string, error)
	ParseToken(token string) (int, error)
}

type Controller struct {
	usecase UseCase
}

func New(usecase UseCase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}
