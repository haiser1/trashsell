package service

import (
	dto_seller "mini-project/dto/seller"
	"mini-project/model/domain"
)

type SellerService interface {
	RegisterSeller(request dto_seller.SellerRequestRegister) (*domain.Seller, error)
	LoginSeller(request dto_seller.SellerRequestLogin) (string, error)
	GetDataSeller(sellerId int) (*domain.Seller, error)
	UpdateSeller(request dto_seller.SellerRequestUpdate, id int) (*domain.Seller, error)
}
