package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/coderianx/flint"
)

// TestFakeData (Sahte veri üretimini test eder)
func TestFakeData(t *testing.T) {
	outputFile := "data.json" // Permanent file name (Kalıcı dosya adı)

	// Remove the file if it exists before the test (Testten önce dosya varsa silelim)
	_ = os.Remove(outputFile)

	// Generate fake data (Sahte veriyi üret)
	flint.FakeData(outputFile, 5)

	// Check if the file exists (Dosya var mı kontrol et)
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Fatalf("Expected file %s to be created (Beklenen dosya %s oluşturulmadı)", outputFile, outputFile)
	}

	// Read file content (Dosya içeriğini oku)
	data, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("File could not be read (Dosya okunamadı): %v", err)
	}

	// Parse JSON (JSON parse et)
	var users []flint.FakeUser
	if err := json.Unmarshal(data, &users); err != nil {
		t.Fatalf("JSON parse error (JSON parse hatası): %v", err)
	}

	// Check if the number of users is correct (Kullanıcı sayısı doğru mu kontrol et)
	if len(users) != 5 {
		t.Errorf("Expected 5 users, got %d (Beklenen kullanıcı sayısı 5, ama %d)", len(users), len(users))
	}
}
