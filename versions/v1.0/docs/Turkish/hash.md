
# 🔐 Hash - Flint (Türkçe Dokümantasyon)

Flint, aşağıdaki kriptografik hash algoritmalarını **yerleşik olarak destekler**:

- `MD5`  
- `SHA256`, `SHA512`  
- `SHA3-256`, `SHA3-512`  
- `Bcrypt`  
- `BLAKE2b`, `BLAKE2s`  
- `Argon2` (şifre saklama için)

---

## 🧩 Ne İçin Kullanılır?

- Şifre hashleme ve doğrulama  
- Veri bütünlüğü kontrolleri  
- Hassas bilgilerin güvenli saklanması  

---

## 🧮 MD5

> ⚠️ **Uyarı:** Kriptografik olarak güvenli değildir. Sadece bütünlük kontrolü için önerilir.

### Örnek:
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "guvenli_sifre"
    hashed := flint.HashMD5(password)
    fmt.Println("MD5 Hash:", hashed)
}
```

### Salt ile:
```go
func main() {
    password := "guvenli_sifre"
    salt := "rastgele_salt"
    hashed := flint.Md5(password, salt)
    fmt.Println("MD5 Hash + Salt:", hashed)
}
```

---

## 🔐 SHA256

> ✅ Güçlü, güvenli ve yaygın olarak kullanılır.

### Örnek:
```go
func main() {
    password := "guvenli_sifre"
    hashed := flint.Sha256(password)
    fmt.Println("SHA256 Hash:", hashed)
}
```

### Salt ile:
```go
func main() {
    password := "guvenli_sifre"
    salt := "rastgele_salt"
    hashed := flint.Sha256Salt(password, salt)
    fmt.Println("SHA256 Hash + Salt:", hashed)
}
```

---

## 🔐 SHA512

> ✅ SHA256'ya göre daha uzun ve güvenli bir algoritmadır.

### Örnek:
```go
func main() {
    password := "guvenli_sifre"
    hashed := flint.Sha512(password)
    fmt.Println("SHA512 Hash:", hashed)
}
```

### Salt ile:
```go
func main() {
    password := "guvenli_sifre"
    salt := "rastgele_salt"
    hashed := flint.Sha512(password, salt)
    fmt.Println("SHA512 Hash + Salt:", hashed)
}
```

---

## 🔐 SHA3-256

> ✅ SHA3 ailesine ait yeni nesil hash algoritmasıdır.

### Örnek:
```go
func main() {
    password := "guvenli_sifre"
    hashed := flint.Sha3_256(password)
    fmt.Println("SHA3-256 Hash:", hashed)
}
```

### Salt ile:
```go
func main() {
    password := "guvenli_sifre"
    salt := "rastgele_salt"
    hashed := flint.Sha3_256(password, salt)
    fmt.Println("SHA3-256 Hash + Salt:", hashed)
}
```

---

## 🔐 SHA3-512

> ✅ SHA3 ailesinin en uzun ve güvenli versiyonu.

### Örnek:
```go
func main() {
    password := "guvenli_sifre"
    hashed := flint.Sha3_512(password)
    fmt.Println("SHA3-512 Hash:", hashed)
}
```

### Salt ile:
```go
func main() {
    password := "guvenli_sifre"
    salt := "rastgele_salt"
    hashed := flint.Sha3_512(password, salt)
    fmt.Println("SHA3-512 Hash + Salt:", hashed)
}
```

---

## 🛡️ Bcrypt

> 🐢 Yavaş çalışır ama çok güvenlidir. Şifre saklama için önerilir. Salt işlemini kendisi yapar.

### Hashleme:
```go
func main() {
    password := "guvenli_sifre"
    hashed, err := flint.Bcrypt(password, 12)
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    fmt.Println("Bcrypt Hash:", hashed)
}
```

### Doğrulama:
```go
func main() {
    password := "guvenli_sifre"
    hashed, _ := flint.Bcrypt(password, 12)
    isValid := flint.CompareBcrypt(password, hashed)

    if isValid {
        fmt.Println("Şifre Doğru!")
    } else {
        fmt.Println("Şifre Hatalı!")
    }
}
```

---

## ⚡ BLAKE2s

> ⚡ 32-bit sistemler için optimize edilmiştir. Çok hızlıdır.

### Örnek:
```go
func main() {
    password := "guvenli_sifre"
    hashed := flint.Blake2s(password)
    fmt.Println("BLAKE2s Hash:", hashed)
}
```

---

## ⚡ BLAKE2b

> ⚡ 64-bit sistemlerde hızlı ve güvenli bir algoritmadır.

### Örnek:
```go
func main() {
    password := "guvenli_sifre"
    hashed := flint.Blake2b(password)
    fmt.Println("BLAKE2b Hash:", hashed)
}
```

---

## 📊 Algoritma Karşılaştırma Tablosu

| Algoritma    | Güvenlik Seviyesi     | Hız         | Uygun Kullanım                  |
|--------------|------------------------|-------------|----------------------------------|
| **MD5**      | ❌ Çok Düşük           | ⚡ Çok Hızlı | Kontrol toplamı, eski sistemler |
| **SHA256**   | ✅ Yüksek              | ⚡ Hızlı     | Genel veri bütünlüğü           |
| **SHA512**   | ✅ Yüksek              | ⚡ Hızlı     | Dosya kontrolü, yedekleme      |
| **SHA3**     | ✅✅ Çok Yüksek        | 🚀 Orta      | Kriptografik güvenlik          |
| **BLAKE2**   | ✅✅ Çok Yüksek        | ⚡⚡ Çok Hızlı| Modern güvenli hash işlemleri  |
| **Argon2**   | 🛡️ En Yüksek          | 🐢 Yavaş     | Şifre saklama (önerilen)       |
