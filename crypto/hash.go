package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(val string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), err
}

func CheckHash(val string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(val))

	return err == nil
}
