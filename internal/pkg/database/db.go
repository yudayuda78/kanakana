package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DB adalah instance koneksi database yang digunakan secara global
var DB *sql.DB

// Connect membuat koneksi ke database PostgreSQL berdasarkan env variable
func Connect() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("gagal membuka koneksi database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("gagal terhubung ke database: %w", err)
	}

	// Konfigurasi connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	DB = db
	return nil
}

// Close menutup koneksi database
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
