package service

import (
	"fmt"
	dto_trashes "mini-project/dto/trashes"
	"mini-project/helper"
	"mini-project/model/domain"
	repository "mini-project/repository/trash_repo"

	"github.com/go-playground/validator"
)

type TrashServiceImpl struct {
	TrashRepository repository.TrashRepository
	Validate        *validator.Validate
}

func NewTrashService(trashRepository repository.TrashRepository, validate *validator.Validate) *TrashServiceImpl {
	return &TrashServiceImpl{
		TrashRepository: trashRepository,
		Validate:        validate,
	}
}

func (service TrashServiceImpl) CreateTypeTrash(request dto_trashes.BuyerCreateTypeTrash) (*domain.TypeTrash, error) {

	err := service.Validate.Struct(request)

	if err != nil {
		return &domain.TypeTrash{}, err
	}

	findNameTypeTrash, _ := service.TrashRepository.FindTypeTrashByName(request.Name)
	if findNameTypeTrash != nil {
		return &domain.TypeTrash{}, helper.TypeTrashAlreadyExist
	}

	typeTrash, err := service.TrashRepository.CreateTypeTrash(&domain.TypeTrash{
		Name: request.Name,
	})

	return typeTrash, err

}

func (service TrashServiceImpl) CreateTrash(request dto_trashes.TrashRequestCreate, buyerId int) (*domain.Trash, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return &domain.Trash{}, err
	}

	findNameTrash, _ := service.TrashRepository.FindTrashByName(request.Name, buyerId)
	if findNameTrash != nil {
		return &domain.Trash{}, helper.TrashAlreadyExist
	}

	typeTrash, _ := service.TrashRepository.FindTypeTrashByName(request.TypeTrash)
	if typeTrash == nil {
		return &domain.Trash{}, helper.TypeTrashNotFound
	}

	trash, err := service.TrashRepository.CreateTrash(&domain.Trash{
		Name:        request.Name,
		Price:       request.Price,
		TypeTrashId: typeTrash.ID,
		BuyerId:     buyerId,
	})
	return trash, err
}

func (service TrashServiceImpl) GetAllTypeTrash() ([]domain.TypeTrash, error) {
	typeTrash, _ := service.TrashRepository.GetAllTypeTrash()
	return typeTrash, nil
}

func (service TrashServiceImpl) UpdateTypeTrash(request dto_trashes.TrashTypeRequestUpdate, trashId int) (*domain.TypeTrash, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return &domain.TypeTrash{}, err
	}

	findIdTyepTrash, _ := service.TrashRepository.GetTypeTrashById(trashId)
	if findIdTyepTrash == nil {
		return &domain.TypeTrash{}, helper.TypeTrashNotFound
	}

	findNameTypeTrash, _ := service.TrashRepository.FindTypeTrashByName(request.Name)
	if findNameTypeTrash != nil {
		return &domain.TypeTrash{}, helper.TrashAlreadyExist
	}

	trash, err := service.TrashRepository.UpdateTypeTrash(&domain.TypeTrash{
		Name: request.Name,
	}, trashId)
	return trash, err
}

func (service TrashServiceImpl) GetAllTrashByBuyerId(buyerId int) ([]domain.Trash, error) {
	trash, _ := service.TrashRepository.GetAllTrashByBuyerId(buyerId)
	return trash, nil
}

func (service TrashServiceImpl) GetAllTrash() ([]domain.Trash, error) {
	trash, _ := service.TrashRepository.GetAllTrash()
	return trash, nil
}

func (service TrashServiceImpl) GetTrashById(trashId, buyerId int) (*domain.Trash, error) {
	trash, _ := service.TrashRepository.GetTrashById(trashId, buyerId)

	fmt.Println("trash: ", trash)

	if trash == nil {
		return &domain.Trash{}, helper.TrashNotFound
	}
	return trash, nil
}

func (service TrashServiceImpl) UpdateTrash(request dto_trashes.TrashRequestUpdate, trashId, buyerId int) (*domain.Trash, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return &domain.Trash{}, err
	}

	findTrash, _ := service.TrashRepository.GetTrashById(trashId, buyerId)
	if findTrash == nil {
		return &domain.Trash{}, helper.TrashNotFound
	}

	findTypeTrash, _ := service.TrashRepository.FindTypeTrashByName(request.TypeTrash)
	if findTypeTrash == nil {
		return &domain.Trash{}, helper.TypeTrashNotFound
	}

	trash, err := service.TrashRepository.UpdateTrash(&domain.Trash{
		Name:  request.Name,
		Price: request.Price,
		TypeTrash: domain.TypeTrash{
			Name: request.Name,
		},
	}, trashId)
	return trash, err
}

func (service TrashServiceImpl) DelteTrash(trashId, buyerId int) (*domain.Trash, error) {
	findTrash, _ := service.TrashRepository.GetTrashById(trashId, buyerId)
	if findTrash == nil {
		return &domain.Trash{}, helper.TrashNotFound
	}

	_, error := service.TrashRepository.DelteTrash(trashId, buyerId)
	return nil, error
}

func (service TrashServiceImpl) GetListTrashPagination(page, pageSize int, nameTrash, typeTrash string, buyerName string) ([]*domain.Trash, error) {
	trash, _ := service.TrashRepository.GetListTrashPagination(page, pageSize, nameTrash, typeTrash, buyerName)
	return trash, nil
}
