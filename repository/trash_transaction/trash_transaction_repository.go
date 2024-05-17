package repository

import "mini-project/model/domain"

type TrashTransactionRepository interface {
	CreateTrashTransaction(trashTransaction *domain.TrashTransaction) (*domain.TrashTransaction, error)
	FindTrashById(trashId int) (*domain.Trash, error)
	FindTrashTransactionById(trashTransactionId int, buyerId int) (*domain.TrashTransaction, error)
	GetAllTrashTransactionBySellerId(sellerId int) ([]domain.TrashTransaction, error)
	GetAllTrashTransactionByBuyerId(buyerId int) ([]domain.TrashTransaction, error)
	CreateTrashTransactionDone(trashTransaction *domain.TrashTransaction, trashTransactionId int, buyerId int) (*domain.TrashTransaction, error)
	GetTrashTransactionDoneByBuyerId(buyerId int) ([]domain.TrashTransaction, error)
	GetTrashTransactionDoneBySellerId(sellerId int) ([]domain.TrashTransaction, error)
}
