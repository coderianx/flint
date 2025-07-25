# 3 - Handle

```go
package main

import (  
    "github.com/grayvort3x/Flint"  
)  

func main() {
    app := flint.NewServer()

    // Handle GET request to "/"
    app.Handle("/", func(ctx *flint.Context) {
        // Codes
    })

```
