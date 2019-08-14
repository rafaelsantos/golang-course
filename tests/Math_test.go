package tests

import "testing"

const error = "Expected value %v, but result was %v"

func TestAverage(t *testing.T) {
	expectedValue := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(error, expectedValue, value)
	}
}
