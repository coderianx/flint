# 2 - Server Çalıştırma

```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    app.Run(":8080")
}
```
### Bu kod, yeni bir Flint sunucusu oluşturur ve 8080 numaralı portta çalıştırır.

**Farklı portlarda bir sunucu açabilirsiniz, Run() içindeki port numarasını değiştirerek.**

### Örnek;

```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    app.Run(":3000")
}
```
### Bu kod, yeni bir Flint sunucusu oluşturur ve 3000 numaralı portta çalıştırır.