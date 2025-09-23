# Ctx/11 | `ctx.Param()` & `ctx.ParamInt()` & `ctx.ParamIntDefault()`

The `ctx.Param()` function is used to retrieve a parameter from the URL path in a web request. It takes a single argument, which is the name of the parameter you want to retrieve, and returns its value as a string.

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app.Handle("/@:username", func(ctx *flint.Context)) {
        username := ctx.Param("username")
        ctx.String(200, "Username: @%s", username)
    }

    app.Run()
}
```
In this example, when a request is made to a URL like `/@john`, the `ctx.Param("username")` call will return the string `"john"`.

---

`ctx.ParamInt()` & `ctx.ParamIntDefault()`

The `ctx.ParamInt()` function is similar to `ctx.Param()`, but it retrieves the parameter as an integer. It takes the same argument (the name of the parameter) and returns its value as an `int`. If the parameter cannot be converted to an integer, it returns `0`.

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/user/:id", func(ctx *flint.Context)) {
        id := ctx.ParamInt("id")
        ctx.String(200, "User ID: %d", id)
    }

    app.Run()
}
```
In this example, when a request is made to a URL like `/user/42`, the `ctx.ParamInt("id")` call will return the integer `42`.

The `ctx.ParamIntDefault()` function is similar to `ctx.ParamInt()`, but it allows you to specify a default value to return if the parameter cannot be converted to an integer. It takes two arguments: the name of the parameter and the default value.

```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/user/:id", func(ctx *flint.Context)) {
        id := ctx.ParamIntDefault("id", 1)
        ctx.String(200, "User ID: %d", id)
    }

    app.Run()
}
```
In this example, when a request is made to a URL like `/user/abc`, the `ctx.ParamIntDefault("id", 1)` call will return the default value `1` since `"abc"` cannot be converted to an integer.