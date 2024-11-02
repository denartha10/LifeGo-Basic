package core

import (
	"fmt"
	"math/rand/v2"
)

type QuotationService interface {
	Service
	GenerateQuotation(c ClientInfo) Quotation
}

// type to store the Quotations returnes by Quotation services
type Quotation struct {
	Company   string
	Reference string
	Price     float64
}

func NewQuotation(company, reference string, price float64) Quotation {
	return Quotation{
		Company:   company,
		Reference: reference,
		Price:     price,
	}
}

// this is a copy of the default abstract class
// We don't have inheritance in go so this will be passes as a composition
type DefaultQuotationService struct {
	counter int64
	random  rand.Rand
}

func (qs *DefaultQuotationService) GenerateReference(prefix string) string {
	ref := prefix
	var length int64 = 100000
	for length > 1000 {
		if qs.counter/length == 0 {
			ref += "0"
		}
		length = length / 10
	}
	qs.counter += 1
	return fmt.Sprintf("%s%d", ref, qs.counter)
}

func (qs *DefaultQuotationService) GeneratePrice(m float64, r int64) float64 {
	return m + float64(r)*qs.random.Float64()
}

func (qs *DefaultQuotationService) Bmi(weight float64, height float64) float64 {
	return weight / (height * height)
}
