package usecase

import (
	"context"
	"github.com/voikin/devan-distribution/internal/DTO"
	"github.com/voikin/devan-distribution/internal/config"
	"github.com/voikin/devan-distribution/internal/entity"
)

type UserService interface {
	CreateUser(ctx context.Context, user DTO.CreateUser) (int64, error)
	GetUserByUsernameWithPassCheck(ctx context.Context, username string, password string) (*entity.User, error)
	GenerateTokens(ctx context.Context, user *entity.User) (string, string, error)
	GetUserByTokenWithCheck(ctx context.Context, token string) (*entity.User, error)
	GetRoles(ctx context.Context) ([]DTO.Role, error)
}

type UseCase struct {
	cfg         config.Config
	userService UserService
}

func New(cfg config.Config, userService UserService) *UseCase {
	return &UseCase{
		userService: userService,
		cfg:         cfg,
	}
}
