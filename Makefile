.PHONY: run dev build test seed install-air

# Menjalankan aplikasi secara biasa
run:
	go run cmd/kanakana/main.go

# Menjalankan dengan fitur auto-reload (pastikan sudah install 'air')
dev:
	air

# Membangun file binary aplikasi
build:
	go build -o bin/kanakana cmd/kanakana/main.go

# Menjalankan unit test
test:
	go test ./... -v

# Menjalankan script database seeder (admin)
seed:
	go run cmd/seeder/main.go

# Menjalankan seeder kosakata N4
seed-kosakata:
	go run cmd/seeder/kosakata/main.go

# Install tool 'air' untuk auto-reload (cukup jalankan sekali)
install-air:
	go install github.com/air-verse/air@latest
