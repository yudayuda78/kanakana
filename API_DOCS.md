# Kanakana API Documentation

Base URL: `http://localhost:8081/api/v1` (Port depends on `.env`)

## Authentication

### 1. Register User
- **URL**: `/auth/register`
- **Method**: `POST`
- **Description**: Daftarkan akun baru.
- **Request Body** (JSON):
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- **Response Success** (`201 Created`):
  ```json
  {
    "ID": 1,
    "CreatedAt": "2023-10-10...",
    "UpdatedAt": "2023-10-10...",
    "DeletedAt": null,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user"
  }
  ```

### 2. Login User
- **URL**: `/auth/login`
- **Method**: `POST`
- **Description**: Login untuk mendapatkan JWT token.
- **Request Body** (JSON):
  ```json
  {
    "email": "john@example.com",
    "password": "password123"
  }
  ```
- **Response Success** (`200 OK`):
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
  ```

---

## Kosakata (Public Endpoints)

Endpoint berikut dapat diakses tanpa autentikasi.

### 1. Dapatkan Semua Kosakata
- **URL**: `/kosakata`
- **Method**: `GET`
- **Response Success** (`200 OK`):
  ```json
  [
    {
      "ID": 1,
      "kanji": "食べる",
      "reading": "たべる",
      "romaji": "taberu",
      "arti": "makan",
      "levels": [
        {
          "ID": 1,
          "CreatedAt": "2023-10-10...",
          "UpdatedAt": "2023-10-10...",
          "DeletedAt": null,
          "name": "N5"
        }
      ]
    }
  ]
  ```

### 2. Dapatkan Kosakata Berdasarkan ID
- **URL**: `/kosakata/detail?id=1`
- **Method**: `GET`

### 3. Dapatkan Kosakata Berdasarkan Level JLPT
- **URL**: `/kosakata/level?level=N5`
- **Method**: `GET`

### 4. Dapatkan Kosakata Berdasarkan Kanji
- **URL**: `/kosakata/kanji?q=食べる`
- **Method**: `GET`

### 5. Dapatkan Kosakata Berdasarkan Reading (Hiragana/Katakana)
- **URL**: `/kosakata/reading?q=たべる`
- **Method**: `GET`

### 6. Dapatkan Kosakata Berdasarkan Romaji
- **URL**: `/kosakata/romaji?q=taberu`
- **Method**: `GET`

---

## Kosakata (Admin Protected Endpoints)

Endpoint berikut membutuhkan header **Authorization** dengan Bearer Token JWT. Token ini harus milik akun dengan `role: admin`.
Format header: `Authorization: Bearer <token_jwt>`

### 1. Tambah Kosakata
- **URL**: `/kosakata`
- **Method**: `POST`
- **Request Body** (JSON):
  ```json
  {
    "kanji": "飲む",
    "reading": "のむ",
    "romaji": "nomu",
    "arti": "minum",
    "levels": [
      {
        "ID": 1
      }
    ]
  }
  ```

### 2. Ubah Kosakata
- **URL**: `/kosakata?id=1`
- **Method**: `PUT`
- **Request Body** (JSON):
  Sama seperti Tambah Kosakata.

### 3. Hapus Kosakata
- **URL**: `/kosakata?id=1`
- **Method**: `DELETE`
- **Response Success** (`200 OK`):
  ```json
  {
    "message": "kosakata berhasil dihapus"
  }
  ```

---

## Health Check
- **URL**: `/health`
- **Method**: `GET`
- **Response Success** (`200 OK`):
  ```json
  {
    "status": "ok",
    "message": "Backend Kanakana siap!"
  }
  ```
