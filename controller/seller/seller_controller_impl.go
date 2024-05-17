package controller

import (
	dto_base "mini-project/dto/base"
	dto_seller "mini-project/dto/seller"
	"mini-project/helper"
	service "mini-project/service/seller_service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SellerControllerImpl struct {
	SellerService service.SellerService
}

func NewSellerController(sellerService service.SellerService) *SellerControllerImpl {
	return &SellerControllerImpl{
		SellerService: sellerService,
	}
}

func (controller SellerControllerImpl) RegisterSellerController(c echo.Context) error {
	request := dto_seller.SellerRequestRegister{}

	c.Bind(&request)

	seller, err := controller.SellerService.RegisterSeller(request)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_seller.SellerResponseRegister{
		Id:       seller.ID,
		Email:    seller.Email,
		Name:     seller.Name,
		Street:   seller.Street,
		Province: seller.Province,
		Regency:  seller.Regency,
		City:     seller.City,
	}
	return c.JSON(http.StatusCreated, dto_base.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    res,
	})
}

func (controller SellerControllerImpl) LoginSellerController(c echo.Context) error {
	request := dto_seller.SellerRequestLogin{}
	c.Bind(&request)

	token, err := controller.SellerService.LoginSeller(request)
	if err != nil {
		return helper.HandleError(c, err)
	}
	tokenRes := dto_seller.SellerResponseLogin{
		Token: token,
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    tokenRes,
	})
}

func (controller SellerControllerImpl) GetDataSellerController(c echo.Context) error {
	sellerId := c.Get("sellerId").(int)
	seller, err := controller.SellerService.GetDataSeller(sellerId)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_seller.SellerResponseGetData{
		Id:       seller.ID,
		Email:    seller.Email,
		Name:     seller.Name,
		Street:   seller.Street,
		Province: seller.Province,
		Regency:  seller.Regency,
		City:     seller.City,
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (controller SellerControllerImpl) UpdateSellerController(c echo.Context) error {
	request := dto_seller.SellerRequestUpdate{}
	c.Bind(&request)
	sellerId := c.Get("sellerId").(int)
	seller, err := controller.SellerService.UpdateSeller(request, sellerId)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_seller.SellerResponseUpdate{
		Id:       seller.ID,
		Email:    seller.Email,
		Name:     seller.Name,
		Street:   seller.Street,
		Province: seller.Province,
		Regency:  seller.Regency,
		City:     seller.City,
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}
