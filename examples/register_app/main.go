package main

import (
	"fmt"

	"github.com/coderianx/flint"
)

func main() {
	app := flint.NewServer()

	app.Handle("/", func(ctx *flint.Context) {
		if ctx.Get() {
			ctx.File("./examples/register_app/index.html")
			return
		}
		if ctx.Post() {
			username := ctx.FormData("username")
			password := ctx.FormData("password")
			ctx.Stringf(200, "Hoşgeldin %s", username)
			fmt.Println(username, ":", password)
		}
	})

	app.Run(":3000")
}
