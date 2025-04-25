# Backend - Go Fiber API

REST API ini dibangun menggunakan [Go](https://golang.org/) dengan framework [Fiber](https://gofiber.io/).

## Fitur Utama
- Autentikasi JWT (Login, Register, Me)
- Manajemen User (CRUD)
- Validasi input
- Dokumentasi API (Swagger/OpenAPI)
- Struktur modular: handler, repository, dto, middleware, routes

## Struktur Direktori
```
backend/
├── config/         # Konfigurasi database, env, migrasi
├── docs/           # Dokumentasi OpenAPI/Swagger
├── dto/            # Data Transfer Object (request/response)
├── handlers/       # HTTP handler (controller)
├── middleware/     # Middleware (JWT, dll)
├── models/         # Model database (GORM)
├── repository/     # Repository pattern (akses DB)
├── routes/         # Definisi endpoint API
├── main.go         # Entry point aplikasi
├── go.mod, go.sum  # Dependency Go
```

## Instalasi & Menjalankan
1. Masuk ke folder backend:
   ```bash
   cd backend
   ```
2. Install dependency:
   ```bash
   go mod tidy
   ```
3. Copy `.env` atau sesuaikan konfigurasi database.
4. Jalankan aplikasi:
   ```bash
   go run main.go
   ```
   Default: http://localhost:8000

## Endpoint Penting
- `/auth/login`     : Login user
- `/auth/register`  : Registrasi user
- `/auth/me`        : Info user login (JWT protected)
- `/users`          : CRUD user (JWT protected)
- `/swagger/`       : Dokumentasi Swagger UI
- `/openapi/swagger.yaml` : File OpenAPI

## Swagger & Dokumentasi
Swagger UI tersedia di: `http://localhost:8000/swagger/`

## Testing
Gunakan Postman/Insomnia atau Swagger UI untuk mencoba endpoint.

## Deployment
- Build binary: `go build -o app`
- (Opsional) Siapkan Dockerfile untuk containerization

---

Aplikasi backend ini adalah bagian dari monorepo Go + React. Untuk frontend, lihat folder `../frontend`.
