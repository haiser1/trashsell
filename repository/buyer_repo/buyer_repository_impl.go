package repository

import (
	"mini-project/model/domain"

	"gorm.io/gorm"
)

type BuyerRepositoryImpl struct {
	DB *gorm.DB
}

func NewBuyerRepository(db *gorm.DB) *BuyerRepositoryImpl {
	return &BuyerRepositoryImpl{DB: db}
}

func (repository BuyerRepositoryImpl) RegisterBuyer(buyer *domain.Buyer) (*domain.Buyer, error) {
	if err := repository.DB.Create(&buyer).Error; err != nil {
		return nil, err
	}
	return buyer, nil
}

func (repository BuyerRepositoryImpl) FindBuyerByEmail(email string) (*domain.Buyer, error) {
	var buyer domain.Buyer
	if err := repository.DB.Where("email = ?", email).First(&buyer).Error; err != nil {
		return nil, err
	}
	return &buyer, nil
}

func (repository BuyerRepositoryImpl) FindBuyerById(id int) (*domain.Buyer, error) {
	var buyer domain.Buyer
	if err := repository.DB.Where("id = ?", id).First(&buyer).Error; err != nil {
		return nil, err
	}
	return &buyer, nil
}

func (repository BuyerRepositoryImpl) UpdateBuyer(buyer *domain.Buyer, buyerId int) (*domain.Buyer, error) {
	if err := repository.DB.Model(&buyer).Where("id = ?", buyerId).Updates(&buyer).Error; err != nil {
		return nil, err
	}
	return buyer, nil
}

func (repository BuyerRepositoryImpl) GetDataBuyer() (*domain.Buyer, error) {
	var buyer domain.Buyer
	if err := repository.DB.Find(&buyer).Error; err != nil {
		return &domain.Buyer{}, err
	}
	return &buyer, nil
}
