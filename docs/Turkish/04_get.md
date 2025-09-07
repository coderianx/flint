# 4 - Get

```go
package main

import (  
    "github.com/coderianx/flint"  
)

func main() {
    app := flint.NewServer()

    app.Get("/", func(ctx *flint.Context) {
        // ctx kodları
    })

    app.Run(":8000")
}
```
### Bu örnekte, yeni bir sunucu oluşturuyoruz ve kök yolu ("/") için bir GET isteği rotası tanımlıyoruz. `Get()` yöntemi, belirtilen yolla ilişkilendirilmiş bir işlevi tanımlamak için kullanılır. Bu yola bir GET isteği yapıldığında, işlev çalıştırılacaktır.
### Flint, HTTP GET isteklerini işlemek için `Get()` yöntemini sağlar. Bu yöntem, belirli bir yola gelen GET isteklerini yakalamak ve işlemek için kullanılır. Flint, bu tür istekleri yönetmek için basit ve etkili bir yol sunar.