package controller

import (
	dto_base "mini-project/dto/base"
	dto_buyer "mini-project/dto/buyer"
	dto_trashes "mini-project/dto/trashes"
	"mini-project/helper"
	service "mini-project/service/trash_service"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TrashControllerImpl struct {
	TrashService service.TrashService
}

func NewTrashController(trashService service.TrashService) *TrashControllerImpl {
	return &TrashControllerImpl{
		TrashService: trashService,
	}
}

func (controller *TrashControllerImpl) CreateTypeTrashController(c echo.Context) error {
	request := dto_trashes.BuyerCreateTypeTrash{}
	c.Bind(&request)

	typeTrash, err := controller.TrashService.CreateTypeTrash(request)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_trashes.BuyerCreateTypeTrash{
		Name: typeTrash.Name,
	}
	return c.JSON(http.StatusCreated, dto_base.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    res,
	})
}

func (controller *TrashControllerImpl) CreateTrashController(c echo.Context) error {
	buyerId := c.Get("buyerId").(int)
	request := dto_trashes.TrashRequestCreate{}
	c.Bind(&request)

	trash, err := controller.TrashService.CreateTrash(request, buyerId)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_trashes.TrashRequestCreate{
		Name:      trash.Name,
		Price:     trash.Price,
		TypeTrash: request.TypeTrash,
	}
	return c.JSON(http.StatusCreated, dto_base.BaseResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    res,
	})
}

func (controller *TrashControllerImpl) GetAllTypeTrashController(c echo.Context) error {

	typeTrash, err := controller.TrashService.GetAllTypeTrash()
	if err != nil {
		return helper.HandleError(c, err)
	}

	res := []dto_trashes.TrashTypeResponse{}
	for _, value := range typeTrash {
		res = append(res, dto_trashes.TrashTypeResponse{
			Id:   value.ID,
			Name: value.Name,
		})
	}

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (controller *TrashControllerImpl) UpdateTypeTrashController(c echo.Context) error {
	request := dto_trashes.TrashTypeRequestUpdate{}

	typeTrashId := c.Param("id")

	id, errConvert := strconv.Atoi(typeTrashId)

	if errConvert != nil {
		return helper.HandleError(c, helper.ErrorIdParam)
	}

	c.Bind(&request)

	typeTrash, err := controller.TrashService.UpdateTypeTrash(request, id)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_trashes.TrashTypeResponse{
		Id:   id,
		Name: typeTrash.Name,
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})

}

func (controller *TrashControllerImpl) GetAllTrashByBuyerIdController(c echo.Context) error {

	buyerId := c.Get("buyerId").(int)
	trash, err := controller.TrashService.GetAllTrashByBuyerId(buyerId)
	if err != nil {
		return helper.HandleError(c, err)
	}

	res := []dto_trashes.TrashResponseGetTrashByIdBuyer{}

	for _, value := range trash {
		res = append(res, dto_trashes.TrashResponseGetTrashByIdBuyer{
			Id:    value.ID,
			Name:  value.Name,
			Price: value.Price,
			Type: dto_trashes.TrashTypeResponse{
				Id:   value.TypeTrash.ID,
				Name: value.TypeTrash.Name,
			},
		})
	}

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (controller *TrashControllerImpl) GetAllTrashController(c echo.Context) error {

	trash, err := controller.TrashService.GetAllTrash()
	if err != nil {
		return helper.HandleError(c, err)
	}

	res := []dto_trashes.TrashResponseGetAll{}

	for _, value := range trash {
		res = append(res, dto_trashes.TrashResponseGetAll{
			Id:   value.ID,
			Name: value.Name,
			Type: dto_trashes.TrashTypeResponse{
				Id:   value.TypeTrash.ID,
				Name: value.TypeTrash.Name,
			},
		})
	}

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (controller *TrashControllerImpl) GetTrashByIdController(c echo.Context) error {
	buyerId := c.Get("buyerId").(int)

	trashId := c.Param("id")

	id, errConvert := strconv.Atoi(trashId)

	if errConvert != nil {
		return helper.HandleError(c, helper.ErrorIdParam)
	}

	trash, err := controller.TrashService.GetTrashById(id, buyerId)

	if err != nil {
		return helper.HandleError(c, err)
	}

	res := dto_trashes.TrashResponseGetTrashByIdBuyer{
		Id:    trash.ID,
		Name:  trash.Name,
		Price: trash.Price,
		Type: dto_trashes.TrashTypeResponse{
			Id:   trash.TypeTrash.ID,
			Name: trash.TypeTrash.Name,
		},
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (controller *TrashControllerImpl) UpdateTrashController(c echo.Context) error {
	buyerId := c.Get("buyerId").(int)
	trashId := c.Param("id")
	id, _ := strconv.Atoi(trashId)

	request := dto_trashes.TrashRequestUpdate{}

	c.Bind(&request)

	trash, err := controller.TrashService.UpdateTrash(request, id, buyerId)
	if err != nil {
		return helper.HandleError(c, err)
	}
	res := dto_trashes.TrashResponseGetTrashByIdBuyer{
		Id:    id,
		Name:  trash.Name,
		Price: trash.Price,
		Type: dto_trashes.TrashTypeResponse{
			Id:   trash.TypeTrash.ID,
			Name: trash.TypeTrash.Name,
		},
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})
}

func (controller *TrashControllerImpl) DeleteTrashController(c echo.Context) error {
	buyerId := c.Get("buyerId").(int)
	trashId := c.Param("id")
	id, _ := strconv.Atoi(trashId)

	_, err := controller.TrashService.DelteTrash(id, buyerId)
	if err != nil {
		return helper.HandleError(c, err)
	}
	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    nil,
	})
}

func (controller *TrashControllerImpl) GetTrashPagginationController(c echo.Context) error {

	page := c.QueryParam("page")
	pageSize := c.QueryParam("page_size")
	trashName := c.QueryParam("name_trash")
	typeTrash := c.QueryParam("type_trash")
	buyerName := c.QueryParam("buyer_name")
	pageInt, _ := strconv.Atoi(page)
	pageSizeint, _ := strconv.Atoi(pageSize)

	if pageInt == 0 {
		pageInt = 1
	}

	if pageSizeint == 0 {
		pageSizeint = 10
	}

	trash, err := controller.TrashService.GetListTrashPagination(pageInt, pageSizeint, trashName, typeTrash, buyerName)
	if err != nil {
		return helper.HandleError(c, err)
	}

	data := dto_trashes.TrashResponsePaggination{}

	for _, value := range trash {
		data.Data = append(data.Data, dto_trashes.TrashResponseDataPaggination{
			Id:    value.ID,
			Name:  value.Name,
			Price: value.Price,
			Type: dto_trashes.TrashTypeResponse{
				Id:   value.TypeTrash.ID,
				Name: value.TypeTrash.Name,
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

	res := dto_trashes.TrashResponsePaggination{
		Data:  data.Data,
		Total: len(trash),
		Page:  pageInt,
		Limit: pageSizeint,
	}

	return c.JSON(http.StatusOK, dto_base.BaseResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    res,
	})

}
