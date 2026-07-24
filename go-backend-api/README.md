# RESTful API Task Management (To-Do List) - Golang

API To-Do List yang dikembangkan dengan Golang, Gin Framework, PostgreSQL, Redis, dan JWT Authentication.

## 🚀 Fitur Utama
- **CRUD Operations**: Mengelola tugas (Task).
- **Filtering & Search**: Pencarian tugas berdasarkan `status` dan kata kunci `search` pada judul/deskripsi.
- **Pagination**: Efisiensi pemanggilan data dengan limit & offset.
- **Concurrency Execution**: Menggunakan Goroutines (`sync.WaitGroup`) saat mengambil list tugas dan total record secara paralel.
- **JWT Authentication**: Pengamanan endpoint API menggunakan Bearer Token.
- **Redis Caching**: Caching otomatis untuk endpoint `GET /tasks` guna mempercepat response.
- **Logging**: Encapsulated logger untuk mencatat error runtime.
- **Validation**: Schema-level request validation menggunakan Gin/Validator.

---

## 🛠️ Prasyarat (Prerequisites)
- [Go](https://go.dev/) (v1.20 atau versi lebih baru)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/) (Opsional, jika tidak ada API tetap berjalan normal)

---

## ⚙️ Cara Menjalankan Aplikasi

### 1. Clone & Install Dependencies
```bash
git clone <repository_url>
cd todo-api
go mod tidy
```

### 2. Pengaturan Database PostgreSQL
Buat sebuah database baru di PostgreSQL Anda:
```sql
CREATE DATABASE todo_db;
```

### 3. Setup Environment Variable
Atur kredensial database Anda (default dapat langsung digunakan jika sesuai):
- **Linux/macOS:**
  ```bash
  export DB_HOST=localhost
  export DB_PORT=5432
  export DB_USER=postgres
  export DB_PASSWORD=postgres
  export DB_NAME=todo_db
  export REDIS_HOST=localhost
  export REDIS_PORT=6379
  ```
- **Windows (PowerShell):**
  ```powershell
  $env:DB_HOST="localhost"
  $env:DB_USER="postgres"
  $env:DB_PASSWORD="postgres"
  $env:DB_NAME="todo_db"
  ```

### 4. Menjalankan Server
```bash
go run main.go
```
Server akan berjalan di `http://localhost:8080`.

---

## 🧪 Menjalankan Unit Test
Untuk menjalankan tes unit pada fungsi validasi API:
```bash
go test -v ./...
```

---

## 🔐 Autentikasi (JWT)

Gunakan endpoint berikut untuk mendapatkan **Bearer Token** sebelum mengakses endpoint Task:

- **Endpoint**: `POST /login`
- **Request Body**:
  ```json
  {
    "username": "admin",
    "password": "password"
  }
  ```
- **Response**:
  ```json
  {
    "token": "<YOUR_JWT_TOKEN>"
  }
  ```

> *Sertakan token di atas pada Header request sebagai:*
> `Authorization: Bearer <YOUR_JWT_TOKEN>`

---

## 📌 Dokumentasi Endpoint API

### 1. Create Task
- **POST** `/tasks`
- **Body**:
  ```json
  {
     "title": "Belajar Golang Concurrency",
     "description": "Mempelajari Goroutine dan Channel",
     "status": "pending",
     "due_date": "2026-08-01"
  }
  ```

### 2. Get All Tasks
- **GET** `/tasks?status=pending&page=1&limit=5&search=Golang`
- **Response**:
  ```json
  {
     "tasks": [
        {
           "id": "e7b0a88b-21d1-4b10-a292-6d2719277f0a",
           "title": "Belajar Golang Concurrency",
           "description": "Mempelajari Goroutine dan Channel",
           "status": "pending",
           "due_date": "2026-08-01"
        }
     ],
     "pagination": {
        "current_page": 1,
        "total_pages": 1,
        "total_tasks": 1
     }
  }
  ```

### 3. Get Task by ID
- **GET** `/tasks/:id`

### 4. Update Task
- **PUT** `/tasks/:id`
- **Body**:
  ```json
  {
     "title": "Belajar Golang Concurrency (Selesai)",
     "description": "Mempelajari Goroutine dan Channel",
     "status": "completed",
     "due_date": "2026-08-01"
  }
  ```

### 5. Delete Task
- **DELETE** `/tasks/:id`
