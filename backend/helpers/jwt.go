package helpers

import (
	"os"
	"time"

	"github.com/alanpv92/events/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECREAT")))
	if err != nil {
		return "", err
	}

	return token, nil
}
