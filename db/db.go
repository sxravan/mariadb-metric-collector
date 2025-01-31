package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/sxravan/mariadb-metric-collector/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect(cfg *config.Config) (*sql.DB, error) {
	var dsn string

	// Check if Unix socket is provided
	if cfg.DBSocket != "" {
		dsn = fmt.Sprintf("%s:%s@unix(%s)/%s",
			cfg.DBUser, cfg.DBPassword, cfg.DBSocket, cfg.DBName)
	} else {
		// Use TCP connection if no socket is provided
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	}

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Ping to make sure the connection works
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}

func Close() {
	if db != nil {
		db.Close()
	}
}
