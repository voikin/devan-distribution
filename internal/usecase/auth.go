package usecase

import (
	"context"
	"github.com/voikin/devan-distribution/internal/DTO"
	"github.com/voikin/devan-distribution/internal/entity"
)

func (uc *UseCase) GenerateToken(ctx context.Context, username, password string) (string, string, error) {
	user, err := uc.userService.GetUserByUsernameWithPassCheck(ctx, username, password)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := uc.userService.GenerateTokens(ctx, user)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (uc *UseCase) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	user, err := uc.userService.GetUserByTokenWithCheck(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}

	newAccessToken, newRefreshToken, err := uc.userService.GenerateTokens(ctx, user)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (uc *UseCase) VerifyToken(ctx context.Context, accessToken string) (*entity.User, error) {
	user, err := uc.userService.GetUserByTokenWithCheck(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UseCase) CreateUser(ctx context.Context, userDTO DTO.CreateUser) (int64, error) {
	return uc.userService.CreateUser(ctx, userDTO)
}

func (uc *UseCase) GetRoles(ctx context.Context) ([]DTO.Role, error) {
	return uc.userService.GetRoles(ctx)
}
