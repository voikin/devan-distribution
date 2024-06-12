package usecase

import (
	"context"
	"github.com/voikin/devan-distribution/internal/entity"
)

type UserSevice interface {
	CreateUser(ctx context.Context, input entity.User) (int64, error)
	GenerateToken(username, password string) (string, string, error)
	RefreshToken(refreshToken string) (string, error)
	ParseToken(token string) (int, error)
}

type UseCase struct {
	userService UserSevice
}

func New(userService UserSevice) *UseCase {
	return &UseCase{
		userService: userService,
	}
}
