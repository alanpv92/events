package database

import (
	"fmt"

	"github.com/alanpv92/events/models"
)

func authUpMigrations() {
	query := `CREATE TABLE IF NOT EXISTS users(
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		email TEXT NOT NULL UNIQUE,
		user_name TEXT NOT NULL,
		password TEXT NOT NULL
	)`
	_, err := Db.Exec(query)
	if err != nil {
		fmt.Println(err);
		panic("could not create users table")
	}
}

func RegisterUser(user models.User){
	
}
