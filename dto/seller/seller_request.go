package dto

type SellerRequestRegister struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Street   string `json:"street" validate:"required"`
	Province string `json:"province" validate:"required"`
	Regency  string `json:"regency" validate:"required"`
	City     string `json:"city" validate:"required"`
}

type SellerRequestLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SellerRequestUpdate struct {
	Name     string `json:"name" validate:"required"`
	Street   string `json:"street" validate:"required"`
	Province string `json:"province" validate:"required"`
	Regency  string `json:"regency" validate:"required"`
	City     string `json:"city" validate:"required"`
}
