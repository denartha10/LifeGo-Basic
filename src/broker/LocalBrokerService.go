package broker

import (
	"strings"

	"github.com/denartha10/lifeGo-basic/core"
)

// Implementation of thebroker service that uses the Service Registry
type LocalBrokerService struct {
	core.BrokerService
}

func (l *LocalBrokerService) GetQuotations(info core.ClientInfo) []core.Quotation {
	quotations := []core.Quotation{}

	for _, t := range core.ServiceRegistryList() {
		if strings.HasPrefix(t, "qs-") {
			service := core.ServiceRegistryLookup(t)
			if qs, ok := service.(core.QuotationService); ok {
				quotations = append(quotations, qs.GenerateQuotation(info))
			}
		}
	}

	return quotations
}
