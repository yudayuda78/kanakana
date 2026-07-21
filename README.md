# Kanakana Backend

Kanakana Backend adalah RESTful API yang dibangun menggunakan [Go (Golang)](https://golang.org/) untuk mengelola data kosakata bahasa Jepang. Proyek ini menggunakan arsitektur standar Go dan terhubung ke database PostgreSQL.

## Fitur Utama

- Endpoint CRUD untuk Kosakata bahasa Jepang (Kanji, Reading, Romaji, Arti, Level JLPT)
- Koneksi ke database PostgreSQL
- Konfigurasi berbasis environment variable (`.env`)

## Prasyarat

Pastikan Anda telah menginstal:
- [Go](https://golang.org/dl/) (versi 1.26.2 atau terbaru)
- [PostgreSQL](https://www.postgresql.org/download/)

## Cara Menjalankan Aplikasi Lokal

1. **Clone repository ini**
   ```bash
   git clone https://github.com/yudayuda78/kanakana.git
   cd kanakana-backend
   ```

2. **Siapkan konfigurasi `.env`**
   Buat file `.env` di root direktori proyek dan sesuaikan nilai konfigurasinya dengan database lokal Anda:
   ```env
   PORT=8081
   
   # Konfigurasi Database PostgreSQL
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=password
   DB_NAME=kanakana
   DB_SSLMODE=disable
   ```

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

4. **Jalankan Aplikasi**
   Untuk development:
   ```bash
   # Masuk ke direktori utama (tempat app dijalankan)
   # (Tergantung struktur spesifik Anda, misal cmd/api)
   # Contoh jika ada main.go di root atau cmd/:
   go run main.go 
   ```
   *Catatan: Pastikan database PostgreSQL sudah menyala dan database `kanakana` sudah dibuat.*

## Menjalankan Unit Test

Proyek ini dilengkapi dengan unit test. Test untuk koneksi database akan otomatis dilewati (`skipped`) jika environment variable tidak disetup.

```bash
go test ./... -v
```

## Struktur Proyek

```text
.
├── .env                # Konfigurasi environment (tidak di-commit)
├── internal/           # Kode privat aplikasi
│   ├── app/            # Setup inisialisasi aplikasi (server, route)
│   ├── handlers/       # HTTP handlers/controllers untuk endpoint
│   ├── models/         # Definisi struct/tipe data
│   └── pkg/            # Package internal (seperti koneksi database)
├── go.mod              # Definisi module & dependencies
└── README.md           # Dokumentasi proyek
```

## Endpoints

Dokumentasi API lengkap beserta panduan _request_ dan _response_-nya bisa dilihat di file [API_DOCS.md](./API_DOCS.md).

## Dependencies Utama

- `github.com/lib/pq` - Driver PostgreSQL
- `github.com/joho/godotenv` - Memuat file `.env`

---
*Dibuat untuk project Kanakana.*
