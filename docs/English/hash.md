# Hash - Flint

Flint provides built-in support for cryptographic hashing algorithms including:

- `MD5`
- `SHA256`, `SHA512`
- `SHA3-256`, `SHA3-512`
- `Bcrypt`
- `BLAKE2b`, `BLAKE2s`
- `Argon2` (secure password hashing)

These functions are useful for:

- Password hashing and verification  
- Data integrity checks  
- Secure storage of sensitive information

You can access these functions directly through the Flint package and use them in your application logic.

---

## MD5

> **Note:** MD5 is a widely used hashing algorithm that produces a 128-bit hash value.  
> It is commonly used for checksums and data integrity, but **not recommended** for cryptographic security due to known vulnerabilities.

### Example;
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"

    // Hash the password using MD5
    hashedPassword := flint.HashMD5(password)
    // Output the hashed password
    fmt.Println("MD5 Hashed Password:", hashedPassword)
    fmt.Println("Password", password)
}
```
### With Salt;
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"
    salt := "random_salt_value"

    // Hash the password with salt using MD5
    hashedPassword := flint.Md5(password, salt)
    // Output the hashed password
    fmt.Println("MD5 Hashed Password with Salt:", hashedPassword)
    fmt.Println("Password", password)
    fmt.Println("Salt", salt)
}

```
### 🧠 Explanation:
This code demonstrates how to hash a password using the MD5 algorithm provided by Flint. The `HashMD5` function takes a string input and returns its MD5 hash.

## SHA256
> **Note:** SHA256 is a cryptographic hash function that produces a 256-bit hash value.
> It is widely used for data integrity and security, providing a higher level of security than MD5.
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"

    // Hash the password using SHA256
    hashedPassword := flint.Sha256(password)
    // Output the hashed password
    fmt.Println("SHA256 Hashed Password:", hashedPassword)
    fmt.Println("Password", password)
}
```
### With Salt;
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"
    salt := "random_salt_value"

    // Hash the password with salt using SHA256
    hashedPassword := flint.Sha256Salt(password, salt)
    // Output the hashed password
    fmt.Println("SHA256 Hashed Password with Salt:", hashedPassword)
    fmt.Println("Password", password)
    fmt.Println("Salt", salt)
}
```
### 🧠 Explanation:
This code demonstrates how to hash a password using the SHA256 algorithm provided by Flint. The `Sha256` function takes a string input and returns its SHA256 hash.

## SHA512
> **Note:** SHA512 is a cryptographic hash function that produces a 512-bit hash value.
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"

    // Hash the password using SHA512
    hashedPassword := flint.Sha512(password)
    // Output the hashed password
    fmt.Println("SHA512 Hashed Password:", hashedPassword)
    fmt.Println("Password", password)
}
```

### With Salt;
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"
    salt := "random_salt_value"

    // Hash the password with salt using SHA512
    hashedPassword := flint.Sha512(password, salt)
    // Output the hashed password
    fmt.Println("SHA512 Hashed Password with Salt:", hashedPassword)
    fmt.Println("Password", password)
    fmt.Println("Salt", salt)
}
```
### 🧠 Explanation:
This code demonstrates how to hash a password using the SHA512 algorithm provided by Flint. The `Sha512` function takes a string input and returns its SHA512 hash.

## SHA3-256
> **Note:** SHA3-256 is part of the SHA-3 family of cryptographic hash functions, providing a higher level of security than SHA2.
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"

    // Hash the password using SHA3-256
    hashedPassword := flint.Sha3_256(password)
    // Output the hashed password
    fmt.Println("SHA3-256 Hashed Password:", hashedPassword)
    fmt.Println("Password", password)
}
```

### With Salt;
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"
    salt := "random_salt_value"

    // Hash the password with salt using SHA3-256
    hashedPassword := flint.Sha3_256(password, salt)
    // Output the hashed password
    fmt.Println("SHA3-256 Hashed Password with Salt:", hashedPassword)
    fmt.Println("Password", password)
    fmt.Println("Salt", salt)
}
```
### 🧠 Explanation:
This code demonstrates how to hash a password using the SHA3-256 algorithm provided by Flint. The `Sha3_256` function takes a string input and returns its SHA3-256 hash.

