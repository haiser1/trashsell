package repository

import "mini-project/model/domain"

type TrashRepository interface {
	FindTypeTrashByName(name string) (*domain.TypeTrash, error)
	FindTrashByName(name string, buyerId int) (*domain.Trash, error)
	GetTypeTrashById(id int) (*domain.TypeTrash, error)
	CreateTypeTrash(typeTrash *domain.TypeTrash) (*domain.TypeTrash, error)
	CreateTrash(trash *domain.Trash) (*domain.Trash, error)
	GetAllTypeTrash() ([]domain.TypeTrash, error)
	UpdateTypeTrash(typeTrash *domain.TypeTrash, id int) (*domain.TypeTrash, error)
	GetAllTrashByBuyerId(buyerId int) ([]domain.Trash, error)
	GetAllTrash() ([]domain.Trash, error)
	GetTrashById(id, buyerId int) (*domain.Trash, error)
	UpdateTrash(trash *domain.Trash, id int) (*domain.Trash, error)
	DelteTrash(trashId, buyerId int) (*domain.Trash, error)
	GetListTrashPagination(page int, pageSize int, nameTrash string, typeTrash string, buyerName string) ([]*domain.Trash, error)
}
