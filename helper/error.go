package helper

import (
	"errors"
)

var (
	EmailAlreadyExist     = errors.New("email already exist")
	InvalidLogin          = errors.New("email or password wrong")
	TypeTrashAlreadyExist = errors.New("name type trash already exist")
	TrashAlreadyExist     = errors.New("name trash already exist")
	TypeTrashNotFound     = errors.New("type trash not found")
	TrashNotFound         = errors.New("trash not found")
	DataNotFound          = errors.New("data not found")
	ErrorIdParam          = errors.New("invalid id param")
)
