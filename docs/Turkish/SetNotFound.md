# SetNotFound - Flint

### .html sayfasını 404 hatalarında nasıl görüntülenir

## Go kodu
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.SetNotFound(func(ctx *flint.Context) {
        ctx.Template404("404.html")
    })

    app.Run(":8080")
}
```
### Bu kod, 404 hatası oluştuğunda 404.html sayfasını görüntüler.

### Orijinal Flint 404 hata sayfasını kullanma
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.SetNotFound(func(ctx *flint.Context) {
        ctx.Default404()
    })

    app.Run(":8080")
}
```
### Bu kod, orijinal Flint 404 hata sayfasını görüntüler.


