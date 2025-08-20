package main

import (
	"fmt"

	"github.com/coderianx/flint"
)

func main() {
	app := flint.NewServer()

	app.Handle("/", func(ctx *flint.Context) {
		useragent := ctx.UserAgent()
		fmt.Println("UserAgent: ", useragent)
		ctx.Stringf(200, "Your Useragent: %s", useragent)
	})

	app.Run()
}
