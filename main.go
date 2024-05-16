package main

import (
	"mini-project/config"
	buyer_controller "mini-project/controller/buyer"
	trash_controller "mini-project/controller/trash"
	"mini-project/driver/mysql"
	buyer_repo "mini-project/repository/buyer_repo"
	trash_repo "mini-project/repository/trash_repo"
	"mini-project/routes"
	buyer_service "mini-project/service/buyer_service"
	trash_service "mini-project/service/trash_service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadEnv()
	config.InitConfigMysql()
	db := mysql.Connection(config.InitConfigMysql())

	app := echo.New()

	buyerRepository := buyer_repo.NewBuyerRepository(db)
	buyerService := buyer_service.NewBuyerService(buyerRepository, validator.New())
	buyerController := buyer_controller.NewBuyerController(buyerService)

	trashRepository := trash_repo.NewTrashRepository(db)
	trashService := trash_service.NewTrashService(trashRepository, validator.New())
	trashController := trash_controller.NewTrashController(trashService)

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes := routes.RouteController{
		BuyerController: buyerController,
		TrashController: trashController,
	}

	routes.InitRoute(app)
	app.Logger.Fatal(app.Start(":8000"))

}
