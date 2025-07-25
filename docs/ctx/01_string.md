# Ctx/01 | `ctx.String()`

```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    // Handle GET request to "/"
    app.Handle("/", func(ctx *flint.Context) {
        // Sends default 200 status with string
        ctx.String("Hello, Flint!")

        // Or specify status code manually
        ctx.String(200, "Hello Flint!")
    })

    app.Run(":8080")
}

```

🧠 This code creates a Flint server, registers a route for /, and responds with plain text using ctx.String().