# 2 - Server Çalıştırma

```go
package main

import (
    "github.com/coderianx/flint"
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
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Run(":3000")
}
```
### Bu kod, yeni bir Flint sunucusu oluşturur ve 3000 numaralı portta çalıştırır.

### Not:
`Run()` içini boş bırakırsanız Flint, 8080 numaralı portu varsayılan olarak kullanır.

**Örnek;**
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Run() // Varsayılan olarak 8080 portunu kullanır.
}
```

Flint, varsayılan olarak 8080 numaralı portu kullanır. Eğer bu portu kullanmak istemiyorsanız, `Run()` fonksiyonuna istediğiniz port numarasını verebilirsiniz.