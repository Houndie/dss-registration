package postgres

import "github.com/jackc/pgx"

type Store struct {
	conn *pgx.Conn
}

func NewStore(conn *pgx.Conn) *Store {
	return &Store{
		conn: conn,
	}
}
