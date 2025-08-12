package test

import (
	"testing"

	"github.com/coderianx/flint"
)

// TestHashFunctions (Hash fonksiyonlarını test eder)
func TestHashFunctions(t *testing.T) {
	password := "mypassword" // Test password (Test parolası)
	salt := "mysalt"         // Test salt (Test tuzu)

	// Generate hashes (Hash'leri oluştur)
	hash_md5 := flint.Md5(password)                 // MD5 hash without salt (Tuzsuz MD5 hash)
	salt_md5 := flint.Md5(password, salt)           // MD5 hash with salt (Tuzlu MD5 hash)
	hash_sha256 := flint.Sha256(password)           // SHA256 hash without salt (Tuzsuz SHA256 hash)
	salt_sha256 := flint.Sha256(password, salt)     // SHA256 hash with salt (Tuzlu SHA256 hash)
	hash_sha512 := flint.Sha512(password)           // SHA512 hash without salt (Tuzsuz SHA512 hash)
	salt_sha512 := flint.Sha512(password, salt)     // SHA512 hash with salt (Tuzlu SHA512 hash)
	salt_sha3_256 := flint.Sha3_256(password)       // SHA3-256 hash (SHA3-256 hash)
	salt_sha3_512 := flint.Sha3_512(password, salt) // SHA3-512 hash with salt (Tuzlu SHA3-512 hash)
	hash_bcrypt := flint.Bcrypt(password, 12)       // Bcrypt hash with cost 12 (Bcrypt hash, maliyet 12)
	hash_argon2, _ := flint.Argon2Hash(password)    // Argon2 hash (Argon2 hash)
	hash_blake2b := flint.Blake2b(password)         // Blake2b hash (Blake2b hash)
	hash_blake2s := flint.Blake2s(password)         // Blake2s hash (Blake2s hash)

	// Log results (Sonuçları yazdır)
	t.Logf("MD5 Hash: %s", hash_md5)
	t.Logf("MD5 Salted: %s", salt_md5)
	t.Logf("SHA256 Hash: %s", hash_sha256)
	t.Logf("SHA256 Salted: %s", salt_sha256)
	t.Logf("SHA512 Hash: %s", hash_sha512)
	t.Logf("SHA512 Salted: %s", salt_sha512)
	t.Logf("SHA3-256 Hash: %s", salt_sha3_256)
	t.Logf("SHA3-512 Salted: %s", salt_sha3_512)
	t.Logf("Bcrypt Hash: %s", hash_bcrypt)
	t.Logf("Argon2 Hash: %s", hash_argon2)
	t.Logf("Blake2b Hash: %s", hash_blake2b)
	t.Logf("Blake2s Hash: %s", hash_blake2s)
}
