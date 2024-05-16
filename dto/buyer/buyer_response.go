package dto

type BuyerResponseRegister struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Street   string `json:"street"`
	Province string `json:"province"`
	Regency  string `json:"regency"`
	City     string `json:"city"`
}

type BuyerResponseLogin struct {
	Token string `json:"token"`
}

type BuyerResponseGetData struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Street   string `json:"street"`
	Province string `json:"province"`
	Regency  string `json:"regency"`
	City     string `json:"city"`
}

type BuyerResponseUpdate struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Street   string `json:"street"`
	Province string `json:"province"`
	Regency  string `json:"regency"`
	City     string `json:"city"`
}
