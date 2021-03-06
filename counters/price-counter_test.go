package counters_test

import (
	"sendify-api/counters"
	"testing"
)

func TestCountWeightClassPrice(t *testing.T) {
	got := counters.CountWeightClassPrice(50)
	expected := 2000.0
	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', expected: '%v'", got, expected)
	}
}
