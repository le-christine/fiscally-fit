package services

import (
	"github.com/le-christine/fiscally-fit/app/models"
	"reflect"
	"testing"
)

func TestCalculateTaxes(t *testing.T) {
	taxService := NewTaxService(models.AllTaxBrackets)

	calculatedTaxes, _ := taxService.CalculateTaxes(100000, 2020)

	expectedTaxes := map[string]interface{}{
		"totalTaxes":       "$17991.78",
		"effectiveTaxRate": 18,
		"taxesByBracket": map[string]string{
			"Tax band: $0.00 to $48535.00":      "$48535.00",
			"Tax band: $48535.00 to $97069.00":  "$48534.00",
			"Tax band: $97069.00 to $150473.00": "$2931.00",
		},
	}

	if reflect.DeepEqual(calculatedTaxes, expectedTaxes) {
		t.Errorf("calculatedTaxes = %v; want %v", calculatedTaxes, expectedTaxes)

	}
}
