# Ctx/02 | `ctx.File()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        ctx.File("example.html")
    })

    app.Run(":8080")
}
```

### 🧠 Bu kod, Flint sunucusu oluşturur, / rotası için bir işleyici kaydeder ve ctx.File() kullanarak bir dosyayla yanıt verir.