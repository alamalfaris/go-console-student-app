package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  //jumlah minimal koneksi yang dibuat saat aplikasi pertama jalan
	db.SetMaxOpenConns(100)                 //jumlah maksimal koneksi yang dibuat
	db.SetConnMaxIdleTime(5 * time.Minute)  //waktu maksimal open koneksi saat idle
	db.SetConnMaxLifetime(60 * time.Minute) //waktu maksimal open koneksi, kalo udah lewat waktu, nanti di re-open connection
	return db
}
