package usecase

import (
	"context"

	"github.com/voikin/devan-distribution/internal/config"
	"github.com/voikin/devan-distribution/internal/entity"
)

type UserService interface {
	CreateUser(ctx context.Context, user entity.User) (int64, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserByID(ctx context.Context, userId int64) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (bool, error)
	DeleteUser(ctx context.Context, userId int64) (bool, error)
}

type UseCase struct {
	cfg         config.Config
	userService UserService
}

func New(cfg config.Config, userService UserService) *UseCase {
	return &UseCase{
		userService: userService,
	}
}
