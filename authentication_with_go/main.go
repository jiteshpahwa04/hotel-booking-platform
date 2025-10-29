package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	config.Load()

	cfg := app.NewConfig()
	app := app.NewApplication(cfg)

	app.Run()
}