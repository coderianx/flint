
# 🔹 Context 04 – `ctx.Get()` & `ctx.Post()` & `ctx.Delete` & `ctx.Put()` (Türkçe)

Bu metodlar, Flint Framework içinde gelen HTTP isteğinin metodunu kontrol etmek için kullanılır.

---

- `ctx.Get()` → İstek `GET` metoduyla geldiyse `true` döner.  
- `ctx.Post()` → İstek `POST` metoduyla geldiyse `true` döner.
- `ctx.Delete()` → İstek `DELETE` metoduyla geldiyse `true` döner.
- `ctx.Put()` → İstek `PUT` metoduyla geldiyse `true` döner. 
 

---

## 🧠 Kullanım Senaryosu

Bu metodlar, aynı route üzerinde hem GET (sayfa gösterimi) hem de POST (form işlemleri) gibi birden fazla HTTP metodunu yönetmek için kullanışlıdır.

---

## ✅ Örnek: `ctx.Get()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        if ctx.Get() {
            // GET isteği işleme kodu
        }
    })

    app.Run(":8080")
}
```

### 🧠 Açıklama:
Sunucu oluşturulur, `/` route'u tanımlanır ve gelen isteğin GET olup olmadığı `ctx.Get()` ile kontrol edilir.

---

## ✅ Örnek: `ctx.Post()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        if ctx.Post() {
            // POST isteği işleme kodu
        }
    })

    app.Run(":8080")
}
```

### 🧠 Açıklama:
İstek `POST` metoduyla geldiyse `ctx.Post()` fonksiyonu `true` döner ve ilgili işlem yapılır.

---

## ✅ Örnek: `ctx.Delete()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        if ctx.Delete() {
            // DELETE isteği işleme kodu
        }
    })

    app.Run(":8080")
```

### 🧠 Açıklama:
İstek `DELETE` metoduyla geldiyse `ctx.Delete()` fonksiyonu `true` döner ve ilgili işlem yapılır.

---

## ✅ Örnek: `ctx.Put()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        if ctx.Put() {
            // PUT isteği işleme kodu
        }
    })

    app.Run(":8080")
}
```
### 🧠 Açıklama:
İstek `PUT` metoduyla geldiyse `ctx.Put()` fonksiyonu `true` döner ve ilgili işlem yapılır.

---

## 💡 Kombine Örnek: GET + POST

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/submit", func(ctx *flint.Context) {
        if ctx.Get() {
            ctx.File("index.html") // HTML sayfası döner
        }
        if ctx.Post() {
            // Form gönderimi işlenir
        }
    })

    app.Run(":8080")
}
```

---

## 🔍 İlgili Metodlar

| Metod            | Açıklama                                |
|------------------|------------------------------------------|
| `ctx.Get()`      | HTTP metodunun GET olup olmadığını kontrol eder |
| `ctx.Delete()`   | HTTP metodunun DELETE olup olmadığını kontrol eder |
| `ctx.Put()`      | HTTP metodunun PUT olup olmadığını kontrol eder |
| `ctx.Post()`     | HTTP metodunun POST olup olmadığını kontrol eder |
| `ctx.File(path)` | Statik dosya (örneğin HTML) sunar        |
| `ctx.String()`   | Düz metin yanıt döner                    |

---

## ⚠️ Notlar

- Bu metodlar, route fonksiyonlarını HTTP metoduna göre ayırmanıza olanak sağlar.
- Özellikle `/login`, `/register`, `/submit` gibi formların olduğu sayfalarda kullanışlıdır.
