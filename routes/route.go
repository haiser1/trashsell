package routes

import (
	buyer_route "mini-project/controller/buyer"
	trash_route "mini-project/controller/trash"

	"github.com/labstack/echo/v4"

	"mini-project/middleware"
)

type RouteController struct {
	BuyerController buyer_route.BuyerController
	TrashController trash_route.TrashController
}

func (route RouteController) InitRoute(app *echo.Echo) {
	route.publicRoute(app)
	route.buyerRoute(app)
	route.trashRoute(app)
}

func (route RouteController) publicRoute(app *echo.Echo) {
	public := app.Group("")
	public.POST("/buyers/register", route.BuyerController.RegisterBuyerController)
	public.POST("/buyers/login", route.BuyerController.LoginBuyerController)

}

func (route RouteController) buyerRoute(app *echo.Echo) {
	buyer := app.Group("", middleware.JwtMiddlewareBuyer)
	buyer.GET("/buyers", route.BuyerController.GetDataBuyerController) //done
	buyer.PUT("/buyers", route.BuyerController.UpdateBuyerController)  // done
}

func (route RouteController) trashRoute(app *echo.Echo) {
	trash := app.Group("", middleware.JwtMiddlewareBuyer)
	trash.POST("/buyers/types-trash", route.TrashController.CreateTypeTrashController)    //done
	trash.POST("/buyers/trash", route.TrashController.CreateTrashController)              //done
	trash.GET("/buyers/types-trash", route.TrashController.GetAllTypeTrashController)     //done
	trash.GET("/buyers/trash", route.TrashController.GetAllTrashByBuyerIdController)      //done
	trash.GET("/buyers/trash/:id", route.TrashController.GetTrashByIdController)          //done
	trash.PUT("/buyers/trash/:id", route.TrashController.UpdateTrashController)           //done
	trash.PUT("/buyers/types-trash/:id", route.TrashController.UpdateTypeTrashController) //done
	trash.DELETE("/buyers/trash/:id", route.TrashController.DeleteTrashController)        //done
}
