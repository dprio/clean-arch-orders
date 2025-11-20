package main

import "github.com/dprio/clean-arch-orders/cmd/app"

func main() {
	app := app.New()

	app.Start()
	println("Acabou !")
}
