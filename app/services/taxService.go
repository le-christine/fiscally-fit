package services

import (
	"fiscally-fit/app/models"
	"fmt"
	"math"
)

// Responsible for retrieving tax brackets from the mock API and calculating the taxes owed.

type TaxService struct {
	TaxBrackets map[int]models.TaxBrackets
}

func NewTaxService(brackets map[int]models.TaxBrackets) *TaxService {
	return &TaxService{
		TaxBrackets: brackets,
	}
}

// CalculateTaxes calculates and displays the total taxes owed for the salary, displays the amount of taxes owed per band, and displays the effective rate.
func (s *TaxService) CalculateTaxes(income float64, year int) (map[string]interface{}, error) {

	// Vars to return
	var totalTax, prev, max float64
	taxesByBracket := make(map[string]string)

	//Find the tax brackets for the given year
	brackets, ok := s.TaxBrackets[year]
	if !ok {
		return nil, fmt.Errorf("no tax brackets found for year %d", year)
	}

	for _, bracket := range brackets.TaxBrackets {

		if bracket.Max == nil {
			max = math.Inf(1)
		} else {
			max = *bracket.Max
		}

		if income <= max || max == 0 {
			taxableAmount := math.Min(income-bracket.Min, max-bracket.Min)
			taxesByBracket[fmt.Sprintf("Tax band: $%.2f to $%.2f", bracket.Min, max)] = fmt.Sprintf("$%.2f", taxableAmount)
			totalTax += taxableAmount * bracket.Rate
			break
		} else {
			taxableAmount := max - prev
			taxesByBracket[fmt.Sprintf("Tax band: $%.2f to $%.2f", bracket.Min, max)] = fmt.Sprintf("$%.2f", taxableAmount)
			totalTax += taxableAmount * bracket.Rate
			prev = max
		}
	}

	var taxRate float64
	if income == 0 {
		taxRate = 0
	} else {
		taxRate = totalTax / income * 100
	}

	response := map[string]interface{}{
		"totalTaxes":       fmt.Sprintf("$%.2f", totalTax),
		"effectiveTaxRate": math.Round(taxRate),
		"taxesByBracket":   taxesByBracket,
	}

	return response, nil
}
