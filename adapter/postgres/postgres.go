package postgres

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	once sync.Once
	conn *pgx.Conn
	err  error
)

func GetDb(ctx context.Context, pgDsn string) (*pgx.Conn, error) {
	once.Do(func() {
		conn, err = pgx.Connect(ctx, pgDsn)
		if err != nil {
			return
		}

		err = conn.Ping(ctx)
		if err != nil {
			conn.Close(ctx)
			conn = nil
			return
		}
	})

	return conn, err
}

func CloseDB(ctx context.Context) error {
	return conn.Close(ctx)
}
