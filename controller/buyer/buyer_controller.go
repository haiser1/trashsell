package controller

import "github.com/labstack/echo/v4"

type BuyerController interface {
	RegisterBuyerController(c echo.Context) (err error)
	LoginBuyerController(c echo.Context) (err error)
	GetDataBuyerController(c echo.Context) (err error)
	UpdateBuyerController(c echo.Context) (err error)
}
