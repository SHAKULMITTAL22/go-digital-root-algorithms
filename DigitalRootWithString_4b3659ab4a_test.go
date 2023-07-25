f'''[INST] <<SYS>>Write code without any explanation, and write code in code block<</SYS>>[/INST]'''

package test

import (
	"testing"
)

func TestDigitalRootWithString(t *testing.T) {

	testCases := []struct {
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
			name:     "two digits",
			input:    123,
			expected: 6,
		},
		{
			name:     "three digits",
			input:    456,
			expected: 9,
		},
		{
			name:     "four digits",
			input:    7890,
			expected: 9,
		},
		{
			name:     "negative number",
			input:    -123,
			expected: -6,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := DigitalRootWithString(tc.input)
			if result!= tc.expected {
				t.Errorf("got %v, expected %v", result, tc.expected)
			}
		})
	}
}

func TestDigitalRootWithString_cornerCases(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		result := DigitalRootWithString("")
		if result!= 0 {
			t.Errorf("got %v, expected 0", result)
		}
	})

	t.Run("only digit", func(t *testing.T) {
		result := DigitalRootWithString("5")
		if result!= 5 {
			t.Errorf("got %v, expected 5", result)
		}
	})
}