package controller

import (
	dto_base "mini-project/dto/base"
	dto_buyer "mini-project/dto/buyer"
	"mini-project/helper"
	service "mini-project/service/buyer_service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BuyerControllerImpl struct {
	BuyerService service.BuyerService
}

func NewBuyerController(buyerService service.BuyerService) *BuyerControllerImpl {
	return &BuyerControllerImpl{
		BuyerService: buyerService,
	}
}

func (controller *BuyerControllerImpl) RegisterBuyerController(c echo.Context) error {

	request := dto_buyer.BayerRequestRegister{}
	c.Bind(&request)

	buyer, err := controller.BuyerService.RegisterBuyer(request)

	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_buyer.BuyerResponseRegister{
		Id:       buyer.ID,
		Email:    buyer.Email,
		Name:     buyer.Name,
		Street:   buyer.Street,
		Province: buyer.Province,
		Regency:  buyer.Regency,
		City:     buyer.City,
	}

	return c.JSON(http.StatusCreated, dto_base.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    res,
	})
}

func (controller *BuyerControllerImpl) LoginBuyerController(c echo.Context) error {
	request := dto_buyer.BuyerRequestLogin{}
	c.Bind(&request)

	token, err := controller.BuyerService.LoginBuyer(request)

	if err != nil {
		return helper.HandleError(c, err)

	}

	tokenRes := dto_buyer.BuyerResponseLogin{
		Token: token,
	}

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    tokenRes,
	})

}

func (controller *BuyerControllerImpl) GetDataBuyerController(c echo.Context) error {
	buyerId := c.Get("buyerId").(int)
	buyer, err := controller.BuyerService.GetDataBuyer(buyerId)

	if err != nil {
		return helper.HandleError(c, err)
	}

	res := dto_buyer.BuyerResponseGetData{
		Id:       buyer.ID,
		Email:    buyer.Email,
		Name:     buyer.Name,
		Street:   buyer.Street,
		Province: buyer.Province,
		Regency:  buyer.Regency,
		City:     buyer.City,
	}

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (controller *BuyerControllerImpl) UpdateBuyerController(c echo.Context) error {
	buyerId := c.Get("buyerId").(int)
	request := dto_buyer.BuyerRequestUpdate{}
	c.Bind(&request)
	buyer, err := controller.BuyerService.UpdateBuyer(request, buyerId)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_buyer.BuyerResponseUpdate{
		Id:       buyer.ID,
		Email:    buyer.Email,
		Name:     buyer.Name,
		Street:   buyer.Street,
		Province: buyer.Province,
		Regency:  buyer.Regency,
		City:     buyer.City,
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}
