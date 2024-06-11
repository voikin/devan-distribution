package application

import (
	userServicePkg "github.com/voikin/devan-distribution/internal/service/user"
)

type services struct {
	userService *userServicePkg.Service
}

func createService(repositories *repositories) *services {
	return &services{
		userService: userServicePkg.New(repositories.userRepository),
	}
}
