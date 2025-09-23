# 6 - Put

```go
package main

import (  
    "github.com/coderianx/flint"  
)

func main() {
    app := flint.NewServer()

    app.Put("/", func(ctx *flint.Context) {
        // ctx codes
    })

    app.Run(":8000")
}
```
### In this example, we create a new server and define a PUT request route for the root path ("/"). The `Put()` method is used to define a function associated with the specified path. When a PUT request is made to this path, the function will be executed.
---
### Flint provides the `Put()` method to handle HTTP PUT requests. This method is used to capture and process PUT requests coming to a specific path. Flint offers a simple and effective way to manage such requests.