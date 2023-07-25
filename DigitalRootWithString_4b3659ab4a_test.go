package digital_root_test

import (
	"testing"
)

func TestDigitalRootWithString(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"single digit", 5, 5},
		{"two digits", 123, 6},
		{"three digits", 456, 9},
		{"four digits", 7890, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DigitalRootWithString(tt.input)
			if result!= tt.expected {
				t.Errorf("got %d, expected %d", result, tt.expected)
			}
		})
	}
}
