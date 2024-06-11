package usecase

import "context"

type UserSevice interface {
	CreateUser(ctx context.Context) error
}

type UseCase struct {
	userService UserSevice
}

func New(userService UserSevice) *UseCase {
	return &UseCase{
		userService: userService,
	}
}
