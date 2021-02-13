package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jpdantur/test-api/internal/app"
)

func main() {
	app := app.New()

	app.Start(os.Getenv("PORT"))
}
