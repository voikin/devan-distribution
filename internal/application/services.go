package application

import (
	"github.com/voikin/devan-distribution/internal/config"
	userServicePkg "github.com/voikin/devan-distribution/internal/service/auth"
)

type services struct {
	userService *userServicePkg.Service
}

func createService(repositories *repositories, cfg *config.Config) *services {
	return &services{
		userService: userServicePkg.New(repositories.userRepository, &cfg.JWT),
	}
}
