package broker

import (
	"log"
	"strings"

	"github.com/denartha10/lifeGo-basic/auldfellas"
	"github.com/denartha10/lifeGo-basic/core"
	"github.com/denartha10/lifeGo-basic/dodgygeezers"
	"github.com/denartha10/lifeGo-basic/girlsallowed"
)

// Implementation of thebroker service that uses the Service Registry
type LocalBrokerService struct {
	core.BrokerService
}

func NewLocalBroker() LocalBrokerService {
	return LocalBrokerService{}
}

func (l *LocalBrokerService) GetQuotations(info *core.ClientInfo) []*core.Quotation {
	quotations := []*core.Quotation{}

	for _, t := range core.ServiceRegistryList() {
		if strings.HasPrefix(t, "qs-") {
			service := core.ServiceRegistryLookup(t)

			switch v := service.(type) {
			case girlsallowed.GAQService:
				quotations = append(quotations, v.GenerateQuotation(info))
			case auldfellas.AFQService:
				quotations = append(quotations, v.GenerateQuotation(info))
			case dodgygeezers.DGQService:
				quotations = append(quotations, v.GenerateQuotation(info))
			default:
				log.Printf("Service is not a QuotationService: %v\n", v)
			}

		}
	}

	return quotations
}
