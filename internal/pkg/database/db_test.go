package database

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// loadTestEnv memuat .env dari root project untuk keperluan test
func loadTestEnv() {
	// Naik 4 level dari internal/pkg/database ke root project
	_ = godotenv.Load("../../../../.env")
}

func TestConnect(t *testing.T) {
	loadTestEnv()

	// Pastikan env variable tersedia
	requiredEnvs := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			t.Skipf("Env variable %s tidak ditemukan, test dilewati", env)
		}
	}

	err := Connect()
	if err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}
	defer Close()

	t.Log("Koneksi ke database berhasil")
}

func TestPing(t *testing.T) {
	loadTestEnv()

	requiredEnvs := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			t.Skipf("Env variable %s tidak ditemukan, test dilewati", env)
		}
	}

	if err := Connect(); err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}
	defer Close()

	if err := DB.Ping(); err != nil {
		t.Fatalf("Ping ke database gagal: %v", err)
	}

	t.Log("Ping ke database berhasil")
}

func TestClose(t *testing.T) {
	loadTestEnv()

	requiredEnvs := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			t.Skipf("Env variable %s tidak ditemukan, test dilewati", env)
		}
	}

	if err := Connect(); err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}

	if err := Close(); err != nil {
		t.Fatalf("Gagal menutup koneksi database: %v", err)
	}

	t.Log("Koneksi database berhasil ditutup")
}

func TestConnectionPool(t *testing.T) {
	loadTestEnv()

	requiredEnvs := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			t.Skipf("Env variable %s tidak ditemukan, test dilewati", env)
		}
	}

	if err := Connect(); err != nil {
		t.Fatalf("Gagal koneksi ke database: %v", err)
	}
	defer Close()

	stats := DB.Stats()
	t.Logf("Connection pool stats - MaxOpenConnections: %d, OpenConnections: %d, InUse: %d, Idle: %d",
		stats.MaxOpenConnections, stats.OpenConnections, stats.InUse, stats.Idle)

	if stats.MaxOpenConnections != 25 {
		t.Errorf("MaxOpenConns seharusnya 25, didapat %d", stats.MaxOpenConnections)
	}
}
