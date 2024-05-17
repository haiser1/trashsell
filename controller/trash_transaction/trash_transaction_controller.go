package controller

import "github.com/labstack/echo/v4"

type TrashTransactionController interface {
	CreateTrashTransactionController(c echo.Context) error
	GetAllTrashTransactionByBuyerId(c echo.Context) error
	GetAllTrashTransactionBySellerId(c echo.Context) error
	CreateTrashTransactionDoneController(c echo.Context) error
	GetTrashTransactionDoneByBuyerIdController(c echo.Context) error
	GetTrashTransactionDoneBySellerIdController(c echo.Context) error
}
