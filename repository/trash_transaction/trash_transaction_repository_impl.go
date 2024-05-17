package repository

import (
	"fmt"
	"mini-project/model/domain"

	"gorm.io/gorm"
)

type TrashTransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTrashTransactionRepository(db *gorm.DB) *TrashTransactionRepositoryImpl {
	return &TrashTransactionRepositoryImpl{DB: db}
}

func (repository *TrashTransactionRepositoryImpl) CreateTrashTransaction(trashTransaction *domain.TrashTransaction) (*domain.TrashTransaction, error) {

	if err := repository.DB.Create(trashTransaction).Error; err != nil {
		return nil, err
	}
	return trashTransaction, nil

}

func (repository *TrashTransactionRepositoryImpl) FindTrashById(trashId int) (*domain.Trash, error) {

	var trash domain.Trash
	if err := repository.DB.Where("id = ?", trashId).Find(&trash).Error; err != nil {
		return nil, err
	}
	if trash.ID == 0 {
		return nil, nil
	}
	return &trash, nil
}

func (repository *TrashTransactionRepositoryImpl) GetAllTrashTransactionBySellerId(sellerId int) ([]domain.TrashTransaction, error) {
	var trashTransaction []domain.TrashTransaction
	if err := repository.DB.
		Joins("Trash").
		Joins("TypeTrash").
		Joins("Seller").
		Joins("Buyer").
		Where("seller_id = ?", sellerId).
		Where("status = ?", false).
		Find(&trashTransaction).Error; err != nil {
		return nil, err
	}
	return trashTransaction, nil
}

func (repository *TrashTransactionRepositoryImpl) GetAllTrashTransactionByBuyerId(buyerId int) ([]domain.TrashTransaction, error) {
	var trashTransaction []domain.TrashTransaction
	if err := repository.DB.
		Joins("Trash").
		Joins("TypeTrash").
		Joins("Seller").
		Joins("Buyer").
		Where("trash_transactions.buyer_id = ?", buyerId).
		Where("status = ?", false).
		Find(&trashTransaction).Error; err != nil {
		return nil, err
	}

	return trashTransaction, nil
}

func (repository *TrashTransactionRepositoryImpl) FindTrashTransactionById(trashTransactionId int, buyerId int) (*domain.TrashTransaction, error) {
	var trashTransaction domain.TrashTransaction
	fmt.Println(buyerId, trashTransactionId)

	if err := repository.DB.Where("id = ?", trashTransactionId).Where("buyer_id = ?", buyerId).Where("status = ?", false).First(&trashTransaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &trashTransaction, nil
}

func (repository *TrashTransactionRepositoryImpl) CreateTrashTransactionDone(trashTransaction *domain.TrashTransaction, trashTransactionId int, buyerId int) (*domain.TrashTransaction, error) {
	if err := repository.DB.
		Model(&domain.TrashTransaction{}).
		Where("id = ? AND buyer_id = ?", trashTransactionId, buyerId).
		Updates(map[string]interface{}{
			"price":           trashTransaction.Price,
			"estimate_weight": trashTransaction.EstimateWeight,
			"status":          trashTransaction.Status,
		}).Error; err != nil {
		return nil, err
	}

	// Fetch the updated transaction to return it
	var updatedTransaction domain.TrashTransaction
	if err := repository.DB.Where("id = ?", trashTransactionId).First(&updatedTransaction).Error; err != nil {
		return nil, err
	}

	return &updatedTransaction, nil
}

func (repository *TrashTransactionRepositoryImpl) GetTrashTransactionDoneByBuyerId(buyerId int) ([]domain.TrashTransaction, error) {
	var trashTransaction []domain.TrashTransaction
	if err := repository.DB.
		Joins("Trash").
		Joins("TypeTrash").
		Joins("Seller").
		Joins("Buyer").
		Where("trash_transactions.buyer_id = ?", buyerId).
		Where("status = ?", true).
		Find(&trashTransaction).Error; err != nil {
		return nil, err
	}
	return trashTransaction, nil
}

func (repository *TrashTransactionRepositoryImpl) GetTrashTransactionDoneBySellerId(sellerId int) ([]domain.TrashTransaction, error) {
	var trashTransaction []domain.TrashTransaction
	if err := repository.DB.
		Joins("Trash").
		Joins("TypeTrash").
		Joins("Seller").
		Joins("Buyer").
		Where("seller_id = ?", sellerId).
		Where("status = ?", true).
		Find(&trashTransaction).Error; err != nil {
		return nil, err
	}
	return trashTransaction, nil
}
