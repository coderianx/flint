# Ctx/03 | `ctx.JSON()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        data := map[string]string{
            "message": "Hello, World!",
        }
        ctx.JSON(200, data)
    })

    app.Run(":8080")
}
```

### 🧠 Bu kod, Flint sunucusu oluşturur, / rotası için bir işleyici kaydeder ve `ctx.JSON()` kullanarak JSON formatında bir yanıt verir.

---

## `ctx.JSONPretty()`

### Eğer okunabilir bir json yanıtı vermek isterseniz `ctx.JSONPretty()` fonksiyonunu kullanabilirsiniz:

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        data := map[string]string{
            "message": "Hello, World!",
            "status":  200,
        }
        ctx.JSONPretty(200, data)
        // veya
        ctx.JSONPretty(200, data, 4) // 4 indentation boşluğu kullanır
    })
}
```

### 🧠 Bu kod, Flint sunucusu oluşturur, / rotası için bir işleyici kaydeder ve `ctx.JSONPretty()` kullanarak okunabilir JSON formatında bir yanıt verir.

---
## `JSONFile()`

### Eğer bir json dosyasını yanıt olarak göndermek isterseniz `ctx.JSONFile()` fonksiyonunu kullanabilirsiniz:

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        ctx.JSONFile(200, "data.json")
    })

    app.Run(":8080")
}
```