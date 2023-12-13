package postgres

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func GetDb(pgDsn string) (*sql.DB, error) {
	once.Do(func() {
		db, err = sql.Open("postgres", pgDsn)
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
