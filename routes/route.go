package routes

import (
	"mini-project/controller"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	BuyerController controller.BuyerController
}

func (route RouteController) InitRoute(app *echo.Echo) {
	app.POST("/register", route.BuyerController.RegisterBuyerController)

}
