# Ctx/01 | `ctx.String()`

```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        ctx.String("Merhaba, Flint!")

        // veya status code ile string gönderilebilir
        ctx.String(200, "Merhaba Flint!")
    })

    app.Run(":8080")
}
```
# 🧠 Bu kod, Flint sunucusu oluşturur, / rotası için bir işleyici kaydeder ve ctx.String() kullanarak düz metinle yanıt verir.

