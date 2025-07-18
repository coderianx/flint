package flint

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(args ...string) string {
	if len(args) == 1 {
		data := args[0]
		sum := sha256.Sum256([]byte(data))
		return hex.EncodeToString(sum[:])
	} else if len(args) == 2 {
		data := args[0]
		salt := args[1]
		sum := sha256.Sum256([]byte(data + salt))
		return hex.EncodeToString(sum[:])
	}
	LogError("Sha256", "expects 1 or 2 arguments")
	return ""
}
