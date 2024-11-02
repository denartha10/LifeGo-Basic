package girlsallowed

import "github.com/denartha10/lifeGo-basic/core"

const (
	company = "Girls Allowed Inc."
	prefix  = "GA"
)

type GAQService struct {
	core.QuotationService
	core.DefaultQuotationService
	PREFIX  string
	COMPANY string
}

func NewGAQService() GAQService {
	return GAQService{
		COMPANY: company,
		PREFIX:  prefix,
	}
}

func (g *GAQService) GenerateQuotation(info *core.ClientInfo) *core.Quotation {
	// create an initial quotation between 600 and 1000
	price := g.GeneratePrice(600, 400)

	// automatic 50% discount for being female
	var discount float64
	if info.Gender == core.FEMALE {
		discount = 50
	} else {
		discount = 30
	}

	// add a points discount
	discount += float64(g.bmiDiscount(info))

	// add a medical weighting
	discount += float64(g.medicalWeighting(info))

	return core.NewQuotation(
		company,
		g.GenerateReference(prefix),
		(price * (100 - discount) / 100),
	)
}

func (g *GAQService) bmiDiscount(info *core.ClientInfo) int64 {
	bmi := g.Bmi(info.Weight, info.Height)
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

func (g *GAQService) medicalWeighting(info *core.ClientInfo) int64 {
	var weighting int64 = 0
	if info.MedicalIssues {
		weighting -= 20
	}

	if info.Smoker {
		weighting -= 40
	}

	return weighting
}
