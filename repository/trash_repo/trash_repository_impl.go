package repository

import (
	"fmt"
	"mini-project/model/domain"

	"gorm.io/gorm"
)

type TrashRepositoryImpl struct {
	DB *gorm.DB
}

func NewTrashRepository(db *gorm.DB) *TrashRepositoryImpl {
	return &TrashRepositoryImpl{DB: db}
}

func (repository TrashRepositoryImpl) FindTypeTrashByName(name string) (*domain.TypeTrash, error) {
	var typeTrash domain.TypeTrash
	if err := repository.DB.Where("name = ?", name).First(&typeTrash).Error; err != nil {
		return nil, err
	}
	return &typeTrash, nil
}

func (repository TrashRepositoryImpl) FindTrashByName(name string, buyerId int) (*domain.Trash, error) {
	var trash domain.Trash
	if err := repository.DB.Where("name = ?", name).Where("buyer_id = ?", buyerId).First(&trash).Error; err != nil {
		return nil, err
	}
	return &trash, nil
}

func (repository TrashRepositoryImpl) CreateTypeTrash(typeTrash *domain.TypeTrash) (*domain.TypeTrash, error) {
	if err := repository.DB.Create(&typeTrash).Error; err != nil {
		return nil, err
	}
	return typeTrash, nil
}

func (repository TrashRepositoryImpl) CreateTrash(trash *domain.Trash) (*domain.Trash, error) {
	if err := repository.DB.Create(&trash).Error; err != nil {
		return nil, err
	}
	return trash, nil
}

func (repository TrashRepositoryImpl) GetAllTypeTrash() ([]domain.TypeTrash, error) {

	var typeTrash []domain.TypeTrash
	if err := repository.DB.Find(&typeTrash).Error; err != nil {
		return nil, err
	}
	return typeTrash, nil
}

func (repository TrashRepositoryImpl) GetTypeTrashById(id int) (*domain.TypeTrash, error) {
	var typeTrash domain.TypeTrash
	if err := repository.DB.Where("id = ?", id).First(&typeTrash).Error; err != nil {
		return nil, err
	}
	return &typeTrash, nil
}

func (repository TrashRepositoryImpl) UpdateTypeTrash(typeTrash *domain.TypeTrash, id int) (*domain.TypeTrash, error) {

	if err := repository.DB.Model(&typeTrash).Where("id = ?", id).Updates(&typeTrash).Error; err != nil {
		return nil, err
	}
	return typeTrash, nil
}

func (repository TrashRepositoryImpl) GetAllTrashByBuyerId(buyerId int) ([]domain.Trash, error) {

	var trash []domain.Trash
	err := repository.DB.Joins("TypeTrash").Where("buyer_id = ?", buyerId).Find(&trash).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(trash)
	return trash, nil
}

func (repository TrashRepositoryImpl) GetAllTrash() ([]domain.Trash, error) {

	var trash []domain.Trash
	if err := repository.DB.Find(&trash).Error; err != nil {
		return nil, err
	}
	return trash, nil
}

func (repository TrashRepositoryImpl) GetTrashById(id, buyerId int) (*domain.Trash, error) {

	var trash domain.Trash

	if err := repository.DB.Joins("TypeTrash").Where("trashes.id = ?", id).Where("buyer_id = ?", buyerId).First(&trash).Error; err != nil {
		return nil, err
	}
	// repository.DB.Joins("TypeTrash").Where("buyer_id = ?", buyerId).Find(&trash).Erro
	fmt.Println("trash: ", trash)
	return &trash, nil
}

func (repository TrashRepositoryImpl) UpdateTrash(trash *domain.Trash, id int) (*domain.Trash, error) {

	if err := repository.DB.Model(&trash).Where("id = ?", id).Updates(&trash).Error; err != nil {
		return nil, err
	}
	return trash, nil
}

func (repository TrashRepositoryImpl) DelteTrash(trashId, buyerId int) (*domain.Trash, error) {

	var trash domain.Trash
	if err := repository.DB.Where("id = ?", trashId).Where("buyer_id = ?", buyerId).Delete(&trash).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
