# Ctx12 | `ctx.XML()` & `ctx.XMLPretty()` & `ctx.XMLFile()`

## `ctx.XML()`
**Description:** The `ctx.XML()` function serializes the specified data structure into XML format and sends it as an HTTP response. This function works similarly to the JSON format but uses XML format instead.
**Usage:**
```go
package main

import (
    "encoding/xml"

    "github.com/coderianx/flint"
)

type User struct {
    XMLName xml.Name `xml:"user"`
    ID      int      `xml:"id"`
    Name    string   `xml:"name"`
    Email   string   `xml:"email"`
}

func main() {
    app := flint.NewServer()

    app.Get("/users", func(ctx *flint.Context) {
        user := User{
            ID:      1,
            Name:    "John Doe",
            Email:   "john.doe@example.com",
        }

        ctx.XML(200, user)
    })

    app.Run(":8080")
}
```
## `ctx.XMLPretty()`
**Description:** The `ctx.XMLPretty()` function serializes the specified data structure into pretty-printed XML format and sends it as an HTTP response. This function works similarly to the JSON format but uses XML format instead.
**Usage:**
```go
package main

import (
    "encoding/xml"

    "github.com/coderianx/flint"
)

type User struct {
    XMLName xml.Name `xml:"user"`
    ID      int      `xml:"id"`
    Name    string   `xml:"name"`
    Email   string   `xml:"email"`
}

func main() {
    app := flint.NewServer()

    app.Get("/users", func(ctx *flint.Context) {
        user := User{
            ID:      1,
            Name:    "John Doe",
            Email:   "john.doe@example.com",
        }

        ctx.XMLPretty(200, user)
    })
    app.Run(":8080")
}
```
## `ctx.XMLFile()`
**Description:** The `ctx.XMLFile()` function reads an XML file from the specified file path and sends its content as an HTTP response with the appropriate XML content type.
**Usage:**
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()

    app.Get("/file", func(ctx *flint.Context) {
        ctx.XMLFile(200, "data.xml")
    })

    app.Run(":8080")
}
```
