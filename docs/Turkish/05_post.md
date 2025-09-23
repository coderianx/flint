# 5 - Post

```go
package main

import (  
    "github.com/coderianx/flint"  
)

func main() {
    app := flint.NewServer()

    app.Post("/", func(ctx *flint.Context) {
        // ctx kodları
    })

    app.Run(":8000")
}
```
### Bu örnekte, yeni bir sunucu oluşturuyoruz ve kök yolu ("/") için bir POST isteği rotası tanımlıyoruz. `Post()` yöntemi, belirtilen yolla ilişkilendirilmiş bir işlevi tanımlamak için kullanılır. Bu yola bir POST isteği yapıldığında, işlev çalıştırılacaktır.
### Flint, HTTP POST isteklerini işlemek için `Post()` yöntemini sağlar. Bu yöntem, belirli bir yola gelen POST isteklerini yakalamak ve işlemek için kullanılır. Flint, bu tür istekleri yönetmek için basit ve etkili bir yol sunar.
