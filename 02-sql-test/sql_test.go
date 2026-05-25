package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	. "github.com/abdul452/belajar-golang-database/connection"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// exec sql tanpa parameter
	script := `INSERT INTO customer(id, name) VALUES('coba123', 'coba nama 1')`
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	// exec sql dengan parameter
	script2 := `INSERT INTO customer(id, name) VALUES(?, ?)`
	_, err = db.ExecContext(ctx, script2, "coba321", "coba nama 2")
	if err != nil {
		panic(err)
	}
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := `SELECT id, name FROM customer`
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() { // untuk cek apakah masih ada data yang bisa di scan
		var id, name string
		err := rows.Scan(&id, &name) // scan hasil query ke variabel / untuk membaca hasil query
		if err != nil {
			panic(err)
		}

		fmt.Println("id:", id)
		fmt.Println("name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := `SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer2`
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() { // untuk cek apakah masih ada data yang bisa di scan
		var id, name string
		var email sql.NullString // untuk handle null value di database, karena jika scan null value ke string akan error, maka gunakan sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime // untuk handle null value di database, karena jika scan null value ke time.Time akan error, maka gunakan sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDate, &married) // scan hasil query ke variabel / untuk membaca hasil query
		if err != nil {
			panic(err)
		}

		fmt.Println("=================")
		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid {
			fmt.Println("email:", email.String)
		} else {
			fmt.Println("email: (null)")
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		fmt.Println("createdAt:", createdAt)
		if birthDate.Valid {
			fmt.Println("birthDate:", birthDate.Time)
		} else {
			fmt.Println("birthDate: (null)")
		}
		fmt.Println("married:", married)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	// username := "admin'; #"
	password := "admin"

	// query dengan parameter, aman dari sql injection
	script := `SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1`
	rows, err := db.QueryContext(ctx, script, username, password)

	// query tanpa parameter, rentan terhadap sql injection
	// script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	// fmt.Println(script) // cek query yang dihasilkan
	// rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("Login berhasil, username:", username)
	} else {
		fmt.Println("Login gagal")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := `INSERT INTO user(username, password) VALUES(?, ?)`
	_, err := db.ExecContext(ctx, script, "abdul", "abdul nama")
	if err != nil {
		panic(err)
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "abl@example.com"
	comment := "ini kentar"

	script := `INSERT INTO comment(email, comment) VALUES(?, ?)`
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId() // untuk mendapatkan id terakhir yang di insert
	if err != nil {
		panic(err)
	}

	fmt.Println("Komentar berhasil ditambahkan dengan id:", id)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := `INSERT INTO comment(email, comment) VALUES(?, ?)`
	statement, err := db.PrepareContext(ctx, script) // untuk membuat prepared statement
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ { // untuk insert data dengan prepared statement, kita bisa menggunakan statement.ExecContext() berkali-kali dengan parameter yang berbeda
		email := fmt.Sprintf("user%d@example.com", i)
		comment := fmt.Sprintf("Komentar %d", i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId() // untuk mendapatkan id terakhir yang di insert
		if err != nil {
			panic(err)
		}

		fmt.Println("Komentar berhasil ditambahkan dengan id:", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin() // untuk memulai transaksi
	if err != nil {
		panic(err)
	}

	script1 := `INSERT INTO comment(email, comment) VALUES(?, ?)`
	// do transaction
	for i := 0; i < 10; i++ { // untuk insert data dengan prepared statement, kita bisa menggunakan statement.ExecContext() berkali-kali dengan parameter yang berbeda
		email := fmt.Sprintf("user%d@example.com", i)
		comment := fmt.Sprintf("Komentar %d", i)
		result, err := tx.ExecContext(ctx, script1, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId() // untuk mendapatkan id terakhir yang di insert
		if err != nil {
			panic(err)
		}

		fmt.Println("Komentar berhasil ditambahkan dengan id:", id)
	}

	// err = tx.Commit() // untuk commit transaksi
	err = tx.Rollback() // untuk rollback transaksi
	if err != nil {
		panic(err)
	}
}

// dari gemini
func TestTransactionAman(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// 1. PASANG DEFER ROLLBACK DI SINI SEBAGAI JARING PENGAMAN
	// Jika kode di bawah ini aman, defer ini diabaikan setelah commit.
	// Jika kode di bawah ini PANIC atau ERROR, defer ini otomatis menyelamatkan database.
	defer tx.Rollback()

	script1 := `INSERT INTO comment(email, comment) VALUES(?, ?)`

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("user%d@example.com", i)
		comment := fmt.Sprintf("Komentar %d", i)

		_, err := tx.ExecContext(ctx, script1, email, comment)
		if err != nil {
			panic(err) // ➡️ Jika panic di sini, defer tx.Rollback() langsung bekerja!
		}
	}

	// 2. JIKA SEMUA PROSES DI ATAS SUKSES, KITA COMMIT
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	// Di titik ini, transaksi SELESAI. Defer rollback di atas saat dipanggil
	// di akhir fungsi tidak akan ngefek apa-apa lagi ke database.
}
