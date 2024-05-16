package dto

type TrashResponseGetAll struct {
	Id   int               `json:"id"`
	Name string            `json:"name"`
	Type TrashTypeResponse `json:"type_trash"`
}

type TrashResponseGetTrashByIdBuyer struct {
	Id    int               `json:"id"`
	Name  string            `json:"name"`
	Price int               `json:"price"`
	Type  TrashTypeResponse `json:"type_trash"`
}

type TrashResponseUpdate struct {
	Id    int               `json:"id"`
	Name  string            `json:"name"`
	Price int               `json:"price"`
	Type  TrashTypeResponse `json:"type"`
}

type TrashTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
