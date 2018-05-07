package data_store

import (
	"database/sql"
	"fmt"
	"log"

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
	return db
}

// func (srv *Server) RunMigrations(db *sql.DB) {
// 	// This format is not ideal =/
// 	log.Println("Running db migrations")
// 	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS origins (id bigserial PRIMARY KEY, name TEXT NOT NULL UNIQUE, default_package_visibility text NOT NULL DEFAULT 'public')"); err != nil {
// 		log.Fatal(err)
// 	}
// }
