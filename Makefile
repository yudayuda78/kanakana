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

# Menjalankan script database seeder
seed:
	go run cmd/seeder/main.go

# Install tool 'air' untuk auto-reload (cukup jalankan sekali)
install-air:
	go install github.com/air-verse/air@latest
