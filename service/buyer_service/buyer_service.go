package service

import (
	dto_buyer "mini-project/dto/buyer"
	"mini-project/model/domain"
)

type BuyerService interface {
	RegisterBuyer(request dto_buyer.BayerRequestRegister) (*domain.Buyer, error)
	LoginBuyer(request dto_buyer.BuyerRequestLogin) (string, error)
	GetDataBuyer(buyerId int) (*domain.Buyer, error)
	UpdateBuyer(request dto_buyer.BuyerRequestUpdate, id int) (*domain.Buyer, error)
}
