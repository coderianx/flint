# SetNotFound - Flint

### How to display .html page on 404 errors

## Go code
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.SetNot(func(ctx *flint.Context) {
        ctx.Template("404.html")
    })
}
```
### This code will display the 404.html page when a 404 error occurs.

### Using the original Flint 404 error page
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.SetNot(func(ctx *flint.Context) {
        ctx.Default404()
    })
}
```

### This code will display the original Flint 404 error page.