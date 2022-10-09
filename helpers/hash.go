package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(value string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func CompareHash(hashedString, valueString string) error {
	hashed := []byte(hashedString)
	value := []byte(valueString)
	err := bcrypt.CompareHashAndPassword(hashed, value)
	return err
}
