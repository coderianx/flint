# 5 - Post

```go
package main

import (  
    "github.com/coderianx/flint"  
)

func main() {
    app := flint.NewServer()

    app.Post("/", func(ctx *flint.Context) {
        // ctx codes
    })

    app.Run(":8000")
}
```
### In this example, we create a new server and define a POST request route for the root path ("/"). The `Post()` method is used to define a function associated with the specified path. When a POST request is made to this path, the function will be executed.
---
### Flint provides the `Post()` method to handle HTTP POST requests. This method is used to capture and process POST requests coming to a specific path. Flint offers a simple and effective way to manage such requests.