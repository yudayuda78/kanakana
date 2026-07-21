package database

import (
	"fmt"
	"kanakana/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect membuat koneksi ke database PostgreSQL menggunakan GORM
// dan mengembalikan instance *gorm.DB yang siap digunakan.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi database: %w", err)
	}

	// Ambil underlying *sql.DB untuk konfigurasi connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan sql.DB dari gorm: %w", err)
	}

	// Verifikasi koneksi aktif
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping ke database: %w", err)
	}

	// Konfigurasi connection pool
	sqlDB.SetMaxOpenConns(cfg.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.DBMaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLifetime) * time.Minute)

	return db, nil
}
