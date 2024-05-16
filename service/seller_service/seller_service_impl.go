package service

import (
	dto_seller "mini-project/dto/seller"
	"mini-project/helper"
	"mini-project/model/domain"
	repository "mini-project/repository/seller_repo"

	"github.com/go-playground/validator"
)

type SellerServiceImpl struct {
	SellerRepository repository.SellerRepository
	Validate         *validator.Validate
}

func NewSellerService(sellerRepository repository.SellerRepository, validate *validator.Validate) *SellerServiceImpl {
	return &SellerServiceImpl{
		SellerRepository: sellerRepository,
		Validate:         validate,
	}
}

func (service SellerServiceImpl) RegisterSeller(request dto_seller.SellerRequestRegister) (*domain.Seller, error) {

	err := service.Validate.Struct(request)

	if err != nil {
		return &domain.Seller{}, err
	}

	findEmailSeller, _ := service.SellerRepository.FindSellerByEmail(request.Email)
	if findEmailSeller != nil {
		return &domain.Seller{}, helper.EmailAlreadyExist
	}

	hashPassword, _ := helper.HashPassword(request.Password)
	seller, err := service.SellerRepository.RegisterSeller(&domain.Seller{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPassword,
		Street:   request.Street,
		Regency:  request.Regency,
		City:     request.City,
		Province: request.Province,
	})
	return seller, err
}

func (service SellerServiceImpl) LoginSeller(request dto_seller.SellerRequestLogin) (string, error) {
	if err := service.Validate.Struct(request); err != nil {
		return "", err
	}
	seller, _ := service.SellerRepository.FindSellerByEmail(request.Email)
	if seller == nil {
		return "", helper.InvalidLogin
	}

	if err := helper.ComparePassword(seller.Password, request.Password); err != nil {
		return "", helper.InvalidLogin
	}

	token, err := helper.NewJwtService().GenerateToken(seller.ID, "seller")

	if err != nil {
		return "", err
	}
	return token, nil
}

func (service SellerServiceImpl) GetDataSeller() (*domain.Seller, error) {
	seller, err := service.SellerRepository.GetDataSeller()
	if err != nil {
		return &domain.Seller{}, err
	}
	return seller, nil
}

func (service SellerServiceImpl) UpdateSeller(request dto_seller.SellerRequestUpdate, id int) (*domain.Seller, error) {
	if err := service.Validate.Struct(request); err != nil {
		return &domain.Seller{}, err
	}

	findSeller, _ := service.SellerRepository.FindSellerById(id)

	if findSeller == nil {
		return &domain.Seller{}, helper.DataNotFound
	}
	seller, err := service.SellerRepository.UpdateSeller(&domain.Seller{
		ID:       id,
		Name:     request.Name,
		Email:    findSeller.Email,
		Password: findSeller.Password,
		Street:   request.Street,
		Regency:  request.Regency,
		City:     request.City,
		Province: request.Province,
	}, id)
	return seller, err
}
