package database

import (
	"errors"

	"github.com/alanpv92/events/models"
)

func authUpMigrations() {
	query := `CREATE TABLE IF NOT EXISTS users(
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		email TEXT NOT NULL UNIQUE,
		user_name TEXT NOT NULL,
		password TEXT NOT NULL,
		email_verified BOOLEAN DEFAULT false
	)`
	_, err := Db.Exec(query)
	if err != nil {

		panic("could not create users table")
	}
}

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email=$1"
	res, err := Db.Query(query, email)
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	var user models.User
	isUserPresent := res.Next()
	if !isUserPresent {
		return nil, nil
	}
	err = res.Scan(&user.Id, &user.Email, &user.UserName, &user.Password, &user.EmailVerified)
	if err != nil {

		return nil, errors.New("something went wrong")
	}
	return &user, nil
}

func InsertUser(user models.User) (string, error) {
	query := `INSERT INTO users(user_name,email,password)
	VALUES($1,$2,$3) RETURNING id;
	`
	var id string
	row := Db.QueryRow(query, user.UserName, user.Email, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil

}
