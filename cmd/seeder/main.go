package main

import (
	"fmt"
	"kanakana/internal/config"
	"kanakana/internal/models"
	"kanakana/internal/pkg/database"
	"kanakana/internal/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Memulai Seeder Admin...")

	cfg := config.Load()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	// Pastikan tabel users ada
	db.AutoMigrate(&models.User{})

	userRepo := repository.NewUserRepository(db)

	adminEmail := "admin@kanakana.com"
	adminPassword := "admin123"

	// Cek apakah admin sudah ada
	_, err = userRepo.FindByEmail(adminEmail)
	if err == nil {
		fmt.Println("Akun admin sudah ada di database.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Gagal hash password: %v", err)
	}

	adminUser := &models.User{
		Name:     "Super Admin",
		Email:    adminEmail,
		Password: string(hashedPassword),
		Role:     "admin",
	}

	if err := userRepo.Create(adminUser); err != nil {
		log.Fatalf("Gagal membuat akun admin: %v", err)
	}

	fmt.Println("==================================================")
	fmt.Println("Sukses membuat akun Admin!")
	fmt.Printf("Email    : %s\n", adminEmail)
	fmt.Printf("Password : %s\n", adminPassword)
	fmt.Println("==================================================")
}
