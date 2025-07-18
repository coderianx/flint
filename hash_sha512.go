package flint

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512(args ...string) string {
	if len(args) == 1 {
		data := args[0]
		sum := sha512.Sum512([]byte(data))
		return hex.EncodeToString(sum[:])
	} else if len(args) == 2 {
		data := args[0]
		salt := args[1]
		sum := sha512.Sum512([]byte(data + salt))
		return hex.EncodeToString(sum[:])
	}
	LogError("Sha512", "expects 1 or 2 arguments")
	return ""
}
