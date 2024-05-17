package dto

type TrashTransactionRequestCreateTransaction struct {
	EstimateWeight int    `json:"estimate_weight" validate:"required"`
	Location       string `json:"location" validate:"required"`
}

type TrashTransactionRequestTransactionDone struct {
	Weight int `json:"weight" validate:"required"`
}
