package models

import "errors"

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

func (user User) ValidateLoginBody() error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (user User) ValidateRegisterBody() error {
	err := user.ValidateLoginBody()
	if err != nil {
		return err
	}
	if user.UserName == "" {
		return errors.New("username is required")
	}
	return nil
}

func (user User) AuthResponse() map[string]interface{} {
	responseMap := make(map[string]interface{})
	responseMap["id"] = user.Id
	responseMap["user_name"] = user.UserName
	responseMap["email"] = user.Email
	responseMap["token"] = user.Token
	return responseMap
}
