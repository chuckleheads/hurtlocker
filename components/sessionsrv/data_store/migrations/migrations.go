package migrations

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func Migrate(db *sql.DB, migrations_dir string) {
	driver, err := cockroachdb.WithInstance(db, &cockroachdb.Config{})
	if err != nil {
		panic(err.Error())
	}
	log.Printf("file://%s", migrations_dir)
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrations_dir),
		"postgres", driver)
	if err != nil {
		panic(err.Error())
	}
	// TED: better handling of errors here when you know more about Go Error handling
	m.Up()
}
