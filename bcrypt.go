package flint

import (
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(args ...string) string {
	if len(args) == 1 {
		data := args[0]
		hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
		if err != nil {
			LogError("Bcrypt", err.Error())
			return ""
		}
		return string(hash)
	} else if len(args) == 2 {
		data := args[0]
		salt := args[1]
		hash, err := bcrypt.GenerateFromPassword([]byte(data+salt), bcrypt.DefaultCost)
		if err != nil {
			LogError("Bcrypt", err.Error())
			return ""
		}
		return string(hash)
	}
	LogError("Bcrypt", "expects 1 or 2 arguments")
	return ""
}

func CompareBcrypt(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
