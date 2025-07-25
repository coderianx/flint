
# 2 - Run Server - Flint

```go
package main  
  
import (  
    "github.com/grayvort3x/Flint"  
)  
func main() {  
    app := flint.NewServer()  
  
    app.Run(":8080")  
}
```

### This code creates a new Flint server and runs it on port 8080


