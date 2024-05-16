package repository

import "mini-project/model/domain"

type SellerRepository interface {
	RegisterSeller(seller *domain.Seller) (*domain.Seller, error)
	FindSellerByEmail(email string) (*domain.Seller, error)
	FindSellerById(id int) (*domain.Seller, error)
	UpdateSeller(seller *domain.Seller, sellerId int) (*domain.Seller, error)
	GetDataSeller() (*domain.Seller, error)
}
