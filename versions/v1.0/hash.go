package flint

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

func Argon2Hash(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32) // t=1, m=64MB, p=4, hashlen=32

	encodedHash := fmt.Sprintf("$argon2id$v=19$m=65536,t=1,p=4$%s$%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return encodedHash, nil
}

func Argon2Verify(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("geçersiz hash formatı")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	comparison := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, uint32(len(hash)))
	return string(comparison) == string(hash), nil
}

func Blake2b(data string) string {
	hash := blake2b.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

func Blake2s(data string) string {
	hash := blake2s.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func Bcrypt(password string, optionalCost ...int) string {
	cost := bcrypt.DefaultCost // default: 10

	if len(optionalCost) > 1 {
		LogError("Bcrypt()", "Too many arguments. Expects: Bcrypt(password, [optional cost])")
		return ""
	}

	if len(optionalCost) == 1 {
		if optionalCost[0] < bcrypt.MinCost || optionalCost[0] > bcrypt.MaxCost {
			LogError("Bcrypt()", "Cost must be between 4 and 31")
			return ""
		}
		cost = optionalCost[0]
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		LogError("Bcrypt()", err.Error())
		return ""
	}

	return string(hash)
}

func CompareBcrypt(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

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

func Sha3_256(args ...string) string {
	if len(args) == 1 {
		data := args[0]
		hash := sha3.New256()          // SHA3-256 hash fonksiyonunu oluştur
		hash.Write([]byte(data))       // Veriyi hash fonksiyonuna yaz
		sum := hash.Sum(nil)           // Hash sonucunu al
		return hex.EncodeToString(sum) // Hex formatına çevir ve döndür
	} else if len(args) == 2 {
		data := args[0]
		salt := args[1]
		hash := sha3.New256()           // SHA3-256 hash fonksiyonunu oluştur
		hash.Write([]byte(data + salt)) // Veri ve salt'ı hash fonksiyonuna yaz
		sum := hash.Sum(nil)            // Hash sonucunu al
		return hex.EncodeToString(sum)  // Hex formatına çevir ve döndür
	}
	LogError("Sha3_256", "expects 1 or 2 argument")
	return ""
}

func Sha3_512(args ...string) string {
	switch len(args) {
	case 1:
		data := args[0]
		hash := sha3.New512()
		hash.Write([]byte(data))
		sum := hash.Sum(nil)
		return hex.EncodeToString(sum)
	case 2:
		data := args[0]
		salt := args[1]
		hash := sha3.New512()
		hash.Write([]byte(data + salt))
		sum := hash.Sum(nil)
		return hex.EncodeToString(sum)
	default:
		LogError("Sha3_512", "expects 1 or 2 argument")
		return ""
	}
}

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
	LogError("Md5()", "expects 1 or 2 arguments")
	return ""
}
