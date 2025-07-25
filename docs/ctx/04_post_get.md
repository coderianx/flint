# Ctx/04 | `ctx.Get()` and `ctx.Post()`

## `ctx.Get()`
```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        // Handle GET request
        if ctx.Get() {
            //codes
        }
    })

    app.Run(":8080")
}
```
### 🧠 This code creates a Flint server, registers a route for /, and checks if the request is a GET request using ctx.Get().

## `ctx.Post()`
```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        // Handle POST request
        if ctx.Post() {
            //codes
        }
    })

    app.Run(":8080")
}
```
### 🧠 This code creates a Flint server, registers a route for /, and checks if the request is a POST request using ctx.Post().

## Examples;
```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/submit", func(ctx *flint.Context) {
        if ctx.Get() {
            ctx.File("indx.html") // Serve a file for GET requests
        }
        if ctx.Post() {
            // codes 
        }
    })

    app.Run(":8080")
}
```