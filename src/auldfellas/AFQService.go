package auldfellas

import (
	"github.com/denartha10/lifeGo-basic/core"
)

const (
	company = "Auld Fellas Ltd."
	prefix  = "AF"
)

type AFQService struct {
	core.QuotationService
	core.DefaultQuotationService
	COMPANY string
	PREFIX  string
}

func NewAFQService() AFQService {
	return AFQService{
		COMPANY: company,
		PREFIX:  prefix,
	}
}

func (afq *AFQService) GenerateQuotation(info *core.ClientInfo) *core.Quotation {
	price := afq.GeneratePrice(600, 600)

	discount := 0.0

	if info.Gender == core.MALE {
		discount = 30
		if info.Age > 50 {
			discount += 10
		}
		if info.Age > 60 {
			discount += 10
		}
	} else {
		discount -= 20
	}

	return core.NewQuotation(company, afq.GenerateReference(prefix), (price * (100 - discount) / 100))
}
