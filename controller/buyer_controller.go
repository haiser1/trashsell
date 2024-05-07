package controller

import "github.com/labstack/echo/v4"

type BuyerController interface {
	// contrroler use echo
	RegisterBuyerController(c echo.Context) (err error)
}
