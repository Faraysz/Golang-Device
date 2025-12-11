# Belajar Golang: Level 1 hingga Level 2  
Proyek ini merupakan rangkaian latihan belajar bahasa pemrograman **Golang (Go)** mulai dari level dasar (CLI Program) hingga membuat **REST API sederhana**.  
Proyek ini dirancang untuk memberikan pemahaman bertahap, mudah diikuti, dan relevan untuk pengembangan backend modern.

---

# 🎯 Tujuan Pembelajaran

1. Memahami sintaks dasar Golang  
2. Memahami cara membuat program CLI sederhana  
3. Memahami struct, slice, fungsi, dan modular code  
4. Memahami dasar-dasar backend menggunakan `net/http`  
5. Membuat REST API sederhana dengan arsitektur modular  

---

# 🚀 LEVEL 1 — Program CLI: Device Manager  
Level pertama bertujuan untuk memahami:

- Dasar Go: package, function, import
- `struct` sebagai model data
- `slice` sebagai penyimpanan data
- Input user dengan `bufio`
- Logika program (`switch`, `for`)
- Modularisasi sederhana

### 🎮 **Fitur Program CLI**
- Tambah device
- Lihat seluruh device
- Cari device berdasarkan UID
- Navigasi menu menggunakan input terminal

### 📌 **File: `main.go` (Level 1)**

Program CLI menampilkan menu seperti:

=== MENU DEVICE MANAGER ===
Tambah Device
Lihat Semua Device
Cari Device by UID
Keluar


Level 1 ini memberikan dasar kuat sebelum masuk ke HTTP server.

---

# 🚀 LEVEL 2 — Basic REST API dengan Golang (Tanpa Framework)

Level ini fokus pada dasar backend:

- Membuat server HTTP
- Routing dengan `http.HandleFunc`
- Handler (controller)
- JSON request & response
- Model & repository pattern
- Modular folder layout seperti backend beneran


### 📌 **Fitur API**
| Method | Endpoint | Fungsi |
|--------|----------|--------|
| GET | `/` | Test API berjalan |
| GET | `/devices` | Menampilkan seluruh device |
| POST | `/device/create` | Menambah device dengan JSON |

---

# ▶️ Cara Menjalankan Project Level 2

Masuk ke folder:

-cd golang-api
-Inisialisasi module Go:
  1. go mod init golang-api
  2. go mod tidy
  3. Jalankan server:
     go run main.go
     Server berjalan di: http://localhost:8080
