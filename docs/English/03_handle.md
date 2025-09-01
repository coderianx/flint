# 3 - Handle

```go
package main

import (  
    "github.com/coderianx/flint"  
)  

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        // ctx codes
    })

    app.Run(":8000")
}
```

### In this example, we create a new server and define a route for the root path ("/"). The `Handle` method is used to associate a function with the specified path. When a request is made to this path, the function will be executed.