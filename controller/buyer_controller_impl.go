package controller

import (
	"mini-project/model/base"
	"mini-project/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BuyerControllerImpl struct {
	BuyerService service.BuyerService
}

func (controller *BuyerControllerImpl) RegisterBuyerController(c echo.Context) error {
	buyerRequest := base.BayerRequestRegister{}
	c.Bind(&buyerRequest)
	buyerResponseRegisterSuccess := controller.BuyerService.RegisterBuyer(buyerRequest)

	if buyerResponseRegisterSuccess.Code != 200 {
		return c.JSON(http.StatusBadRequest, buyerResponseRegisterSuccess)
	}

	return c.JSON(http.StatusOK, buyerResponseRegisterSuccess)
}

func NewBuyerController(buyerService service.BuyerService) BuyerController {
	return &BuyerControllerImpl{
		BuyerService: buyerService,
	}
}
