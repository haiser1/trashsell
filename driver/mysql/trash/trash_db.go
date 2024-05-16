package trash

import (
	"mini-project/driver/mysql/buyer"
	"time"
)

type Trash struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Price       int
	TypeTrashId int
	TypeTrash   TypeTrash `gorm:"foreignKey:TypeTrashId"`
	BuyerId     int
	Buyer       buyer.Buyer `gorm:"foreignKey:BuyerId"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
}

type TypeTrash struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
