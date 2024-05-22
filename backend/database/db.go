package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Init() {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic("could not open database")
	}
	err = db.Ping()
	if err != nil {
		panic("could not open database")
	}
	Db = db
	migrations()
}

func migrations() {
	upMigrations()
}

func upMigrations() {
	authUpMigrations()
}
