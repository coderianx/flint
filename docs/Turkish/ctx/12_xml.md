# Ctx12 | `ctx.XML()` & `ctx.XMLPretty()` & `ctx.XMLFile()`

## `ctx.XML()``
**Açıklama:** `ctx.XML()` fonksiyonu, belirtilen veri yapısını XML formatında serileştirir ve HTTP yanıtı olarak gönderir. Bu fonksiyon, JSON formatına benzer şekilde çalışır, ancak XML formatını kullanır.
**Kullanımı:**
```go
package main

import (
    "encoding/xml"

    "github.com/coderianx/flint"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

func main() {
    app := flint.NewServer()

    user := User{
        ID:    1,
        Name:  "John Doe",
        Email: "john.doe@example.com",
    }

    app.Get("/user", func(ctx *flint.Context) {
        ctx.XML(200, user)
    })

    app.Run(":8080")
}
```
---
## `ctx.XMLPretty()`
**Açıklama:** `ctx.XMLPretty()` fonksiyonu, belirtilen veri yapısını XML formatında serileştirir ve HTTP yanıtı olarak gönderir. Bu fonksiyon, XML çıktısını daha okunabilir hale getirmek için girintili (pretty) formatta döner.
**Kullanımı:**
```go
package main

import (
    "encoding/xml"

    "github.com/coderianx/flint"
)

type User struct {
    XMLName xml.Name `xml:"user"`
    ID      int      `xml:"id"`
    Name    string   `xml:"name"`
    Email   string   `xml:"email"`
}

func main() {
    app := flint.NewServer()

    user := User{
        ID:    1,
        Name:  "John Doe",
        Email: "john.doe@example.com",
    }

    app.Get("/user", func(ctx *flint.Context) {
        ctx.XMLPretty(200, user)
    })
    app.Run(":8080")
}
```
---
## `ctx.XMLFile()`
**Açıklama:** `ctx.XMLFile()` fonksiyonu, belirtilen dosya yolundaki XML dosyasını okuyarak HTTP yanıtı olarak gönderir. Bu fonksiyon, sunucuda bulunan bir XML dosyasını istemciye göndermek için kullanılır.
**Kullanımı:**
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Get("/file", func(ctx *flint.Context) {
        ctx.XMLFile(200, "path/to/your/file.xml")
    })

    app.Run(":8080")
}
```
---