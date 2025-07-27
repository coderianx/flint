package flint

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type FakeUser struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	HashedPwd string `json:"hashed_password"`
}

var (
	letters      = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	lowercase    = []rune("abcdefghijklmnopqrstuvwxyz")
	digits       = []rune("0123456789")
	emailDomains = []string{"@flintgo.dev", "@mail.com", "@example.com", "@fakemail.com"}
)

// Rastgele string üretme fonksiyonu
func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Gerçekçi email üretme fonksiyonu
func randomEmail() string {
	nameLength := rand.Intn(4) + 5 // 5-8 harf
	name := make([]rune, nameLength)
	for i := range name {
		name[i] = lowercase[rand.Intn(len(lowercase))]
	}

	// %50 ihtimalle birkaç sayı ekle
	if rand.Float32() < 0.5 {
		for i := 0; i < rand.Intn(3)+1; i++ {
			name = append(name, digits[rand.Intn(len(digits))])
		}
	}

	domain := emailDomains[rand.Intn(len(emailDomains))]
	return string(name) + domain
}

// FakeData fonksiyonu, sahte kullanıcı verisi oluşturup bir JSON dosyasına kaydeder
func FakeData(filename string, count int) {
	rand.Seed(time.Now().UnixNano())
	var users []FakeUser

	for i := 1; i <= count; i++ {
		username := "user" + strconv.Itoa(i) + "_" + randomString(3)
		email := randomEmail()

		rawPass := randomString(10)
		hash := sha256.Sum256([]byte(rawPass))
		hashedPass := hex.EncodeToString(hash[:])

		user := FakeUser{
			ID:        i,
			Username:  username,
			Email:     email,
			Password:  rawPass,
			HashedPwd: hashedPass,
		}

		users = append(users, user)
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("❌ Hata (dosya oluşturulamadı): %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(users); err != nil {
		fmt.Printf("❌ Hata (JSON yazılamadı): %v\n", err)
		return
	}

	fmt.Printf("✅ %d fake kullanıcı '%s' dosyasına yazıldı.\n", count, filename)
}
