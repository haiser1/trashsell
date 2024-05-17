package controller

import (
	"fmt"
	dto_base "mini-project/dto/base"
	dto_buyer "mini-project/dto/buyer"
	dto_seller "mini-project/dto/seller"
	dto_trash_transaction "mini-project/dto/trash_transaction"
	dto_trashes "mini-project/dto/trashes"
	"mini-project/helper"
	service "mini-project/service/trash_transaction"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TrashTransactionControllerImpl struct {
	TrashTransactionService service.TrashTrnsactionService
}

func NewTrashTransactionController(trashTransactionService service.TrashTrnsactionService) TrashTransactionControllerImpl {
	return TrashTransactionControllerImpl{
		TrashTransactionService: trashTransactionService,
	}
}

func (controller TrashTransactionControllerImpl) CreateTrashTransactionController(c echo.Context) error {
	id := c.Get("sellerId").(int)
	trashId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.HandleError(c, helper.ErrorIdParam)
	}

	request := dto_trash_transaction.TrashTransactionRequestCreateTransaction{}
	c.Bind(&request)

	trashTransaction, err := controller.TrashTransactionService.CreateTrashTransactionService(request, trashId, id)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_trash_transaction.TrashTransactionCreateResponse{
		EstimateWeight: trashTransaction.EstimateWeight,
		Location:       trashTransaction.Location,
		Price:          trashTransaction.Price,
		TrashId:        trashTransaction.TrashId,
		BuyerId:        trashTransaction.BuyerId,
		SellerId:       trashTransaction.SellerId,
	}
	return c.JSON(http.StatusCreated, dto_base.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    res,
	})

}

func (controller TrashTransactionControllerImpl) GetAllTrashTransactionByBuyerId(c echo.Context) error {

	id := c.Get("buyerId").(int)
	trashTransaction, err := controller.TrashTransactionService.GetAllTrashTransactionByBuyerId(id)
	if err != nil {
		return helper.HandleError(c, err)
	}
	data := []dto_trash_transaction.TrashTransactionGetResponse{}

	for _, value := range trashTransaction {
		data = append(data, dto_trash_transaction.TrashTransactionGetResponse{
			Id:             value.ID,
			EstimateWeight: value.EstimateWeight,
			Location:       value.Location,
			Price:          value.Price,
			Trash: dto_trashes.TrashResponseGetAll{
				Id:    value.Trash.ID,
				Name:  value.Trash.Name,
				Price: value.Trash.Price,
				Type: dto_trashes.TrashTypeResponse{
					Id:   value.Trash.TypeTrashId,
					Name: value.TypeTrash.Name,
				},
			},
			Seller: dto_seller.SellerResponseGetData{
				Id:       value.Seller.ID,
				Name:     value.Seller.Name,
				Email:    value.Seller.Email,
				Street:   value.Seller.Street,
				Province: value.Seller.Province,
				Regency:  value.Seller.Regency,
				City:     value.Seller.City,
			},

			Buyer: dto_buyer.BuyerResponseGetData{
				Id:       value.Buyer.ID,
				Name:     value.Buyer.Name,
				Email:    value.Buyer.Email,
				Street:   value.Buyer.Street,
				Province: value.Buyer.Province,
				Regency:  value.Buyer.Regency,
				City:     value.Buyer.City,
			},
		})
	}

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (controller TrashTransactionControllerImpl) GetAllTrashTransactionBySellerId(c echo.Context) error {

	id := c.Get("sellerId").(int)
	trashTransaction, err := controller.TrashTransactionService.GetAllTrashTransactionBySellerId(id)
	if err != nil {
		return helper.HandleError(c, err)
	}
	data := []dto_trash_transaction.TrashTransactionGetResponse{}

	for _, value := range trashTransaction {
		data = append(data, dto_trash_transaction.TrashTransactionGetResponse{
			Id:             value.ID,
			EstimateWeight: value.EstimateWeight,
			Location:       value.Location,
			Price:          value.Price,
			Trash: dto_trashes.TrashResponseGetAll{
				Id:    value.Trash.ID,
				Name:  value.Trash.Name,
				Price: value.Trash.Price,
				Type: dto_trashes.TrashTypeResponse{
					Id:   value.Trash.TypeTrashId,
					Name: value.TypeTrash.Name,
				},
			},
			Seller: dto_seller.SellerResponseGetData{
				Id:       value.Seller.ID,
				Name:     value.Seller.Name,
				Email:    value.Seller.Email,
				Street:   value.Seller.Street,
				Province: value.Seller.Province,
				Regency:  value.Seller.Regency,
				City:     value.Seller.City,
			},

			Buyer: dto_buyer.BuyerResponseGetData{
				Id:       value.Buyer.ID,
				Name:     value.Buyer.Name,
				Email:    value.Buyer.Email,
				Street:   value.Buyer.Street,
				Province: value.Buyer.Province,
				Regency:  value.Buyer.Regency,
				City:     value.Buyer.City,
			},
		})
	}

	fmt.Print(data)

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (controller TrashTransactionControllerImpl) CreateTrashTransactionDoneController(c echo.Context) error {
	id := c.Param("id")
	transactionId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		return helper.HandleError(c, errConvert)
	}

	buyerId := c.Get("buyerId").(int)
	request := dto_trash_transaction.TrashTransactionRequestTransactionDone{}
	c.Bind(&request)

	trashTransaction, err := controller.TrashTransactionService.CreateTrashTransactionDoneService(request, transactionId, buyerId)

	if err != nil {
		return helper.HandleError(c, err)
	}

	res := dto_trash_transaction.TrashTransactionResponseDone{
		Id:       trashTransaction.ID,
		Weight:   trashTransaction.EstimateWeight,
		Price:    trashTransaction.Price,
		Location: trashTransaction.Location,
	}

	return c.JSON(http.StatusCreated, dto_base.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    res,
	})

}

