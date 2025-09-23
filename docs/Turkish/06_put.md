# 6 - Put

```go
package main

import (  
    "github.com/coderianx/flint"  
)

func main() {
    app := flint.NewServer()

    app.Put("/", func(ctx *flint.Context) {
        // ctx kodları
    })

    app.Run(":8000")
}
```
### Bu örnekte, yeni bir sunucu oluşturuyoruz ve kök yolu ("/") için bir PUT isteği rotası tanımlıyoruz. `Put()` yöntemi, belirtilen yolla ilişkilendirilmiş bir işlevi tanımlamak için kullanılır. Bu yola bir PUT isteği yapıldığında, işlev çalıştırılacaktır.
### Flint, HTTP PUT isteklerini işlemek için `Put()` yöntemini sağlar. Bu yöntem, belirli bir yola gelen PUT isteklerini yakalamak ve işlemek için kullanılır. Flint, bu tür istekleri yönetmek için basit ve etkili bir yol sunar.
