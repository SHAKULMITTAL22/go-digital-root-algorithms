package digital_root_test

import (
	"testing"
	"strconv"
)

func TestDigitalRootWithString(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "single digit",
			input:     5,
			expected: 5,
		},
		{
			name:     "multi digit",
			input:    123,
			expected: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DigitalRootWithString(test.input)
			if result!= test.expected {
				t.Errorf("got %d, expected %d", result, test.expected)
			}
		})
	}
}

func TestDigitalRootWithStringEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "-1",
			input:    -1,
			expected: -1,
		},
		{
			name:     "0",
			input:      0,
			expected:  0,
		},
		{
			name:     "negative",
			input:    -45,
			expected: -6,
		},
		{
			name:     "large positive",
			input:     999,
			expected: 9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DigitalRootWithString(test.input)
			if result!= test.expected {
				t.Errorf("got %d, expected %d", result, test.expected)
			}
		})
	}
}
