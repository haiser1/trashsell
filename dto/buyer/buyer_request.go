package dto

type BayerRequestRegister struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Street   string `json:"street" validate:"required"`
	Province string `json:"province" validate:"required"`
	Regency  string `json:"regency" validate:"required"`
	City     string `json:"city" validate:"required"`
}

type BuyerRequestLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type BuyerRequestUpdate struct {
	Name     string `json:"name" validate:"required"`
	Street   string `json:"street" validate:"required"`
	Province string `json:"province" validate:"required"`
	Regency  string `json:"regency" validate:"required"`
	City     string `json:"city" validate:"required"`
}
