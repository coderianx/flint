package main

import "github.com/coderianx/flint"

func main() {
	app := flint.NewServer()

	app.Handle("/", func(ctx *flint.Context) {
		ctx.String(200, "Hello Flint Familyyy")
	})

	app.Run(":8080")
}
