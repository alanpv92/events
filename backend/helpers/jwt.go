package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/alanpv92/events/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User, isRefresh bool) (string, error) {
	var expTime int64
	if isRefresh {
		expTime = time.Now().Add(time.Hour * 24).Unix()
	} else {
		expTime = time.Now().Add(time.Hour).Unix()
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   expTime,
	})
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyJwtToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	userDataMap := make(map[string]interface{})
	if err != nil {
		if err.Error() == "token has invalid claims: token is expired" {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, errors.New("could not get claims")
			}
			userDataMap["id"] = claims["id"].(string)
			userDataMap["email"] = claims["email"].(string)
			return userDataMap, errors.New("token expired")
		}

		return nil, errors.New("invalid token")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("could not get claims")
	}
	userDataMap["id"] = claims["id"].(string)
	userDataMap["email"] = claims["email"].(string)
	return userDataMap,nil
}
