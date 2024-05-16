package controller

import (
	"fmt"
	dto_base "mini-project/dto/base"
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
	fmt.Println("typeTrashId: ", typeTrashId)

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
