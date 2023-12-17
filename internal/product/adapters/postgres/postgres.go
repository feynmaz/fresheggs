package postgres

import (
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	once sync.Once
	db   *sqlx.DB
	err  error
)

func GetDb(pgDsn string) (*sqlx.DB, error) {
	once.Do(func() {
		db, err = sqlx.Open("postgres", pgDsn)
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
