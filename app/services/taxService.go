package services

import (
	"fiscally-fit/app/models"
)

// Responsible for retrieving tax brackets from the mock API and calculating the taxes owed.

type TaxService struct {
	TaxBrackets models.TaxBracket // Use the updated struct name here
}

func NewTaxService(brackets models.TaxBracket) *TaxService { // Use the updated struct name here
	return &TaxService{
		TaxBrackets: brackets,
	}
}

// CalculateTaxes calculates and displays the total taxes owed for the salary, displays the amount of taxes owed per band, and displays the effective rate.
func (s *TaxService) CalculateTaxes(salary float64, year int) (float64, error) {
	// Find the tax brackets for the given year
	//brackets, ok := s.TaxBrackets[year]
	//if !ok {
	//	return 0, fmt.Errorf("no tax brackets found for year %d", year)
	//}
	//
	//// Calculate taxes for each bracket and add them up
	//var totalTax float64
	//for _, bracket := range brackets {
	//	if salary > float64(bracket.Max) {
	//		totalTax += float64(bracket.Max-bracket.Min) * bracket.Rate
	//		fmt.Printf("Tax band: $%d to $%d. Tax owed: $%.2f\n", bracket.Min, bracket.Max, float64(bracket.Max-bracket.Min)*bracket.Rate)
	//	} else if salary > float64(bracket.Min) {
	//		totalTax += (salary - float64(bracket.Min)) * bracket.Rate
	//		fmt.Printf("Tax band: $%d to $%.0f. Tax owed: $%.2f\n", bracket.Min, salary, (salary-float64(bracket.Min))*bracket.Rate)
	//	}
	//}
	//
	//fmt.Printf("Total tax owed: $%.2f\n", totalTax)
	//fmt.Printf("Effective tax rate: %.2f%%\n", totalTax/salary*100)

	return 0, nil
}
