package connection

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// database pooling
func GetConnection() *sql.DB {
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_db")
	// parseTime=true untuk mengatasi error ketika scan time ke struct
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // Jumlah koneksi idle yang diperbolehkan
	db.SetMaxOpenConns(100)                 // Jumlah koneksi maksimum yang diperbolehkan
	db.SetConnMaxIdleTime(5 * time.Minute)  // Waktu maksimum koneksi idle
	db.SetConnMaxLifetime(60 * time.Minute) // Waktu maksimum koneksi hidup

	return db
}
