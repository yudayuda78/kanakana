package database

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"kanakana/internal/config"
)

// loadTestEnv memuat .env dari root project untuk keperluan test.
func loadTestEnv() {
	// Naik 4 level dari internal/pkg/database ke root project
	_ = godotenv.Load("../../../../.env")
}

// skipIfNoDB memeriksa apakah env variable DB tersedia.
// Jika tidak, test akan di-skip agar aman dijalankan tanpa database.
func skipIfNoDB(t *testing.T) {
	t.Helper()
	required := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, env := range required {
		if os.Getenv(env) == "" {
			t.Skipf("Env variable %s tidak ditemukan, test dilewati", env)
		}
	}
}

func TestConnect(t *testing.T) {
	loadTestEnv()
	skipIfNoDB(t)

	cfg := config.Load()
	db, err := Connect(cfg)
	if err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Gagal mendapatkan sql.DB: %v", err)
	}
	defer sqlDB.Close()

	t.Log("Koneksi ke database berhasil")
}

func TestPing(t *testing.T) {
	loadTestEnv()
	skipIfNoDB(t)

	cfg := config.Load()
	db, err := Connect(cfg)
	if err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Gagal mendapatkan sql.DB: %v", err)
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Ping ke database gagal: %v", err)
	}

	t.Log("Ping ke database berhasil")
}

func TestClose(t *testing.T) {
	loadTestEnv()
	skipIfNoDB(t)

	cfg := config.Load()
	db, err := Connect(cfg)
	if err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Gagal mendapatkan sql.DB: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		t.Fatalf("Gagal menutup koneksi database: %v", err)
	}

	t.Log("Koneksi database berhasil ditutup")
}

func TestConnectionPool(t *testing.T) {
	loadTestEnv()
	skipIfNoDB(t)

	cfg := config.Load()
	db, err := Connect(cfg)
	if err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Gagal mendapatkan sql.DB: %v", err)
	}
	defer sqlDB.Close()

	stats := sqlDB.Stats()
	t.Logf("Connection pool stats - MaxOpenConnections: %d, OpenConnections: %d, InUse: %d, Idle: %d",
		stats.MaxOpenConnections, stats.OpenConnections, stats.InUse, stats.Idle)

	if stats.MaxOpenConnections != cfg.DBMaxOpenConns {
		t.Errorf("MaxOpenConns seharusnya %d, didapat %d", cfg.DBMaxOpenConns, stats.MaxOpenConnections)
	}
}
