package helpers

import "golang.org/x/crypto/bcrypt"

func HashPasswod(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", nil
	}
	return string(hashByte), nil
}

func VerifyPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
