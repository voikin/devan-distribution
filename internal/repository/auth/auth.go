package auth

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/voikin/devan-distribution/internal/entity"
	"github.com/voikin/devan-distribution/internal/errs"
)

type Repo struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (r *Repo) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	var id int64
	row := r.conn.QueryRow(ctx, createUserQuery, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	var user entity.User
	// TODO: Change user properties if we need
	err := r.conn.QueryRow(ctx, getUserQuery, username, password).Scan(&user.Username, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewErrorNotFound(username)
		}
		return nil, err
	}

	return &user, nil
}
