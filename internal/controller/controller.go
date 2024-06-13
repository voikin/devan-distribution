package controller

import (
	"context"
	"github.com/voikin/devan-distribution/internal/DTO"
	"github.com/voikin/devan-distribution/internal/entity"
)

// UseCase TODO are we really need this ?
type UseCase interface {
	CreateUser(ctx context.Context, input DTO.CreateUser) (int64, error)
	GenerateToken(ctx context.Context, username, password string) (string, string, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, string, error)
	VerifyToken(ctx context.Context, accessToken string) (*entity.User, error)
	GetRoles(ctx context.Context) ([]DTO.Role, error)
}

type Controller struct {
	usecase UseCase
}

func New(usecase UseCase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}
