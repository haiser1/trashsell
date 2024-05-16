package seller

import "time"

type Seller struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Street    string
	Province  string
	Regency   string
	City      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
