# Ctx 05 - `ctx.FormData()`

`ctx.FormData()` metodu Flint Framework'te `POST` istekleri aracılığıyla gönderilen form verilerini almak için kullanılır.

---

## 🧠 Kullanım Senaryosu

`ctx.FormData()` metodunu, HTML form gönderimlerini `POST` istekleri ile işlerken kullanın.
Kullanıcı giriş formları, kayıt formları, geri bildirim formları gibi durumlarda faydalıdır.
---

## ✅ Temel Kullanım

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/login", func(ctx *flint.Context) {
        if ctx.Post() {
            username := ctx.FormData("username")
            password := ctx.FormData("password")

            // Örnek cevap
            ctx.String("Başarılı")
        }
    })

    app.Run(":8080")
}
```

---

## 💡 Gelişmiş Örnek: Form İşleme

```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        if ctx.Get() {
            ctx.File("index.html")
            return
        }
        if ctx.Post() {
            username := ctx.FormData("username")
            password := ctx.FormData("password")

            // Örnek cevap
            ctx.String("Hoşgeldin")
        }
    })

    app.Run(":8080")
}
```

---
## 🌐 Örnek `index.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Flint Form Example</title>
</head>
<body>
  <h2>Login with Flint</h2>
  <form action="/" method="POST">
    <label>Username:</label><br>
    <input type="text" name="username" required><br><br>

    <label>Password:</label><br>
    <input type="password" name="password" required><br><br>

    <button type="submit">Submit</button>
  </form>
</body>
</html>
```

---

## 🔍 İlgili Metodlar

| Metod            | Açıklama                                |
|------------------|-----------------------------------------|
| `ctx.FormData(key string) string` | Form verisinden belirtilen anahtara karşılık gelen değeri alır. |
| `ctx.Post()`     | İstek `POST` metoduyla geldiyse `true` döner. |
| `ctx.Get()`      | İstek `GET` metoduyla geldiyse `true` döner. |
| `ctx.Method()`   | İstek yöntemini döndürür. |
| `ctx.String(msg string)` | İstek cevabını düz metin olarak döndürür. |
| `ctx.File(filename string)` | İstek cevabını dosya olarak döndürür. |
