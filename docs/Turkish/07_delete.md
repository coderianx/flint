# 7 - Delete

```go
package main

import (  
    "github.com/coderianx/flint"  
)

func main() {
    app := flint.NewServer()

    app.Delete("/", func(ctx *flint.Context) {
        // ctx kodları
    })

    app.Run(":8000")
}
```
### Bu örnekte, yeni bir sunucu oluşturuyoruz ve kök yolu ("/") için bir DELETE isteği rotası tanımlıyoruz. `Delete()` yöntemi, belirtilen yolla ilişkilendirilmiş bir işlevi tanımlamak için kullanılır. Bu yola bir DELETE isteği yapıldığında, işlev çalıştırılacaktır.
---
### Flint, HTTP DELETE isteklerini işlemek için `Delete()` yöntemini sağlar. Bu yöntem, belirli bir yola gelen DELETE isteklerini yakalamak ve işlemek için kullanılır. Flint, bu tür istekleri yönetmek için basit ve etkili bir yol sunar.