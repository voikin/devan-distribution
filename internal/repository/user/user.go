package user

import "github.com/jackc/pgx/v5"

type Repo struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Repo {
	return &Repo{
		conn: conn,
	}
}
