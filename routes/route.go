package routes

import (
	buyer "mini-project/controller/buyer"
	seller "mini-project/controller/seller"
	trash_route "mini-project/controller/trash"
	trash_transaction_route "mini-project/controller/trash_transaction"

	"github.com/labstack/echo/v4"

	"mini-project/middleware"
)

type RouteController struct {
	BuyerController            buyer.BuyerController
	SellerController           seller.SellerController
	TrashController            trash_route.TrashController
	TrashTransactionController trash_transaction_route.TrashTransactionController
}

func (route RouteController) InitRoute(app *echo.Echo) {
	route.publicRoute(app)
	route.buyerRoute(app)
	route.sellerRoute(app)
	route.trashRoute(app)
	route.trashTransactionRoute(app)
}

func (route RouteController) publicRoute(app *echo.Echo) {
	public := app.Group("")
	public.POST("/buyers/register", route.BuyerController.RegisterBuyerController)
	public.POST("/buyers/login", route.BuyerController.LoginBuyerController)

	public.POST("/sellers/register", route.SellerController.RegisterSellerController)
	public.POST("/sellers/login", route.SellerController.LoginSellerController)

}

func (route RouteController) buyerRoute(app *echo.Echo) {
	buyer := app.Group("", middleware.JwtMiddlewareBuyer)
	buyer.GET("/buyers", route.BuyerController.GetDataBuyerController) //done
	buyer.PUT("/buyer", route.BuyerController.UpdateBuyerController)   // done
}

func (route RouteController) sellerRoute(app *echo.Echo) {
	seller := app.Group("", middleware.JwtMiddlewareSeller)
	seller.GET("/sellers", route.SellerController.GetDataSellerController)
	seller.PUT("/seller", route.SellerController.UpdateSellerController)
	seller.GET("/sellers/trashes", route.TrashController.GetTrashPagginationController)
}

func (route RouteController) trashTransactionRoute(app *echo.Echo) {
	trashTransactionSeller := app.Group("", middleware.JwtMiddlewareSeller)
	trashTransactionSeller.POST("/sellers/trash-transaction/:id", route.TrashTransactionController.CreateTrashTransactionController)
	trashTransactionSeller.GET("/sellers/trash-transactions", route.TrashTransactionController.GetAllTrashTransactionBySellerId)
	trashTransactionSeller.GET("/sellers/trash-transaction", route.TrashTransactionController.GetTrashTransactionDoneBySellerIdController)

	trashTransactionBuyer := app.Group("", middleware.JwtMiddlewareBuyer)
	trashTransactionBuyer.PUT("/buyers/trash-transaction/:id", route.TrashTransactionController.CreateTrashTransactionDoneController)
	trashTransactionBuyer.GET("/buyers/trash-transactions", route.TrashTransactionController.GetAllTrashTransactionByBuyerId)
	trashTransactionBuyer.GET("/buyers/trash-transaction", route.TrashTransactionController.GetTrashTransactionDoneByBuyerIdController)
}

func (route RouteController) trashRoute(app *echo.Echo) {
	trash := app.Group("", middleware.JwtMiddlewareBuyer)
	trash.POST("/buyers/type-trash", route.TrashController.CreateTypeTrashController)
	trash.GET("/buyers/types-trash", route.TrashController.GetAllTypeTrashController)
	trash.PUT("/buyers/types-trash/:id", route.TrashController.UpdateTypeTrashController)

	trash.POST("/buyers/trash", route.TrashController.CreateTrashController)
	trash.GET("/buyers/trashes", route.TrashController.GetAllTrashByBuyerIdController)
	trash.GET("/buyers/trash/:id", route.TrashController.GetTrashByIdController)
	trash.PUT("/buyers/trashes/:id", route.TrashController.UpdateTrashController)
	trash.DELETE("/buyer/trash/:id", route.TrashController.DeleteTrashController)
}
