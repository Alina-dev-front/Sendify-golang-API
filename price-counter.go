package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func countWeightClassPrice(weight float64) float64 {
	weightClassPrice := 0.0

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
		fmt.Println("Invalid")
	}

	return weightClassPrice
}

func setFinalPrice(countryCode string, weight float64) string {
	weightClassPrice := countWeightClassPrice(weight)

	finalPrice := decimal.NewFromFloat(0.0)
	priceDecimal := decimal.NewFromFloat(weightClassPrice)

	if belongToListNordicRegion(countryCode) == true {
		finalPrice = priceDecimal
	} else if belongToListEU(countryCode) == true {
		finalPrice = priceDecimal.Mul(decimal.NewFromFloat(1.5))
	} else {
		finalPrice = priceDecimal.Mul(decimal.NewFromFloat(2.5))
	}
	return finalPrice.String()
}

func belongToListNordicRegion(lookup string) bool {
	list := []string{
		"SE",
		"DK",
		"NO",
		"FI",
	}
	for _, val := range list {
		if val == lookup {
			return true
		}
	}
	return false
}

func belongToListEU(lookup string) bool {
	list := []string{
		"AD",
		"AT",
		"BE",
	}
	for _, val := range list {
		if val == lookup {
			return true
		}
	}
	return false
}
