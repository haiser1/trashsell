package service

import (
	dto_buyer "mini-project/dto/buyer"
	"mini-project/helper"
	"mini-project/model/domain"
	repository "mini-project/repository/buyer_repo"

	"github.com/go-playground/validator"
)

type BuyerServiceImpl struct {
	BuyerRepository repository.BuyerRepository
	Validate        *validator.Validate
}

func NewBuyerService(buyerRepository repository.BuyerRepository, validate *validator.Validate) *BuyerServiceImpl {
	return &BuyerServiceImpl{
		BuyerRepository: buyerRepository,
		Validate:        validate,
	}
}

func (service BuyerServiceImpl) RegisterBuyer(request dto_buyer.BayerRequestRegister) (*domain.Buyer, error) {
	err := service.Validate.Struct(request)

	if err != nil {
		return &domain.Buyer{}, err
	}

	buyerEmail, _ := service.BuyerRepository.FindBuyerByEmail(request.Email)

	if buyerEmail != nil {
		return &domain.Buyer{}, helper.EmailAlreadyExist
	}

	hashPassword, _ := helper.HashPassword(request.Password)

	buyer, err := service.BuyerRepository.RegisterBuyer(&domain.Buyer{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPassword,
		Street:   request.Street,
		Province: request.Province,
		Regency:  request.Regency,
		City:     request.City,
	})

	return buyer, err

}

func (service BuyerServiceImpl) LoginBuyer(request dto_buyer.BuyerRequestLogin) (string, error) {

	if err := service.Validate.Struct(request); err != nil {
		return "", err
	}

	buyer, _ := service.BuyerRepository.FindBuyerByEmail(request.Email)

	if buyer == nil {
		return "", helper.InvalidLogin
	}

	if err := helper.ComparePassword(buyer.Password, request.Password); err != nil {
		return "", helper.InvalidLogin
	}

	token, err := helper.NewJwtService().GenerateToken(buyer.ID, "buyer")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service BuyerServiceImpl) GetDataBuyer(buyerId int) (*domain.Buyer, error) {
	buyer, err := service.BuyerRepository.FindBuyerById(buyerId)
	if err != nil {
		return &domain.Buyer{}, err
	}
	return buyer, nil
}

func (service BuyerServiceImpl) UpdateBuyer(request dto_buyer.BuyerRequestUpdate, id int) (*domain.Buyer, error) {

	errValidate := service.Validate.Struct(request)
	if errValidate != nil {
		return &domain.Buyer{}, errValidate
	}
	findBuyer, _ := service.BuyerRepository.FindBuyerById(id)
	if findBuyer == nil {
		return &domain.Buyer{}, helper.DataNotFound
	}
	buyer, err := service.BuyerRepository.UpdateBuyer(&domain.Buyer{
		ID:       findBuyer.ID,
		Name:     request.Name,
		Email:    findBuyer.Email,
		Password: findBuyer.Password,
		Street:   request.Street,
		Province: request.Province,
		Regency:  request.Regency,
		City:     request.City,
	}, id)
	return buyer, err
}
