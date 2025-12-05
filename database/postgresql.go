package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB{

	dsn := os.Getenv("POSGRES_URI")

	// Koneksi database
	db, err := sql.Open("postgres", dsn)

	if err != nil{
		log.Fatal("Gagal koneksi ke database",err)
	}

	// Tes Koneksi
	if err = db.Ping(); err != nil{
		log.Fatal("Gagal ping database", err)
	}

	fmt.Println("Berhasil terhubung ke database PostgreSQL!")
	return db
}