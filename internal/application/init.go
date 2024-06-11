package application

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/voikin/devan-distribution/internal/config"
	"github.com/voikin/devan-distribution/internal/controller"
)

func initApp(ctx context.Context, cfg *config.Config) (*App, error) {
	conn, err := createPostgres(ctx, cfg)
	if err != nil {
		return nil, err
	}

	repository := createRepository(conn)
	services := createService(repository)
	usecase := createUseCase(services)
	controller := controller.New(usecase)

	ginEngine := gin.Default()

	initRoutes(ginEngine, controller)

	return &App{
		config:  cfg,
		handler: ginEngine,
	}, nil
}
