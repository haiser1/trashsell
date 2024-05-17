package service

import (
	"fmt"
	dto_trash_transaction "mini-project/dto/trash_transaction"
	"mini-project/helper"
	"mini-project/model/domain"
	repository "mini-project/repository/trash_transaction"

	"github.com/go-playground/validator"
)

type TrashTransactionServiceImpl struct {
	TrashTransactionRepository repository.TrashTransactionRepository
	Validate                   *validator.Validate
}

func NewTrashTransactionService(trashTransactionRepository repository.TrashTransactionRepository, validate *validator.Validate) *TrashTransactionServiceImpl {
	return &TrashTransactionServiceImpl{
		TrashTransactionRepository: trashTransactionRepository,
		Validate:                   validate,
	}
}

func (service TrashTransactionServiceImpl) CreateTrashTransactionService(request dto_trash_transaction.TrashTransactionRequestCreateTransaction, trashId int, sellerId int) (*domain.TrashTransaction, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return &domain.TrashTransaction{}, err
	}

	findTrash, err := service.TrashTransactionRepository.FindTrashById(trashId)
	if err != nil {
		return &domain.TrashTransaction{}, err
	}

	if findTrash == nil {
		return &domain.TrashTransaction{}, helper.TrashNotFound
	}

	priceTrash := request.EstimateWeight * findTrash.Price

	buyerId := findTrash.BuyerId

	trashTransaction, err := service.TrashTransactionRepository.CreateTrashTransaction(&domain.TrashTransaction{
		EstimateWeight: request.EstimateWeight,
		Location:       request.Location,
		Price:          priceTrash,
		TrashId:        trashId,
		SellerId:       sellerId,
		BuyerId:        buyerId,
		TypeTrashId:    findTrash.TypeTrashId,
	})
	if err != nil {
		return &domain.TrashTransaction{}, err
	}

	return trashTransaction, nil
}

func (service TrashTransactionServiceImpl) GetAllTrashTransactionByBuyerId(buyerId int) ([]domain.TrashTransaction, error) {
	trashTransactionBuyer, err := service.TrashTransactionRepository.GetAllTrashTransactionByBuyerId(buyerId)

	if err != nil {
		return []domain.TrashTransaction{}, err
	}
	return trashTransactionBuyer, nil
}

func (service TrashTransactionServiceImpl) GetAllTrashTransactionBySellerId(sellerId int) ([]domain.TrashTransaction, error) {

	trashTransactionSeller, err := service.TrashTransactionRepository.GetAllTrashTransactionBySellerId(sellerId)

	if err != nil {
		return []domain.TrashTransaction{}, err
	}

	return trashTransactionSeller, nil
}

func (service TrashTransactionServiceImpl) CreateTrashTransactionDoneService(request dto_trash_transaction.TrashTransactionRequestTransactionDone, trashTransactionId int, buyerId int) (*domain.TrashTransaction, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return &domain.TrashTransaction{}, err
	}

	findTransaction, errFind := service.TrashTransactionRepository.FindTrashTransactionById(trashTransactionId, buyerId)
	if errFind != nil {
		return &domain.TrashTransaction{}, errFind
	}

	if findTransaction == nil {
		return &domain.TrashTransaction{}, helper.DataNotFound
	}

	findTrashPrice, errFindTrash := service.TrashTransactionRepository.FindTrashById(findTransaction.TrashId)
	if errFindTrash != nil {
		return &domain.TrashTransaction{}, errFindTrash
	}
	fmt.Println(findTrashPrice.Price)
	priceTrash := request.Weight * findTrashPrice.Price

	trashTransactionDone, err := service.TrashTransactionRepository.CreateTrashTransactionDone(&domain.TrashTransaction{
		EstimateWeight: request.Weight,
		Price:          priceTrash,
		Status:         true,
	}, trashTransactionId, buyerId)

	if err != nil {
		return &domain.TrashTransaction{}, err
	}
	return trashTransactionDone, nil
}

func (service TrashTransactionServiceImpl) GetTrashTransactionDoneServiceByBuyerId(buyerId int) ([]domain.TrashTransaction, error) {
	trashTransactionBuyer, err := service.TrashTransactionRepository.GetTrashTransactionDoneByBuyerId(buyerId)

	if err != nil {
		return []domain.TrashTransaction{}, err
	}
	return trashTransactionBuyer, nil
}

func (service TrashTransactionServiceImpl) GetTrashTransactionDoneServiceBySellerId(sellerId int) ([]domain.TrashTransaction, error) {

	trashTransactionSeller, err := service.TrashTransactionRepository.GetTrashTransactionDoneBySellerId(sellerId)

	if err != nil {
		return []domain.TrashTransaction{}, err
	}

	return trashTransactionSeller, nil
}
