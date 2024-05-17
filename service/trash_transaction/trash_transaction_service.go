package service

import (
	dto_trash_transaction "mini-project/dto/trash_transaction"
	"mini-project/model/domain"
)

type TrashTrnsactionService interface {
	CreateTrashTransactionService(request dto_trash_transaction.TrashTransactionRequestCreateTransaction, trashId int, sellerId int) (*domain.TrashTransaction, error)
	GetAllTrashTransactionByBuyerId(buyerId int) ([]domain.TrashTransaction, error)
	GetAllTrashTransactionBySellerId(sellerId int) ([]domain.TrashTransaction, error)
	CreateTrashTransactionDoneService(request dto_trash_transaction.TrashTransactionRequestTransactionDone, trashTransactionId int, buyerId int) (*domain.TrashTransaction, error)
	GetTrashTransactionDoneServiceByBuyerId(buyerId int) ([]domain.TrashTransaction, error)
	GetTrashTransactionDoneServiceBySellerId(sellerId int) ([]domain.TrashTransaction, error)
}
