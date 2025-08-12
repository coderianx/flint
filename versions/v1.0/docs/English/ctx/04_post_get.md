# 🔹 Context 04 – `ctx.Get()` and `ctx.Post()`

These methods are used to check the HTTP method of an incoming request in the Flint Framework.

---


- `ctx.Get()` returns `true` if the HTTP method is `GET`.
- `ctx.Post()` returns `true` if the HTTP method is `POST`.

---

## 🧠 Use Case

These methods are useful for handling routes that support multiple request methods, such as GET for rendering pages and POST for processing form submissions.

---

## ✅ Example: `ctx.Get()`

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
            // Your GET handling code here
        }
    })

    app.Run(":8080")
}
```

### 🧠 Explanation:
This code creates a Flint server, registers a route for `/`, and checks if the request method is GET using `ctx.Get()`.

---

## ✅ Example: `ctx.Post()`

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
            // Your POST handling code here
        }
    })

    app.Run(":8080")
}
```

### 🧠 Explanation:
This code checks if the request method is POST using `ctx.Post()`.

---

## 💡 Combined Example: GET + POST

```go
package main

import (
    "github.com/grayvort3x/Flint"
)

func main() {
    app := flint.NewServer()

    app.Handle("/submit", func(ctx *flint.Context) {
        if ctx.Get() {
            ctx.File("index.html") // Serve HTML page
        }
        if ctx.Post() {
            // Handle form submission
        }
    })

    app.Run(":8080")
}
```

---

## 🔍 Related Methods

| Method           | Description                          |
|------------------|--------------------------------------|
| `ctx.Get()`      | Checks if the request method is GET  |
| `ctx.Post()`     | Checks if the request method is POST |
| `ctx.File(path)` | Serves static files (e.g. HTML)      |
| `ctx.String()`   | Sends a plain text response          |

---

## ⚠️ Notes

- These methods help organize your route handlers by HTTP method.
- Great for routes like `/login`, `/register`, `/submit`, etc.

---
