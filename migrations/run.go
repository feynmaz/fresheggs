// migrations/run.go
package migrations

import (
	"database/sql"
	"embed"
	
	migrate "github.com/rubenv/sql-migrate"
	_ "github.com/jackc/pgx/v5/stdlib"
)

//go:embed *
var migrationsFiles embed.FS

func Run(postgresConn string) error {
	db, err := sql.Open("pgx", postgresConn)
	if err != nil {
		return err
	}
	
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFiles, 
		Root:       ".",
	}
	
	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		return err
	}
	
	return nil
}
