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