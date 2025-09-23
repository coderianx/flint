package flint

import (
	"crypto/sha256"
	"database/sql"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

// FakeUser represents a fake user with extended fields.
type FakeUser struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	HashedPwd string    `json:"hashed_password"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	letters      = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	lowercase    = []rune("abcdefghijklmnopqrstuvwxyz")
	digits       = []rune("0123456789")
	firstNames   = []string{"Ali", "Eren", "Mert", "Ahmet", "Ayşe", "Zeynep", "John", "Sarah", "Emma", "David", "Efe", "Ronaldo", "Tony", "James"}
	lastNames    = []string{"Yılmaz", "Demir", "Kaya", "Smith", "Johnson", "Brown", "Lee", "Garcia"}
	streets      = []string{"Main St", "High St", "Park Ave", "2nd St", "Cedar Rd", "Elm St"}
	countries    = []string{"Turkey", "USA", "Germany", "UK", "Canada", "Azerbajian"}
	emailDomains = []string{"@flintgo.dev", "@mail.com", "@example.com", "@fakemail.com"}
)

// randomString generates a random string of the given length.
func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// randomEmail generates a realistic-looking random email address.
func randomEmail(name string) string {
	name = strings.ToLower(strings.ReplaceAll(name, " ", ""))
	if rand.Float32() < 0.3 {
		name += strconv.Itoa(rand.Intn(100))
	}
	domain := emailDomains[rand.Intn(len(emailDomains))]
	return name + domain
}

// randomPhone generates a fake phone number.
func randomPhone() string {
	return fmt.Sprintf("+%d-%03d-%03d-%04d",
		90+rand.Intn(60), // country code 90–150
		rand.Intn(900)+100,
		rand.Intn(900)+100,
		rand.Intn(10000),
	)
}

// randomAddress generates a fake address.
func randomAddress() string {
	street := streets[rand.Intn(len(streets))]
	return fmt.Sprintf("%d %s", rand.Intn(200)+1, street)
}

// FakeData generates fake users and saves as JSON or CSV.
func FakeData(filename string, count int, format string) {
	rand.Seed(time.Now().UnixNano())
	var users []FakeUser

	for i := 1; i <= count; i++ {
		first := firstNames[rand.Intn(len(firstNames))]
		last := lastNames[rand.Intn(len(lastNames))]
		fullName := first + " " + last
		username := strings.ToLower(first) + strconv.Itoa(i) + "_" + randomString(3)

		rawPass := randomString(10)
		hash := sha256.Sum256([]byte(rawPass))
		hashedPass := hex.EncodeToString(hash[:])

		user := FakeUser{
			ID:        i,
			Username:  username,
			Email:     randomEmail(fullName),
			Password:  rawPass,
			HashedPwd: hashedPass,
			FullName:  fullName,
			Phone:     randomPhone(),
			Address:   randomAddress(),
			Country:   countries[rand.Intn(len(countries))],
			CreatedAt: time.Now().AddDate(0, 0, -rand.Intn(1000)), // random past date
		}

		users = append(users, user)
	}

	switch strings.ToLower(format) {
	case "json":
		file, err := os.Create(filename)
		if err != nil {
			LogError("FakeData", err.Error())
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(users); err != nil {
			LogError("FakeData", err.Error())
			return
		}
	case "csv":
		file, err := os.Create(filename)
		if err != nil {
			LogError("FakeData", err.Error())
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		// header
		writer.Write([]string{"ID", "Username", "Email", "FullName", "Phone", "Address", "Country", "CreatedAt", "Password", "HashedPwd"})

		// data
		for _, u := range users {
			writer.Write([]string{
				strconv.Itoa(u.ID), u.Username, u.Email, u.FullName, u.Phone, u.Address,
				u.Country, u.CreatedAt.Format(time.RFC3339), u.Password, u.HashedPwd,
			})
		}
	default:
		LogError("FakeData", "Unsupported format: "+format)
		return
	}

	fmt.Printf("✅ %d fake users written to '%s' (%s).\n", count, filename, format)
}

// FakeDataDB creates/opens a sqlite database file (modernc.org/sqlite) and inserts `count` fake users.
// filename: path to .db file (e.g. "users.db")
// count: number of fake users to insert
func FakeDataDB(filename string, count int) {
	rand.Seed(time.Now().UnixNano())

	// Open (or create) sqlite database using modernc.org/sqlite driver
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		LogError("FakeDataDB", err.Error())
		return
	}
	defer db.Close()

	// Create users table if not exists
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username TEXT,
		email TEXT,
		password TEXT,
		hashed_password TEXT,
		full_name TEXT,
		phone TEXT,
		address TEXT,
		country TEXT,
		created_at TEXT
	);`
	if _, err := db.Exec(createTable); err != nil {
		LogError("FakeDataDB", err.Error())
		return
	}

	tx, err := db.Begin()
	if err != nil {
		LogError("FakeDataDB", err.Error())
		return
	}

	stmt, err := tx.Prepare(`INSERT INTO users 
		(id, username, email, password, hashed_password, full_name, phone, address, country, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		LogError("FakeDataDB", err.Error())
		_ = tx.Rollback()
		return
	}
	defer stmt.Close()

	for i := 1; i <= count; i++ {
		first := firstNames[rand.Intn(len(firstNames))]
		last := lastNames[rand.Intn(len(lastNames))]
		fullName := first + " " + last
		username := strings.ToLower(first) + strconv.Itoa(i) + "_" + randomString(3)

		rawPass := randomString(10)
		hash := sha256.Sum256([]byte(rawPass))
		hashedPass := hex.EncodeToString(hash[:])

		createdAt := time.Now().AddDate(0, 0, -rand.Intn(1000)).Format(time.RFC3339)

		if _, err := stmt.Exec(i, username, randomEmail(fullName), rawPass, hashedPass, fullName, randomPhone(), randomAddress(), countries[rand.Intn(len(countries))], createdAt); err != nil {
			LogError("FakeDataDB", err.Error())
			_ = tx.Rollback()
			return
		}
	}

	if err := tx.Commit(); err != nil {
		LogError("FakeDataDB", err.Error())
		return
	}

	fmt.Printf("✅ %d fake users inserted into SQLite DB '%s'.\n", count, filename)
}
