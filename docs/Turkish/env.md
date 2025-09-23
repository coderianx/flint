# Env - Flint
Flint ile .env dosyası kullanımı
Flint projelerinde .env dosyası kullanarak çevresel değişkenleri yönetmek oldukça kolaydır. Bu dosya, uygulamanızın yapılandırma ayarlarını ve hassas bilgilerini saklamak için idealdir.
## .env Dosyasını Oluşturma
Projenizin kök dizininde bir `.env` dosyası oluşturun. Bu dosya, anahtar-değer çiftleri şeklinde çevresel değişkenlerinizi içerecektir. Örneğin:
```
API_KEY=your_api_key_here
DATABASE_URL=your_database_url_here
DEBUG_MODE=true
PORT=8080
```
## Çevresel Değişkenleri Yükleme
```go
package main

import (
    "github.com/coderianx/flint"
)

func main() {
    app := flint.NewServer()
    env := flint.EnvLoad(".env") // .env dosyasını yükle

    app.Get("/", func(ctx *flint.Context) {
        api_key := env.GetEnv("API_KEY") // API_KEY değişkenini al
        db_url := env.GetEnv("DATABASE_URL") // DATABASE_URL değişkenini al
        debug_mode := env.GetBool("DEBUG_MODE") // DEBUG_MODE değişkenini al
        port := env.GetInt("PORT") // PORT değişkenini al
        ctx.Stringf("API Key: %s\nDatabase URL: %s\nDebug Mode: %t\nPort: %d", api_key, db_url, debug_mode, port)
    })

    app.Run(":8080")
}
```
## Açıklama:
- `flint.EnvLoad(".env")`: Bu fonksiyon, belirtilen .env dosyasını yükler ve çevresel değişkenleri erişilebilir hale getirir.
- `env.GetEnv("KEY")`: Belirtilen anahtara karşılık gelen değeri döner.
- `env.GetBool("KEY")`: Belirtilen anahtara karşılık gelen değeri boolean olarak döner.
- `env.GetInt("KEY")`: Belirtilen anahtara karşılık gelen değeri integer olarak döner.
## Notlar:
- .env dosyasını versiyon kontrol sistemlerine dahil etmeyin, çünkü bu dosya genellikle hassas bilgileri içerir.
- Çevresel değişkenler, uygulamanızın farklı ortamlar (geliştirme, test, üretim) için yapılandırılmasını kolaylaştırır.
## Daha Fazla Bilgi
