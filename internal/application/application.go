package application

import (
	"context"
	"net/http"
	"time"

	"github.com/voikin/devan-distribution/internal/config"
)

type App struct {
	httpServer *http.Server
	config     *config.Config
	handler    http.Handler
}

func New() (*App, error) {
	ctx := context.Background()

	cfg, err := config.InitConfig()
	if err != nil {
		return nil, err
	}

	app, err := initApp(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	a.httpServer = &http.Server{
		Addr:           ":" + a.config.HTTP.Port,
		Handler:        a.handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return nil
}

func (a *App) Stop() error {
	if a.httpServer == nil {
		return nil
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), a.config.HTTP.ShutdownServerTimeout)
	defer cancel()

	return a.httpServer.Shutdown(shutdownCtx)
}
