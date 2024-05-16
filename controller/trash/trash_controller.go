package controller

import "github.com/labstack/echo/v4"

type TrashController interface {
	CreateTypeTrashController(c echo.Context) (err error)
	CreateTrashController(c echo.Context) (err error)
	GetAllTypeTrashController(c echo.Context) (err error)
	UpdateTypeTrashController(c echo.Context) (err error)
	GetAllTrashByBuyerIdController(c echo.Context) (err error)
	GetAllTrashController(c echo.Context) (err error)
	GetTrashByIdController(c echo.Context) (err error)
	UpdateTrashController(c echo.Context) (err error)
	DeleteTrashController(c echo.Context) (err error)
}
