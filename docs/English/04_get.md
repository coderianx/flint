# 4 - Get

```go
package main

import (  
    "github.com/coderianx/flint"  
)

func main() {
    app := flint.NewServer()

    app.Get("/", func(ctx *flint.Context) {
        // ctx codes
    })

    app.Run(":8000")
}
```
### In this example, we create a new server and define a GET request route for the root path ("/"). The `Get()` method is used to define a function associated with the specified path. When a GET request is made to this path, the function will be executed.
---
### Flint provides the `Get()` method to handle HTTP GET requests. This method is used to capture and process GET requests coming to a specific path. Flint offers a simple and effective way to manage such requests.