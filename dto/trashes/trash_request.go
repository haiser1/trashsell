package dto

type TrashRequestUpdate struct {
	Name      string `json:"name" validate:"required"`
	Price     int    `json:"price" validate:"required"`
	TypeTrash string `json:"type_trash" validate:"required"`
}

type TrashRequestCreate struct {
	Name      string `json:"name" validate:"required"`
	Price     int    `json:"price" validate:"required"`
	TypeTrash string `json:"type_trash" validate:"required"`
}

type TrashRequestById struct {
	Id int `json:"id" validate:"required"`
}
