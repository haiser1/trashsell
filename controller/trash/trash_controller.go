package controller

import "github.com/labstack/echo/v4"

type TrashController interface {
	CreateTypeTrashController(c echo.Context) error
	CreateTrashController(c echo.Context) error
	GetAllTypeTrashController(c echo.Context) error
	UpdateTypeTrashController(c echo.Context) error
	GetAllTrashByBuyerIdController(c echo.Context) error
	GetAllTrashController(c echo.Context) error
	GetTrashByIdController(c echo.Context) error
	UpdateTrashController(c echo.Context) error
	DeleteTrashController(c echo.Context) error
	GetTrashPagginationController(c echo.Context) error
}
