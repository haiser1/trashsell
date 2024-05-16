package dto

type TrashTypeRequestUpdate struct {
	Name string `json:"name" validate:"required"`
}

type BuyerCreateTypeTrash struct {
	Name string `json:"name" validate:"required"`
}
