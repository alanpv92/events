package helpers

import "golang.org/x/crypto/bcrypt"

func HashPasswod(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", nil
	}
	return string(hashByte), nil
}
