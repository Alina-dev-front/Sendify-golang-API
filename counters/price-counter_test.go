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
		{input: 918.5, expected: 3000},
	}

	for _, s := range scenarios {
		got := counters.CountWeightClassPrice(s.input)
		if got != s.expected {
			t.Errorf("Did not get expected result fot input '%v'. Got: '%v', expected: '%v'", s.input, got, s.expected)
		}
	}
}
