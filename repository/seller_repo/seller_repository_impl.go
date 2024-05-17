package repository

import (
	"mini-project/model/domain"

	"gorm.io/gorm"
)

type SellerRepositoryImpl struct {
	DB *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepositoryImpl {
	return &SellerRepositoryImpl{DB: db}
}

func (repository SellerRepositoryImpl) RegisterSeller(seller *domain.Seller) (*domain.Seller, error) {
	if err := repository.DB.Create(&seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

func (repository SellerRepositoryImpl) FindSellerByEmail(email string) (*domain.Seller, error) {
	var seller domain.Seller
	if err := repository.DB.Where("email = ?", email).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}

func (repository SellerRepositoryImpl) FindSellerById(id int) (*domain.Seller, error) {
	var seller domain.Seller
	if err := repository.DB.Where("id = ?", id).First(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}

func (repository SellerRepositoryImpl) UpdateSeller(seller *domain.Seller, sellerId int) (*domain.Seller, error) {
	if err := repository.DB.Model(&seller).Where("id = ?", sellerId).Updates(&seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

func (repository SellerRepositoryImpl) GetDataSeller(sellerId int) (*domain.Seller, error) {
	var seller domain.Seller
	if err := repository.DB.Where("id = ?", sellerId).Find(&seller).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}
