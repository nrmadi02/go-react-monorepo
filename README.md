# Go React Monorepo

Monorepo ini berisi dua project utama:
- **frontend**: Aplikasi web berbasis React + Vite + TailwindCSS
- **backend**: REST API menggunakan Go (Golang) dengan Fiber framework

## Struktur Direktori

```
├── backend/        # Source code backend (Go + Fiber)
├── frontend/       # Source code frontend (React + Vite)
├── node_modules/   # Dependency root (monorepo)
├── package.json    # Dependency management (pnpm workspaces)
├── pnpm-lock.yaml  # Lockfile untuk pnpm
```

## Prasyarat
- Node.js & pnpm (untuk frontend/monorepo)
- Go (untuk backend)
- Docker (opsional, untuk deployment)

## Instalasi

### 1. Clone Repository
```bash
git clone https://github.com/username/go_react_monorepo.git
cd go_react_monorepo
```

### 2. Instalasi Dependency
#### Monorepo (semua workspace)
```bash
pnpm install
```

#### Backend
```bash
cd backend
go mod tidy
```

#### Frontend
```bash
cd frontend
pnpm install
```

## Menjalankan Project

### Backend (Go + Fiber)
```bash
cd backend
go run main.go
```
Akses API di: `http://localhost:8000`

### Frontend (React + Vite)
```bash
cd frontend
pnpm run dev
```
Akses web di: `http://localhost:5173`

## Build & Deployment

- **Frontend**: Ikuti instruksi di `frontend/README.md` (sudah support Docker, Vite, Tailwind)
- **Backend**: Build Go binary, atau gunakan Dockerfile (tambahkan jika diperlukan)

## Dokumentasi API
- Swagger/OpenAPI tersedia di endpoint `/swagger/` pada backend
- File OpenAPI YAML: `backend/docs/openapi3.yaml`

## Kontribusi
1. Fork repo
2. Buat branch fitur: `git checkout -b fitur-anda`
3. Commit perubahan
4. Push dan buat Pull Request

---

Monorepo ini dibuat untuk memudahkan pengembangan aplikasi modern dengan stack Go dan React dalam satu repository.
