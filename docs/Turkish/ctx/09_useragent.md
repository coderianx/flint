# User Agent

**flint framework** ile User Agent bilgilerini nasıl alabileceğinizi gösteren bir örnek.

## User Agent Bilgisi Alma

User Agent, istemcinin (tarayıcı, uygulama vb.) sunucuya gönderdiği bir başlıktır. Flint ile User Agent bilgisine erişmek için `ctx.UserAgent()` metodunu kullanabilirsiniz.

```go
package main

import (
    "fmt"
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        if ctx.Get() {
            useragent := ctx.UserAgent() // User Agent bilgisini alır
            fmt.Printf("User Agent: %s\n", useragent)
        }
    })
}
```
**Açıklama:**
- `ctx.UserAgent()`: İstemcinin gönderdiği User Agent bilgisini alır.
- `fmt.Printf("User Agent: %s\n", useragent)`: User Agent bilgisini konsola yazdırır.
