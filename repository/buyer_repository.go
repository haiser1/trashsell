package repository

import (
	"mini-project/model/domain"
)

type BuyerRepository interface {
	RegisterBuyer(buyer *domain.Buyer) *domain.Buyer
	FindBuyerByEmail(email string) *domain.Buyer
}
