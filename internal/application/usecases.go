package application

import (
	"github.com/voikin/devan-distribution/internal/usecase"
)

func createUseCase(services *services) *usecase.UseCase {
	return usecase.New(services.userService)
}
