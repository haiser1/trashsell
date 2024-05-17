package dto

import (
	dto_buyer "mini-project/dto/buyer"
	dto_seller "mini-project/dto/seller"
	dto_trases "mini-project/dto/trashes"
)

type TrashTransactionGetResponse struct {
	Id             int    `json:"id" validate:"required"`
	EstimateWeight int    `json:"estimate_weight" validate:"required"`
	Location       string `json:"location" validate:"required"`
	Price          int    `json:"price" validate:"required"`
	Trash          dto_trases.TrashResponseGetAll
	Buyer          dto_buyer.BuyerResponseGetData
	Seller         dto_seller.SellerResponseGetData
}

type TrashTransactionCreateResponse struct {
	EstimateWeight int    `json:"estimate_weight" validate:"required"`
	Location       string `json:"location" validate:"required"`
	Price          int    `json:"price" validate:"required"`
	TrashId        int    `json:"trash_id" validate:"required"`
	BuyerId        int    `json:"buyer_id" validate:"required"`
	SellerId       int    `json:"seller_id" validate:"required"`
}

type TrashTransactionResponseDone struct {
	Id       int    `json:"id" `
	Weight   int    `json:"weight" `
	Price    int    `json:"price"`
	Location string `json:"location"`
}
