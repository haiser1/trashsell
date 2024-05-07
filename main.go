package main

import (
	"mini-project/config"
	"mini-project/controller"
	"mini-project/driver/mysql"
	"mini-project/repository"
	"mini-project/routes"
	"mini-project/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.InitConfigMysql()
	db := mysql.Connection(config.InitConfigMysql())

	app := echo.New()

	buyerRepo := repository.BuyerRepositoryImpl{DB: db}
	buyerService := service.BuyerServiceImpl{BuyerRepository: buyerRepo, Validate: validator.New()}
	buyerController := controller.NewBuyerController(buyerService)

	routes := routes.RouteController{
		BuyerController: buyerController,
	}
	routes.InitRoute(app)

	app.Logger.Fatal(app.Start(":8080"))

}
