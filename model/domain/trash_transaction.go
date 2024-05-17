package domain

import "time"

type TrashTransaction struct {
	ID             int
	EstimateWeight int
	Location       string
	Price          int
	TrashId        int
	Trash          Trash
	SellerId       int
	Seller         Seller
	BuyerId        int
	Buyer          Buyer
	TypeTrashId    int
	TypeTrash      TypeTrash
	Status         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
