# Ctx/03 | `ctx.JSON()`

```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        data := map[string]string{
            "message": "Hello, World!",
        }
        ctx.JSON(200, data)
    })

    app.Run(":8080")
}
```
### 🧠 This code creates a Flint server, registers a handler for the `/` route, and responds with JSON format using `ctx.JSON()`.