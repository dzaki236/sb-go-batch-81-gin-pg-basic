package config

import (
	"database/sql"
	"fmt"
	// "log"
	"os"
	// "rest-api/config"
)

var DB *sql.DB

func InitDatabase() error {
	keys := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSLMODE",
	}

	values := make([]string, len(keys))
	for i, key := range keys {
		values[i] = os.Getenv(key)
		if values[i] == "" {
			return fmt.Errorf("konfigurasi %s belum diisi", key)
		}
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		values[0], values[1], values[2], values[3], values[4], values[5],
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("gagal membuka database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("gagal terhubung ke PostgreSQL: %w", err)
	}

	DB = db
	return nil
}
