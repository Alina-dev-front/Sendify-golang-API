package counters

import (
	"github.com/shopspring/decimal"
)

//SetFinalPrice counts shipment price based on weight class and sender country code
func SetFinalPrice(countryCode string, weight float64) string {
	priceDecimal := decimal.NewFromInt(CountWeightClassPrice(weight))

	var finalPrice decimal.Decimal

	if belongToListNordicRegion(countryCode) == true {
		finalPrice = priceDecimal
	} else if belongToListEU(countryCode) == true {
		finalPrice = priceDecimal.Mul(decimal.NewFromFloat(1.5))
	} else {
		finalPrice = priceDecimal.Mul(decimal.NewFromFloat(2.5))
	}
	return finalPrice.String()
}

//CountWeightClassPrice price depends on weight class of the package
func CountWeightClassPrice(weight float64) int64 {
	var weightClassPrice int64 = 0

	switch {
	case weight < 10:
		weightClassPrice = 100
	case weight >= 10 && weight < 25:
		weightClassPrice = 300
	case weight >= 25 && weight < 50:
		weightClassPrice = 500
	case weight >= 50 && weight <= 1000:
		weightClassPrice = 2000
	default:
		panic("Invalid weight. Weight should not be more than 1000")
	}

	return weightClassPrice
}

func belongToListNordicRegion(lookup string) bool {
	listNordic := []string{"SE", "DK", "NO", "FI"}
	for _, val := range listNordic {
		if val == lookup {
			return true
		}
	}
	return false
}

func belongToListEU(lookup string) bool {
	listEU := []string{"AT", "BE", "BG", "HR", "CY", "CZ", "EE", "FR", "DE", "GR", "HU", "IE", "IT",
		"LV", "LT", "LU", "MT", "NL", "PL", "PT", "RO", "SK", "SI", "ES"}
	for _, val := range listEU {
		if val == lookup {
			return true
		}
	}
	return false
}
