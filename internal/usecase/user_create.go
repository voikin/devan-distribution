package usecase

import "context"

func (uc *UseCase) CreateUser(ctx context.Context) error {
	return uc.userService.CreateUser(ctx)
}
