package flint

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(args ...string) string {
	if len(args) == 1 {
		data := args[0]
		sum := md5.Sum([]byte(data))
		return hex.EncodeToString(sum[:])
	} else if len(args) == 2 {
		data := args[0]
		salt := args[1]
		sum := md5.Sum([]byte(data + salt))
		return hex.EncodeToString(sum[:])
	}
	LogError("Md5", "expects 1 or 2 arguments")
	return ""
}