## SHA3-512
> **Note:** SHA3-512 is part of the SHA-3 family of cryptographic hash functions, providing a higher level of security than SHA2.
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"

    // Hash the password using SHA3-512
    hashedPassword := flint.Sha3_512(password)
    // Output the hashed password
    fmt.Println("SHA3-512 Hashed Password:", hashedPassword)
    fmt.Println("Password", password)
}
```
### With Salt;
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"
    salt := "random_salt_value"

    // Hash the password with salt using SHA3-512
    hashedPassword := flint.Sha3_512(password, salt)
    // Output the hashed password
    fmt.Println("SHA3-512 Hashed Password with Salt:", hashedPassword)
    fmt.Println("Password", password)
    fmt.Println("Salt", salt)
}
```

### 🧠 Explanation:
This code demonstrates how to hash a password using the SHA3-512 algorithm provided by Flint. The `Sha3_512` function takes a string input and returns its SHA3-512 hash.

## Bcrypt
> **Note:** Bcrypt is a password hashing function designed to be slow and resistant to brute-force attacks. It automatically handles salting and is widely used for secure password storage.
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"

    // Hash the password using Bcrypt
    hashedPassword, err := flint.Bcrypt(password, 12)
    if err != nil {
        fmt.Println("Error hashing password:", err)
        return
    }
    // Output the hashed password
    fmt.Println("Bcrypt Hashed Password:", hashedPassword)
    fmt.Println("Password", password)
}
```
### Verification;
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"
    hashedPassword, _ := flint.Bcrypt(password, 12)

    // Verify the password against the hashed password
    isValid := flint.CompareBcrypt(password, hashedPassword)
    if isValid {
        fmt.Println("Password is valid!")
    } else {
        fmt.Println("Invalid password.")
    }
}
```
### 🧠 Explanation:
This code demonstrates how to hash a password using the Bcrypt algorithm provided by Flint. The `Bcrypt` function takes a string input and a cost factor (12 in this case) and returns the hashed password. The `CompareBcrypt` function is used to verify the password against the hashed value.


## BLAKE2s
> **Note:** BLAKE2 is a cryptographic hash function that is faster than MD5, SHA-1, and SHA-2 while providing a high level of security. BLAKE2s is optimized for 32-bit platforms.
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"

    // Hash the password using BLAKE2s
    hashedPassword := flint.Blake2s(password)
    // Output the hashed password
    fmt.Println("BLAKE2s Hashed Password:", hashedPassword)
    fmt.Println("Password", password)
}
```
### 🧠 Explanation:
This code demonstrates how to hash a password using the BLAKE2s algorithm provided by Flint. The `Blake2s` function takes a string input and returns its BLAKE2s hash.

## BLAKE2b
> **Note:** BLAKE2b is a cryptographic hash function that is faster than MD5, SHA-1, and SHA-2 while providing a high level of security. BLAKE2b is optimized for 64-bit platforms.
```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    password := "my_secure_password"

    // Hash the password using BLAKE2b
    hashedPassword := flint.Blake2b(password)
    // Output the hashed password
    fmt.Println("BLAKE2b Hashed Password:", hashedPassword)
    fmt.Println("Password", password)
}
```
### 🧠 Explanation:
This code demonstrates how to hash a password using the BLAKE2b algorithm provided by Flint. The `Blake2b` function takes a string input and returns its BLAKE2b hash.


| Algorithm | Security        | Speed       | Suitable For           |
|-----------|------------------|-------------|--------------------------|
| **MD5**   | ❌ Very Low      | ⚡ Very Fast | Checksums, legacy code  |
| **SHA256**| ✅ High          | ⚡ Fast      | General-purpose hashing |
| **SHA512**| ✅ High          | ⚡ Fast      | File integrity, backups |
| **SHA3**  | ✅✅ Very High   | 🚀 Moderate | Cryptographic use       |
| **BLAKE2**| ✅✅ Very High   | ⚡⚡ Very Fast| Modern general hashing  |
| **Argon2**| 🛡️ Top Level    | 🐢 Slow      | Password storage        |
