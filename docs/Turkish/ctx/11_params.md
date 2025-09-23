# Ctx/11 - `ctx.Param()` & `ctx.ParamInt()` & `ctx.ParamIntDefault()`

## `ctx.Param()`
`ctx.Param(key string) string` yöntemi, URL yolundaki belirli bir parametrenin değerini döndürür. Bu, dinamik rotalarla çalışırken kullanışlıdır.
### Örnek
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    // Dinamik rota tanımlama
    app.Handle("/@:username", func(ctx *flint.Context) {
        username := ctx.Param("username") // "id" parametresini al
        ctx.String(200, "Kullanıcı Adı : @%s", username)
    })

    app.Run(":8080")
}
```
### Açıklama
Bu kod, Flint sunucusu oluşturur, `/@:username` rotası için bir işleyici kaydeder ve `ctx.Param("username")` kullanarak URL'den kullanıcı adını alır.

http://localhost:8080/@coderianx

Cevap: `Kullanıcı adı: @coderianx`

---

## `ctx.ParamInt()`
`ctx.ParamInt(key string) (int, error)` yöntemi, URL yolundaki belirli bir parametrenin tamsayı değerini döndürür. Parametre tamsayıya dönüştürülemezse bir hata döner.

### Örnek
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    // Dinamik rota tanımlama
    app.Handle("/user/:id", func(ctx *flint.Context) {
        id, err := ctx.ParamInt("id") // "id" parametresini tamsayı olarak al
        if err != nil {
            ctx.String(400, "Geçersiz ID")
            return
        }
        ctx.String(200, "Kullanıcı ID: %d", id)
    })

    app.Run(":8080")
}
```
### Açıklama
Bu kod, Flint sunucusu oluşturur, `/user/:id` rotası için bir işleyici kaydeder ve `ctx.ParamInt("id")` kullanarak URL'den kullanıcı ID'sini tamsayı olarak alır.

http://localhost:8080/user/123

Cevap: `Kullanıcı ID: 123`