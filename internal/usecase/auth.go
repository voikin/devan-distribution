package usecase

import (
	"context"
	"github.com/voikin/devan-distribution/internal/entity"
)

func (uc *UseCase) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	return uc.userService.CreateUser(ctx, user)
}

func (uc *UseCase) GenerateToken(username, password string) (string, string, error) {
	return uc.userService.GenerateToken(username, password)
}

func (uc *UseCase) RefreshToken(refreshToken string) (string, error) {
	return uc.userService.RefreshToken(refreshToken)
}

func (uc *UseCase) ParseToken(token string) (int, error) {
	return uc.userService.ParseToken(token)
}
