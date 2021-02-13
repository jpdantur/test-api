package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jpdantur/test-api/internal/http/ping"
)

type App struct {
	Router *gin.Engine
}

func New() *App {
	app := &App {
		Router: gin.Default(),
	}
	app.setupRoutes()
	return app
}

func (app *App) Start(port string) {
	if err := app.Router.Run(":" + port); err != nil {
		panic("Error running server")
	}
}

func (app *App) setupRoutes() {
	pingController := ping.NewController()

	app.Router.GET("/ping", pingController.HandlePing)
}