package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jpdantur/test-api/internal/domain/transactions"
	transactionsController "github.com/jpdantur/test-api/internal/http/controllers/transactions"
)

type App struct {
	Router              *gin.Engine
	transactionsService transactions.Service
}

func New() *App {
	transactionsService := transactions.NewService()
	app := &App{
		Router:              gin.Default(),
		transactionsService: transactionsService,
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
	transactionsController := transactionsController.NewController(app.transactionsService)

	transactions := app.Router.RouterGroup.Group("/transactions")
	{
		transactions.POST("", transactionsController.HandleAdd)
		transactions.GET("/:id", transactionsController.HandleGetByID)
		transactions.GET("", transactionsController.HandleGetHistory)
	}
	app.Router.GET("/balance", transactionsController.HandleGetBalance)
}
