package dto

import (
	buyer "mini-project/dto/buyer"
)

type TrashResponseGetAll struct {
	Id    int               `json:"id"`
	Name  string            `json:"name"`
	Price int               `json:"price"`
	Type  TrashTypeResponse `json:"type_trash"`
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

type TrashResponseDataPaggination struct {
	Id    int                        `json:"id"`
	Name  string                     `json:"name"`
	Price int                        `json:"price"`
	Type  TrashTypeResponse          `json:"type"`
	Buyer buyer.BuyerResponseGetData `json:"buyer"`
}

type TrashResponsePaggination struct {
	Data  []TrashResponseDataPaggination
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
