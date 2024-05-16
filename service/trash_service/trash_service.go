package service

import (
	dto_trashes "mini-project/dto/trashes"
	"mini-project/model/domain"
)

type TrashService interface {
	CreateTypeTrash(request dto_trashes.BuyerCreateTypeTrash) (*domain.TypeTrash, error)
	CreateTrash(request dto_trashes.TrashRequestCreate, buyerId int) (*domain.Trash, error)
	GetAllTypeTrash() ([]domain.TypeTrash, error)
	UpdateTypeTrash(request dto_trashes.TrashTypeRequestUpdate, trashId int) (*domain.TypeTrash, error)
	GetAllTrashByBuyerId(buyerId int) ([]domain.Trash, error)
	GetAllTrash() ([]domain.Trash, error)
	GetTrashById(trashId, buyerId int) (*domain.Trash, error)
	UpdateTrash(request dto_trashes.TrashRequestUpdate, trashId, buyerId int) (*domain.Trash, error)
	DelteTrash(trashId, buyerId int) (*domain.Trash, error)
	GetListTrashPagination(page, pageSize int, nameTrash, typeTrash string, buyerName string) ([]*domain.Trash, error)
}
