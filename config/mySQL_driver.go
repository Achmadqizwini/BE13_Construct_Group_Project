package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectToDB() *sql.DB {
	// format --> <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	var connectionString = os.Getenv("DB_Connection")
	db, err := sql.Open("mysql", connectionString) // membuka koneksi ke database
	if err != nil {                                // pengecekan error yang terjadi ketika proses open connection
		log.Fatal("error open connection", err.Error())
	}
	errPing := db.Ping() // mengecek apakah aplikasi masih terkoneksi ke database
	if errPing != nil {  // handling error ketika gagal connect ke db
		log.Fatal("error connect to db", errPing.Error())
	} else {
		fmt.Println("koneksi berhasil")
	}

	return db
}
