package postgres

import (
	"database/sql"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func GetDb(pgDsn string) (*sql.DB, error) {
	once.Do(func() {
		db, err = sql.Open("pgx", pgDsn)
		if err != nil {
			return
		}

		err = db.Ping()
		if err != nil {
			db.Close()
			db = nil
			return
		}
	})

	return db, err
}

func CloseDB() error {
	return db.Close()
}
