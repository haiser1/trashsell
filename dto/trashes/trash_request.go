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

type TrashRequestPaggination struct {
	Limit     int    `json:"limit" validate:"min=1 isdefault=10 max=100"`
	Offset    int    `json:"offset" validate:"min=1 isdefault=0 max=100"`
	NameTrash string `json:"name_trash"`
	TypeTrash string `json:"type_trash"`
}
