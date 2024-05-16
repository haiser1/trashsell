package controller

import "github.com/labstack/echo/v4"

type BuyerController interface {
	RegisterBuyerController(c echo.Context) error
	LoginBuyerController(c echo.Context) error
	GetDataBuyerController(c echo.Context) error
	UpdateBuyerController(c echo.Context) error
}
