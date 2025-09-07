# Ctx/01 | `ctx.String()`

```go
package main

import (
    "github.com/coderianx/flint"s
)

func main() {
    app := flint.NewServer()

    // Handle GET request to "/"
    app.Handle("/", func(ctx *flint.Context) {
        // Sends default 200 status with 
        ctx.String(200, "Hello Flint!")
    })

    app.Run(":8080")
}

```

🧠 This code creates a Flint server, registers a route for /, and responds with plain text using ctx.String().

## `ctx.Stringf()`
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    // Handle GET request to "/"
    app.Handle("/", func(ctx *flint.Context) {
        // Sends default 200 status with formatted string
        name := "Flint"
        ctx.Stringf(200, "Hello %s!", name)
    })

    app.Run(":8080")
}
```

🧠 This code creates a Flint server, registers a route for /, and responds with a formatted string using ctx.Stringf().

