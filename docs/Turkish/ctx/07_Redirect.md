# Ctx/07 | `ctx.Redirect()`

### Redirect yönlendirme işlemleri için kullanılır.

### Örnek kullanım:

```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.New()

    app.Get("/", func(ctx *flint.Context) {
        ctx.Redirect("/new-path")
    })

    app.Get("/new-path", func(ctx *flint.Context) {
        ctx.Send("/ adresinden yönlendirildiniz.")
    })

    app.Listen(":8080")
}
```
### Açıklama:
`ctx.Redirect()` fonksiyonu, HTTP yönlendirmesi yapmak için kullanılır. Bu fonksiyon, istemciyi belirtilen URL'ye yönlendirir. Örnekte, `/` adresine gelen istekler `/new-path` adresine yönlendirilir ve bu adrese gelen istekler için bir mesaj gönderilir.
### Notlar:
- `ctx.Redirect()` fonksiyonu, HTTP durum kodu olarak 302 (Found) kullanır. Eğer farklı bir durum kodu kullanmak isterseniz, `ctx.RedirectWithStatus()` fonksiyonunu kullanabilirsiniz.
- Yönlendirme işlemi, istemcinin yeni bir URL'ye yönlendirilmesini sağlar. Bu, genellikle kullanıcı deneyimini iyileştirmek veya URL yapısını değiştirmek için kullanılır.

