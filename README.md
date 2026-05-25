# Belajar Golang Database (MySQL)

Repositori ini berisi kumpulan materi, implementasi kode, dan unit pengujian (*unit testing*) untuk memahami interaksi database menggunakan bahasa pemrograman Go (Golang) dan driver MySQL secara mendalam. Proyek ini mencakup dasar pemrosesan SQL hingga penerapan arsitektur **Repository Pattern**.

---

## 🚀 Fitur & Materi yang Dipelajari

* **Database Connection Pool**: Manajemen *lifecycle* koneksi database menggunakan `database/sql` agar efisien dan mencegah kebocoran koneksi (*connection leak*).
* **CRUD Operations**: Eksekusi perintah SQL manipulasi data menggunakan `ExecContext` dan membaca data dengan `QueryContext`.
* **SQL Injection Protection**: Mengamankan query database dari serangan injeksi kode menggunakan *argument binding / parameterized query* (`?`).
* **Null Value Handling**: Menangani data kolom yang bernilai `NULL` di database agar tidak menyebabkan *crash/error* di Go menggunakan `sql.NullString` dan `sql.NullTime`.
* **Prepared Statement**: Optimasi performa eksekusi query yang dilakukan berulang kali secara massal melalui metode *pre-compiled* query.
* **Database Transaction**: Implementasi mekanisme ACID transaksi (`Begin`, `Commit`, `Rollback`) yang aman dari *panic* menggunakan kombinasi *control flow* `defer tx.Rollback()`.
* **Repository Pattern (Clean Architecture)**: Memisahkan logika manipulasi database (layer infrastruktur) dengan layer bisnis menggunakan abstraksi berbasis *Interface*.

---

## 🏗️ Arsitektur Proyek (Repository Pattern)

Aplikasi ini menerapkan pemisahan komponen (*decoupling*) dengan struktur file sebagai berikut:

```text
├── connection/
│   └── connection.go       # Konfigurasi database pool
├── entity/
│   ├── comment.go          # Struct cetakan data comment
│   └── customer.go         # Struct cetakan data customer
├── repository/
│   ├── comment_repository.go      # Interface (Kontrak fungsi comment)
│   ├── comment_repository_impl.go # Implementasi query SQL comment
│   ├── customer_repository.go     # Interface (Kontrak fungsi customer)
│   └── customer_repository_impl.go# Implementasi query SQL customer
├── 02-sql-test/
│   └── sql_test.go         # Kumpulan Unit Test inti database
├── go.mod                  # Manajemen modul aplikasi
└── go.sum                  # Verifikasi checksum dependency
```

## Alur Kerja Aliran Data
1. Entity: Representasi objek struktur tabel database dalam bentuk tipe data Go (`struct`).
2. Interface: Kontrak blueprint atau daftar fungsi yang mendefinisikan apa saja operasi yang bisa dilakukan ke database.
3. Implementation: Kelas konkrit yang membawa koneksi objek `*sql.DB` dan mengeksekusi query SQL mentah yang sebenarnya.

## 🧪 Dokumentasi Unit Test (`sql_test.go`)
Seluruh fungsionalitas diuji menggunakan ekosistem bawaan pengujian Go (testing). Berikut adalah beberapa pengujian utama yang diimplementasikan:

1. Eksekusi SQL & SQL Parameter
- `TestExecSql`: Menguji teknik insert data manual langsung maupun menggunakan parameter ? untuk keamanan.
- `TestSqlInjection`: Membuktikan kerentanan manipulasi string query dan mendemonstrasikan solusinya menggunakan parameterized query.

2. Penanganan Nilai NULL secara Aman
- `TestQuerySqlComplex`: Menunjukkan cara melakukan teknik `.Scan()` data kompleks dari database ke variabel internal Go. Menggunakan `sql.NullString` pada kolom email dan `sql.NullTime` pada kolom `birth_date` untuk validasi kondisi data kosong tanpa memicu kegagalan sistem.

3. Manajemen Transaksi Aman (Dapat Diandalkan)
- `TestTransactionAman`: Mengimplementasikan standard best practice industri dalam mengelola state transaksi:

```go
tx, _ := db.Begin()
defer tx.Rollback() // Jaring pengaman otomatis saat terjadi panic/error

// ... proses transaksi internal ...

tx.Commit() // Menyimpan permanen jika seluruh proses tanpa kendala
```
## 🛠️ Cara Menjalankan Project
1. Prasyarat Sistem
- Go (Versi terbaru)
- MySQL Server / MariaDB aktif

2. Konfigurasi Skema Database
- Pastikan Anda telah membuat tabel berikut di dalam basis data MySQL Anda sebelum menjalankan pengujian:

```SQL
CREATE TABLE customer (
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE customer2 (
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    balance INT DEFAULT 0,
    rating DOUBLE DEFAULT 0.0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    birth_date DATE,
    married BOOLEAN DEFAULT false,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE comment (
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comment TEXT NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;
```
3. Menjalankan Seluruh Pengujian
Gunakan perintah berikut di terminal root proyek untuk mengeksekusi seluruh skenario pengujian database:

```Bash
go test -v ./...
```
