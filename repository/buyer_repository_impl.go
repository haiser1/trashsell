package repository

import (
	"errors"
	"mini-project/helper"
	"mini-project/model/domain"

	"gorm.io/gorm"
)

type BuyerRepositoryImpl struct {
	DB *gorm.DB
}

func (repository BuyerRepositoryImpl) RegisterBuyer(buyer *domain.Buyer) *domain.Buyer {
	err := repository.DB.Create(&buyer).Error
	helper.PanicIfError(err)
	return buyer
}

func (repository BuyerRepositoryImpl) FindBuyerByEmail(email string) *domain.Buyer {
	var buyer domain.Buyer
	err := repository.DB.Where("email = ?", email).First(&buyer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Email tidak ditemukan
			return nil
		}
		// Error lain, tangani sesuai kebutuhan Anda
		helper.PanicIfError(err)
	}
	return &buyer
}
