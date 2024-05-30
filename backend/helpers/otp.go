package helpers

import (
	"math/rand"
	"strconv"
)

func GenerateOtp() string {
	var otp string
	for i := 0; i < 4; i++ {
		num := rand.Intn(10)
		otp += strconv.Itoa(num)
	}
	return otp
}
