package database

import (
	"errors"
	"fmt"
	"time"
)

func otpUpMigrations() {
	query := `CREATE TABLE IF NOT EXISTS otps (
		id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id uuid REFERENCES users(id) UNIQUE,
		otp TEXT NOT NULL ,
		created_at DATE DEFAULT CURRENT_DATE
	)`
	_, err := Db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func DeleteOtp(userId string, shouldWait bool) {
	if shouldWait {
		time.Sleep(time.Second * 500)
	}

	query := `DELETE FROM otps WHERE user_id = $1`
	_, err := Db.Exec(query, userId)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("could not delete otp")
		return
	}
	fmt.Println("otp deleted")
}

func AddOtp(otp string, userId string) error {
	query := `INSERT INTO otps(user_id,otp) 
   VALUES(
    $1,$2
   ) ON CONFLICT(user_id) DO UPDATE SET otp=EXCLUDED.otp `
	_, err := Db.Exec(query, userId, otp)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("could not generate otp")
	}
	go DeleteOtp(userId, true)
	return nil
}

func VerifyOtp(id string) (string, error) {
	query := `SELECT otp FROM otps WHERE user_id=$1`
	rows := Db.QueryRow(query, id)
	var dbOtp string
	err := rows.Scan(&dbOtp)
	if err != nil {
		return "false", errors.New("could not get otp")
	}
    DeleteOtp(id,false)
	return dbOtp, nil
}
