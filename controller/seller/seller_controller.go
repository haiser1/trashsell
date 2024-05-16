package controller

import "github.com/labstack/echo/v4"

type SellerController interface {
	RegisterSellerController(c echo.Context) error
	LoginSellerController(c echo.Context) error
	GetDataSellerController(c echo.Context) error
	UpdateSellerController(c echo.Context) error
}
