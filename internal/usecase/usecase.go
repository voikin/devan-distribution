package usecase

import (
	"context"
	"github.com/voikin/devan-distribution/internal/entity"
)

type UserSevice interface {
	CreateUser(ctx context.Context, input entity.User) (int64, error)
}

type UseCase struct {
	userService UserSevice
}

func New(userService UserSevice) *UseCase {
	return &UseCase{
		userService: userService,
	}
}
