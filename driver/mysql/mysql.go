package mysql

import (
	"fmt"
	"mini-project/driver/mysql/buyer"
	"mini-project/driver/mysql/seller"
	"mini-project/driver/mysql/trash"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Connection(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Migrate(db)

	return db

}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&buyer.Buyer{}, &trash.TypeTrash{}, &trash.Trash{}, &seller.Seller{})
}
