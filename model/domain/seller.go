package domain

import "time"

type Seller struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Street    string
	Province  string
	Regency   string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
