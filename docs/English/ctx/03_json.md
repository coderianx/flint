# Ctx/03 | `ctx.JSON()`

```go
package main

import (
    "github.com/coderianx/flint"
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

---

## `ctx.JSONPretty()`

### If you want to return a readable JSON response, you can use the `ctx.JSONPretty()` function:

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        data := map[string]string{
            "message": "Hello, World!",
            "status":  "200",
        }
        ctx.JSONPretty(200, data)
        // or
        ctx.JSONPretty(200, data, 4) // uses 4 spaces for indentation
    })
}
```

### 🧠 This code creates a Flint server, registers a handler for the `/` route, and responds with a readable JSON format using `ctx.JSONPretty()`.

---
## `JSONFile()`

### If you want to send a JSON file as a response, you can use the `ctx.JSONFile()` function:
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        ctx.JSONFile(200, "data.json")
    })

    app.Run(":8080")
}
```
### 🧠 This code creates a Flint server, registers a handler for the `/` route, and responds with the contents of `data.json` file using `ctx.JSONFile()`.
