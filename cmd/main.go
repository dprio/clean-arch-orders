package main

import "github.com/dprio/clean-arch-orders/cmd/app"

func main() {
	app := app.New()

	if err := app.Start(); err != nil {
		println(err.Error())
	}
	println("Acabou !")
}
