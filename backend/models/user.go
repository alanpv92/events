package models

import "errors"

type User struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	UserName     string `json:"user_name"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	EmailVerified bool 	`json:"email_verified"`
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
	responseMap["refresh_token"] = user.RefreshToken
	responseMap["email_verified"]=user.EmailVerified
	return responseMap
}

func(user User) TokenResponse()map[string]interface{}{
	responseMap := make(map[string]interface{})
	responseMap["token"] = user.Token
	responseMap["refresh_token"] = user.RefreshToken
	return responseMap;
}

func (user User) ValidateRefreshToken() error {
	if user.RefreshToken == "" {
		return errors.New("refresh token is required")
	}
	return nil
}
