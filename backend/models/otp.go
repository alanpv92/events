package models

import "errors"

type Otp struct{
	Otp string `json:"otp"`
}

func(otp Otp) Validate() error{
	if otp.Otp==""{
		return errors.New("no otp")
	}
	return nil
}