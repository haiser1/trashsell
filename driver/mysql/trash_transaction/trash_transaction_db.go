package trash_transaction

import (
	"mini-project/driver/mysql/buyer"
	"mini-project/driver/mysql/seller"
	"mini-project/driver/mysql/trash"
	"time"
)

type TrashTransaction struct {
	ID             int `gorm:"primaryKey"`
	EstimateWeight int
	Location       string
	Price          int
	TrashId        int
	Trash          trash.Trash `gorm:"foreignKey:TrashId"`
	SellerId       int
	Seller         seller.Seller `gorm:"foreignKey:SellerId"`
	BuyerId        int
	Buyer          buyer.Buyer `gorm:"foreignKey:BuyerId"`
	TypeTrashId    int
	TypeTrash      trash.TypeTrash `gorm:"foreignKey:TypeTrashId"`
	Status         bool            `gorm:"default:false"`
	CreatedAt      time.Time       `gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime"`
}
