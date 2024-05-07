package helper

import "errors"

var (
	BuyerAlreadyExist = errors.New("buyer already exist")
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
