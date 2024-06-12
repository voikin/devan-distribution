package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	var roleId string
	err := r.conn.QueryRow(ctx, getUserQuery, username).Scan(&user.ID, &user.Username, &user.Password, &roleId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewErrorNotFound(username)
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repo) GetUserById(ctx context.Context, userId int64) (*entity.User, error) {
	var user entity.User
	err := r.conn.QueryRow(ctx, getUserQuery, userId).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// TODO: Fix This
			return nil, errs.NewErrorNotFound(fmt.Sprintf("user with id: %d not found", userId))
		}
		return nil, err
	}

	return &user, nil
}
