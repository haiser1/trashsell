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

	trash_transaction_controller "mini-project/controller/trash_transaction"
	trash_transaction_repo "mini-project/repository/trash_transaction"
	trash_transaction_service "mini-project/service/trash_transaction"

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

	trashTransactionRepository := trash_transaction_repo.NewTrashTransactionRepository(db)
	trashTransactionService := trash_transaction_service.NewTrashTransactionService(trashTransactionRepository, validator.New())
	trashTransactionController := trash_transaction_controller.NewTrashTransactionController(trashTransactionService)

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes := routes.RouteController{
		BuyerController:            buyerController,
		TrashController:            trashController,
		SellerController:           sellerController,
		TrashTransactionController: trashTransactionController,
	}

	routes.InitRoute(app)
	app.Logger.Fatal(app.Start(":8000"))

}
