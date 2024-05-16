package repository

import (
	"mini-project/model/domain"
)

type BuyerRepository interface {
	RegisterBuyer(buyer *domain.Buyer) (*domain.Buyer, error)
	FindBuyerByEmail(email string) (*domain.Buyer, error)
	FindBuyerById(id int) (*domain.Buyer, error)
	UpdateBuyer(buyer *domain.Buyer, buyerId int) (*domain.Buyer, error)
}
