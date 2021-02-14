package main

import (
	"github.com/jpdantur/test-api/internal/app"
)

func main() {
	app := app.New()

	app.Start("8080")
}
