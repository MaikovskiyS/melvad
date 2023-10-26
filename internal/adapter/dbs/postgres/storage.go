package psql

import (
	"context"
	"errors"
	"melvad/internal/model"

	"github.com/jackc/pgx/v5"
)

var (
	ErrInsert = errors.New("cant insert")
)

type store struct {
	conn *pgx.Conn
}

func New(c *pgx.Conn) *store {
	return &store{
		conn: c,
	}
}
func (s *store) Save(ctx context.Context, u model.User) (uint64, error) {
	sql := "INSERT INTO users(name,age) values ($1,$2) RETURNING ID"

	row := s.conn.QueryRow(ctx, sql, u.Name, u.Age)
	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return 0, ErrInsert
	}
	return id, nil
}
