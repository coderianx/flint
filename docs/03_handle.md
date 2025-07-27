# 3 - Handle

```go
package main

import (  
    "github.com/grayvort3x/Flint"  
)  

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        // Codes
    })

    app.Run(":8000")
}
```