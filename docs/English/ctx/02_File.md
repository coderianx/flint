# Ctx/02 | `ctx.File()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    // Handle GET request to "/"
    app.Handle("/", func(ctx *flint.Context) {
        // Sends default 200 status with file
        ctx.File("index.html")

        // Or specify status code manually
        ctx.File(200, "index.html")
    })

    app.Run(":8080")
}
```
### 🧠 This code creates a Flint server, registers a route for /, and responds with a file using ctx.File().
