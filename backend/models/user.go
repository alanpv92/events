package models

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}
