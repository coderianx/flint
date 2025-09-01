# Ctx/01 | `ctx.String()`

```go
package main

import (
    "github.com/coderianx/flint"
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

## `ctx.Stringf()`
```go
package main

import (
    flint "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        name := "Flint"
        ctx.Stringf("Merhaba %s!", name)

        // veya status code ile formatlanmış string gönderilebilir
        ctx.Stringf(200, "Merhaba %s!", name)
    })

    app.Run(":8080")
}
```
# 🧠 Bu kod, Flint sunucusu oluşturur, / rotası için bir işleyici kaydeder ve ctx.Stringf() kullanarak formatlanmış bir string ile yanıt verir.

