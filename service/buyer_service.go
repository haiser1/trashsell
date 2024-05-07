package service

import "mini-project/model/base"

type BuyerService interface {
	RegisterBuyer(request base.BayerRequestRegister) base.BuyerResponseRegister
}
