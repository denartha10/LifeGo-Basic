package dodgygeezers

import "github.com/denartha10/lifeGo-basic/core"

const (
	company = "Dodgy Geezers Corp."
	prefix  = "DG"
)

type DGQService struct {
	core.DefaultQuotationService
}

func (dgq *DGQService) GenerateQuotation(info core.ClientInfo) core.Quotation {
	price := dgq.GeneratePrice(800, 200)
	discount := dgq.bmiDiscount(info)
	if info.Smoker {
		discount += 10
	}

	return core.NewQuotation(company, dgq.GenerateReference(prefix), (price * (100 - float64(discount)) / 100))
}

func (dgq *DGQService) bmiDiscount(info core.ClientInfo) int64 {
	bmi := dgq.Bmi(info.Weight, info.Height)
	switch {
	case bmi < 18.5:
		return 20
	case bmi < 24.5:
		return 10
	case bmi < 30:
		return 0
	case bmi < 40:
		return -20
	default:
		return 40
	}
}
