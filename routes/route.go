package routes

import (
	buyer "mini-project/controller/buyer"
	seller "mini-project/controller/seller"
	trash_route "mini-project/controller/trash"

	"github.com/labstack/echo/v4"

	"mini-project/middleware"
)

type RouteController struct {
	BuyerController  buyer.BuyerController
	SellerController seller.SellerController
	TrashController  trash_route.TrashController
}

func (route RouteController) InitRoute(app *echo.Echo) {
	route.publicRoute(app)
	route.buyerRoute(app)
	route.sellerRoute(app)
	route.trashRoute(app)
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

func (route RouteController) trashRoute(app *echo.Echo) {
	trash := app.Group("", middleware.JwtMiddlewareBuyer)
	trash.POST("/buyers/type-trash", route.TrashController.CreateTypeTrashController)     //done
	trash.POST("/buyers/trash", route.TrashController.CreateTrashController)              //done
	trash.GET("/buyers/types-trash", route.TrashController.GetAllTypeTrashController)     //done
	trash.GET("/buyers/trashes", route.TrashController.GetAllTrashByBuyerIdController)    //done
	trash.GET("/buyers/trash/:id", route.TrashController.GetTrashByIdController)          //done
	trash.PUT("/buyers/trashes/:id", route.TrashController.UpdateTrashController)         //done
	trash.PUT("/buyers/types-trash/:id", route.TrashController.UpdateTypeTrashController) //done
	trash.DELETE("/buyer/trash/:id", route.TrashController.DeleteTrashController)         //done
}
