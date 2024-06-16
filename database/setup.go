package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	// Gunakan driver MySQL dan koneksi string yang sesuai dengan konfigurasi Anda
	connStr := "root:@tcp(127.0.0.1:3306)/mydb"
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")
}