func (controller TrashTransactionControllerImpl) GetTrashTransactionDoneByBuyerIdController(c echo.Context) error {
	buyerId := c.Get("buyerId").(int)
	fmt.Println(buyerId)

	trashTransaction, err := controller.TrashTransactionService.GetTrashTransactionDoneServiceByBuyerId(buyerId)

	if err != nil {
		return helper.HandleError(c, err)
	}

	data := []dto_trash_transaction.TrashTransactionGetResponse{}
	for _, value := range trashTransaction {
		data = append(data, dto_trash_transaction.TrashTransactionGetResponse{
			Id:             value.ID,
			EstimateWeight: value.EstimateWeight,
			Location:       value.Location,
			Price:          value.Price,
			Trash: dto_trashes.TrashResponseGetAll{
				Id:    value.Trash.ID,
				Name:  value.Trash.Name,
				Price: value.Trash.Price,
				Type: dto_trashes.TrashTypeResponse{
					Id:   value.Trash.TypeTrashId,
					Name: value.TypeTrash.Name,
				},
			},
			Seller: dto_seller.SellerResponseGetData{
				Id:       value.Seller.ID,
				Name:     value.Seller.Name,
				Email:    value.Seller.Email,
				Street:   value.Seller.Street,
				Province: value.Seller.Province,
				Regency:  value.Seller.Regency,
				City:     value.Seller.City,
			},
			Buyer: dto_buyer.BuyerResponseGetData{
				Id:       value.Buyer.ID,
				Name:     value.Buyer.Name,
				Email:    value.Buyer.Email,
				Street:   value.Buyer.Street,
				Province: value.Buyer.Province,
				Regency:  value.Buyer.Regency,
				City:     value.Buyer.City,
			},
		})
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func (controller TrashTransactionControllerImpl) GetTrashTransactionDoneBySellerIdController(c echo.Context) error {
	sellerId := c.Get("sellerId").(int)

	trashTransaction, err := controller.TrashTransactionService.GetTrashTransactionDoneServiceBySellerId(sellerId)

	if err != nil {
		return helper.HandleError(c, err)
	}

	data := []dto_trash_transaction.TrashTransactionGetResponse{}
	for _, value := range trashTransaction {
		data = append(data, dto_trash_transaction.TrashTransactionGetResponse{
			Id:             value.ID,
			EstimateWeight: value.EstimateWeight,
			Location:       value.Location,
			Price:          value.Price,
			Trash: dto_trashes.TrashResponseGetAll{
				Id:    value.Trash.ID,
				Name:  value.Trash.Name,
				Price: value.Trash.Price,
				Type: dto_trashes.TrashTypeResponse{
					Id:   value.Trash.TypeTrashId,
					Name: value.TypeTrash.Name,
				},
			},
			Seller: dto_seller.SellerResponseGetData{
				Id:       value.Seller.ID,
				Name:     value.Seller.Name,
				Email:    value.Seller.Email,
				Street:   value.Seller.Street,
				Province: value.Seller.Province,
				Regency:  value.Seller.Regency,
				City:     value.Seller.City,
			},
			Buyer: dto_buyer.BuyerResponseGetData{
				Id:       value.Buyer.ID,
				Name:     value.Buyer.Name,
				Email:    value.Buyer.Email,
				Street:   value.Buyer.Street,
				Province: value.Buyer.Province,
				Regency:  value.Buyer.Regency,
				City:     value.Buyer.City,
			},
		})
	}

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
