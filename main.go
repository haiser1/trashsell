package main

import (
	"mini-project/config"
	buyer_controller "mini-project/controller/buyer"
	seller_controller "mini-project/controller/seller"
	trash_controller "mini-project/controller/trash"
	"mini-project/driver/mysql"
	buyer_repo "mini-project/repository/buyer_repo"
	seller_repo "mini-project/repository/seller_repo"
	trash_repo "mini-project/repository/trash_repo"
	"mini-project/routes"
	buyer_service "mini-project/service/buyer_service"
	seller_service "mini-project/service/seller_service"
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

	sellerRepository := seller_repo.NewSellerRepository(db)
	sellerService := seller_service.NewSellerService(sellerRepository, validator.New())

	sellerController := seller_controller.NewSellerController(sellerService)

	trashRepository := trash_repo.NewTrashRepository(db)
	trashService := trash_service.NewTrashService(trashRepository, validator.New())
	trashController := trash_controller.NewTrashController(trashService)

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes := routes.RouteController{
		BuyerController:  buyerController,
		TrashController:  trashController,
		SellerController: sellerController,
	}

	routes.InitRoute(app)
	app.Logger.Fatal(app.Start(":8000"))

}
