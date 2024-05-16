package dto

type SellerResponseRegister struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Street   string `json:"street"`
	Province string `json:"province"`
	Regency  string `json:"regency"`
	City     string `json:"city"`
}

type SellerResponseLogin struct {
	Token string `json:"token"`
}

type SellerResponseGetData struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Street   string `json:"street"`
	Province string `json:"province"`
	Regency  string `json:"regency"`
	City     string `json:"city"`
}

type SellerResponseUpdate struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Street   string `json:"street"`
	Province string `json:"province"`
	Regency  string `json:"regency"`
	City     string `json:"city"`
}
