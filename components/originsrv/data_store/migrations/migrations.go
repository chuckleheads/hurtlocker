package migrations

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func Migrate(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err.Error())
	}
	// TED: really man? yeah this sucks but we need to hardcode
	// this for now until we get some actual config
	m, err := migrate.NewWithDatabaseInstance(
		"file:///src/components/originsrv/migrations",
		"postgres", driver)
	m.Up()
}
