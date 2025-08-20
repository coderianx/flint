# 🔹 Context 05 – `ctx.FormData()`

The `ctx.FormData()` method is used to retrieve form data submitted through `POST` requests in the Flint Framework.

---

## 🧠 Use Case

Use `ctx.FormData()` when handling HTML form submissions via `POST` requests.  
It’s useful for login forms, registration, feedback forms, etc.

---

## ✅ Basic Usage

```go
app.Handle("/login", func(ctx *flint.Context) {
    if ctx.Post() {
        username := ctx.FormData("username")
        password := ctx.FormData("password")

        // Example response
        ctx.String("Welcome")
    }
})
```

---

## 💡 Advanced Example: Form Handling

```go
package main

import (
    "github.com/grayvort3x/Flint"
    "fmt"
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        if ctx.Get() {
            ctx.File("index.html")
            return
        }
        if ctx.Post() {
            username := ctx.FormData("username")
            password := ctx.FormData("password")

            ctx.String("Welcome")
        }
    })
}
```

---

## 🌐 Example `index.html`

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Flint Form Example</title>
</head>
<body>
  <h2>Login with Flint</h2>
  <form action="/" method="POST">
    <label>Username:</label><br>
    <input type="text" name="username" required><br><br>

    <label>Password:</label><br>
    <input type="password" name="password" required><br><br>

    <button type="submit">Submit</button>
  </form>
</body>
</html>
```

---

## 🔍 Related Methods

| Method           | Description                          |
|------------------|--------------------------------------|
| `ctx.Get()`      | Checks if the request method is GET  |
| `ctx.Post()`     | Checks if the request method is POST |
| `ctx.File(path)` | Serves a static file (e.g., HTML)    |
| `ctx.String(str)`| Sends plain text response            |

---