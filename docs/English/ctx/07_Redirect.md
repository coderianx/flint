# Ctx/07 | `ctx.Redirect()`

### Redirect() is used for redirection operations.

### Example usage:

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
        ctx.Send("You have been redirected from the root path.")
    })

    app.Listen(":8080")
}
```
### Explanation:
`ctx.Redirect()` function is used to perform HTTP redirection. This function redirects the client to the specified URL. In the example, requests to the `/` path are redirected to `/new-path`, and a message is sent for requests to this new path.
### Notes:
- The `ctx.Redirect()` function uses the HTTP status code 302 (Found) by default. If you want to use a different status code, you can use the `ctx.RedirectWithStatus()` function.
- The redirection operation allows the client to be redirected to a new URL. This is often used to improve user experience or change the URL structure.
