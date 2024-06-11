package application

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/voikin/devan-distribution/internal/config"
)

func createPostgres(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, cfg.Pg.GetDSN())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
