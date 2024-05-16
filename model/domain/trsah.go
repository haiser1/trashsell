package domain

import "time"

type Trash struct {
	ID          int
	Name        string
	Price       int
	TypeTrashId int
	TypeTrash   TypeTrash
	BuyerId     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TypeTrash struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
