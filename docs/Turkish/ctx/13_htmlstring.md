# Ctx12 | `ctx.HTMLString()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/html", func(ctx *flint.Context) {
        htmlContent := `<h1>Merhaba, Flint!</h1><p>Bu, ctx.HTMLString() kullanılarak oluşturulmuş bir HTML sayfasıdır.</p>`
        ctx.HTMLString(200, htmlContent)
    })

    app.Run(":8080")
}
```
### Açıklama:
Bu örnekte, `/html` yoluna yapılan isteklerde `ctx.HTMLString()` fonksiyonu kullanılarak dinamik olarak oluşturulmuş bir HTML içeriği döndürülür. `htmlContent` değişkeni içinde HTML kodu bir string olarak tanımlanır ve `ctx.HTMLString(200, htmlContent)` ile istemciye gönderilir. Bu yöntem, basit HTML içeriklerini hızlıca sunmak için kullanışlıdır.
