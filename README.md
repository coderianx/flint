<p align="center">
  <img src="assets/flint.jpg" alt="Flint Logo" width="200"/>
</p>

<h1 align="center">Flint</h1>
<p align="center">A blazing fast and lightweight web framework for Go.</p>

[![Docs](https://img.shields.io/badge/Web_site-Flint-%23098687?style=for-the-badge&logo=firefox-browser&logoColor=white)](https://flintgo.netlify.app)

[![Telegram Channel](https://img.shields.io/badge/Telegram-Channel-blue?logo=telegram)](https://t.me/flint_framework)

[![Telegram Turkey Channel](https://img.shields.io/badge/Telegram-Channel_Türkiye-blue?logo=telegram)](https://t.me/flint_framework_tr)

[![X](https://img.shields.io/badge/X-Account-black?logo=twitter)](https://x.com/flintframework)

[![YouTube](https://img.shields.io/badge/YouTube-Channel-red?logo=youtube)](https://youtube.com/@flint_framework)


---

## ⚡️ What is Flint?

Flint is a simple and minimal backend framework written in Go, made for fast API development with clean structure and easy usage. 

## 🐾 Flinex
<p align="center">
  <img src="assets/flinex.png" alt="Flint Mascot" width="150"/>
  <br>
</p>

## 🚀 Quick Start

```bash
go get github.com/coderianx/flint
```

```go
package main

import "github.com/coderianx/flint"

func main() {
    app := flint.NewServer()

    app.Handle("/", func(ctx flint.Context) {
        ctx.String("Hello from Flint!")
    })

    app.Run(":8080")
}
```
---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

> Built with ❤️ using Go
