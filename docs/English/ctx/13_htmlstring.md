# Ctx13 | `ctx.HTMLString()`

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/html", func(ctx *flint.Context) {
        htmlContent := `<h1>Hello, Flint!</h1><p>This is an HTML page generated using ctx.HTMLString().</p>`
        ctx.HTMLString(200, htmlContent)
    })

    app.Run(":8080")
}
```
### Explanation:
In this example, when a request is made to the `/html` route, the `ctx.HTMLString()` function is used to return dynamically generated HTML content. The HTML code is defined as a string in the `htmlContent` variable and sent to the client using `ctx.HTMLString(200, htmlContent)`. This method is useful for quickly serving simple HTML content.
