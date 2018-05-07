package data_store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/chuckleheads/hurtlocker/components/originsrv/data_store/migrations"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Username string
	Password string
	SSLMode  string
	Host     string
	Port     int
	Database string
}

func New(cfg *DBConfig) *sql.DB {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s host=%s port=%d", cfg.Username, cfg.Database, cfg.SSLMode, cfg.Host, cfg.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	migrations.Migrate(db)
	return db
}
