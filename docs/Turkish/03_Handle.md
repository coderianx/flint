# 3 - Handle

```go
package main

import (  
    "github.com/grayvort3x/Flint"  
)

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx *flint.Context) {
        // ctx kodları
    })

    app.Run(":8000")
}
```
### Bu örnekte, yeni bir sunucu oluşturuyoruz ve kök yolu ("/") için bir rota tanımlıyoruz. `Handle()` yöntemi, belirtilen yolla ilişkilendirilmiş bir işlevi tanımlamak için kullanılır. Bu yola bir istek yapıldığında, işlev çalıştırılacaktır.