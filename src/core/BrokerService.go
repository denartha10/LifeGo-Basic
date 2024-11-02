package core

// Interface for defining the behaviours of the broker service
type BrokerService interface {
	GetQuotations(info *ClientInfo) []*Quotation
}
