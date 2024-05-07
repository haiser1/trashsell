package service

import (
	"mini-project/helper"
	"mini-project/model/base"
	"mini-project/model/domain"
	"mini-project/repository"
	"strings"

	"github.com/go-playground/validator"
)

type BuyerServiceImpl struct {
	BuyerRepository repository.BuyerRepository
	Validate        *validator.Validate
}

func (service BuyerServiceImpl) RegisterBuyer(request base.BayerRequestRegister) base.BuyerResponseRegister {
	err := service.Validate.Struct(request)

	if err != nil {
		errorMessages := helper.HandleValidationErrors(err)
		errorMessage := strings.Join(errorMessages, ", ")
		return base.BuyerResponseRegister{
			Code:    400,
			Message: errorMessage,
			Data:    nil,
		}
	}
	findBuyer := service.BuyerRepository.FindBuyerByEmail(request.Email)

	if findBuyer != nil {
		return base.BuyerResponseRegister{
			Code:    400,
			Message: "Email already exist",
			Data:    nil,
		}
	}

	buyer := domain.Buyer{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	buyerId := service.BuyerRepository.RegisterBuyer(&buyer).ID
	responseData := base.ConvertResponseBuyerRegister{
		Id:    buyerId,
		Name:  buyer.Name,
		Email: buyer.Email,
	}
	return base.BuyerResponseRegister{
		Code:    200,
		Message: "Success",
		Data:    responseData,
	}
}
