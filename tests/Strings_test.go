package tests

import (
	"strings"
	"testing"
)

const stringError = "%s (part: %s) - index: expected (%d) <> found (%d)"

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Cod3r is show", "Cod3r", 0},
		{"", "", 0},
		{"Oops", "oops", -1},
		{"Rafael", "f", 2},
	}

	for _, test := range tests {
		t.Logf("= %v", test)
		actual := strings.Index(test.text, test.part)

		if actual != test.expected {
			t.Errorf(stringError, test.text, test.part, test.expected, actual)
		}
	}
}
