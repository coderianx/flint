# 2 - Run Server - Flint

```go
package main  
  
import (  
    flint "github.com/coderianx/flint"  
)  
func main() {  
    app := flint.NewServer()  
  
    app.Run(":8080")  
}
```

### This code creates a new Flint server and runs it on port 8080

**You can open a server on different ports by changing the port number in run().**

### Example;

```go
package main 

import (  
    "github.com/coderianx/flint"  
)  
func main() {  
    app := flint.NewServer()  
  
    app.Run(":3000")  
}
```

### This code creates a new Flint server and runs it on port 3000

### Note:
If you leave `Run()` blank, Flint will use port 8080 by default.
### Example;

```go
package main

import (  
    "github.com/coderianx/flint"  
)

func main() {  
    app := flint.NewServer()  
  
    app.Run() // Uses port 8080 by default
}
```
Flint uses port 8080 by default. If you do not want to use this port, you can provide any port number to the `Run()` function.