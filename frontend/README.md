# Frontend - React + Vite

Aplikasi web frontend ini dibuat dengan [React](https://react.dev/), [Vite](https://vitejs.dev/), dan [TailwindCSS](https://tailwindcss.com/). Bagian ini merupakan bagian dari monorepo Go + React.

## Fitur Utama
- React (TypeScript)
- Vite (fast build & HMR)
- TailwindCSS (utility-first styling)
- Routing dengan React Router
- Siap untuk deployment (Docker, static build)

## Struktur Direktori
```
frontend/
├── app/          # Komponen utama aplikasi
├── common/       # Komponen/utilitas bersama
├── public/       # Static assets
├── screens/      # Halaman utama
├── .env          # Environment variable
├── package.json  # Dependency frontend
├── vite.config.ts# Konfigurasi Vite
```

## Instalasi & Menjalankan
1. Masuk ke folder frontend:
   ```bash
   cd frontend
   ```
2. Install dependency:
   ```bash
   pnpm install
   ```
3. Copy `.env.example` ke `.env` dan sesuaikan jika perlu.
4. Jalankan development server:
   ```bash
   pnpm run dev
   ```
   Default: http://localhost:5173

## Build & Deployment
- Build untuk production:
  ```bash
  pnpm run build
  ```
- Jalankan preview production:
  ```bash
  pnpm run preview
  ```
- Untuk deployment Docker:
  ```bash
  docker build -t frontend-app .
  docker run -p 3000:3000 frontend-app
  ```

## Integrasi dengan Backend
- Endpoint API dapat disesuaikan di file `.env`
- Pastikan backend berjalan sebelum menggunakan fitur login/dll

## Styling
Menggunakan TailwindCSS, bisa dikombinasikan dengan framework lain sesuai kebutuhan.

---

Frontend ini adalah bagian dari monorepo Go + React. Untuk backend, lihat folder `../backend`.

