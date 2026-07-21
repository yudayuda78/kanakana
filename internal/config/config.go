package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config menyimpan semua konfigurasi aplikasi yang dibaca dari environment variable.
type Config struct {
	Port string

	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBSSLMode         string
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime int // dalam menit

	JWTSecret          string
	JWTExpirationHours int
}

// Load membaca environment variable dan mengembalikan Config.
// Setiap field memiliki nilai default jika env variable tidak diset.
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: file .env tidak ditemukan, menggunakan environment default")
	}

	return &Config{
		Port:               getEnv("PORT", "8081"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", "kanakana"),
		DBSSLMode:          getEnv("DB_SSLMODE", "disable"),
		DBMaxOpenConns:     getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
		DBMaxIdleConns:     getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
		DBConnMaxLifetime:  getEnvAsInt("DB_CONN_MAX_LIFETIME", 15),
		JWTSecret:          getEnv("JWT_SECRET", "supersecretkey"),
		JWTExpirationHours: getEnvAsInt("JWT_EXPIRATION_HOURS", 24),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil && parsed > 0 {
			return parsed
		}
	}
	return defaultVal
}
