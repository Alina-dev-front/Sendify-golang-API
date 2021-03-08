package counters_test

import (
	"sendify-api/counters"
	"testing"
)

func TestCountWeightClassPriceTableDriven(t *testing.T) {
	scenarios := []struct {
		input    float64
		expected int64
	}{
		{input: 0.806, expected: 100},
		{input: 17, expected: 300},
		{input: 32.002, expected: 500},
		{input: 918.5, expected: 2000},
	}

	for _, s := range scenarios {
		got := counters.CountWeightClassPrice(s.input)
		if got != s.expected {
			t.Errorf("Did not get expected result fot input '%v'. Got: '%v', expected: '%v'", s.input, got, s.expected)
		}
	}
}

func TestSetFinalPrice(t *testing.T) {
	scenarios := []struct {
		inputCountryCode string
		inputWeight      float64
		expected         string
	}{
		{inputCountryCode: "FI", inputWeight: 1, expected: "100"},
		{inputCountryCode: "FR", inputWeight: 1, expected: "150"},
		{inputCountryCode: "RU", inputWeight: 1, expected: "250"},
		{inputCountryCode: "SE", inputWeight: 100, expected: "2000"},
		{inputCountryCode: "ES", inputWeight: 100, expected: "3000"},
		{inputCountryCode: "US", inputWeight: 100, expected: "5000"},
	}

	for _, s := range scenarios {
		got := counters.SetFinalPrice(s.inputCountryCode, s.inputWeight)
		if got != s.expected {
			t.Errorf("Did not get expected result fot input '%v'. Got: '%v', expected: '%v'", s.inputCountryCode, got, s.expected)
		}
	}
}
