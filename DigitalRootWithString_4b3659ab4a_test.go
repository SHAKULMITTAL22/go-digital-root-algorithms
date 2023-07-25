f'''[INST] <<SYS>>Write code without any explanation, and write code in code block<</SYS>>[/INST]'''

package test

import (
	"fmt"
	"testing"
)

func TestDigitalRootWithString(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"empty", 0, 0},
		{"one digit", 5, 5},
		{"two digits", 123, 6},
		{"three digits", 4567, 8},
		{"four digits", 90123, 9},
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

func TestDigitalRootWithString_cornerCases(t *testing.T) {
	cases := map[int]bool{
		0: true, // empty string
		1: true, // one digit string
	}

	for input, expect := range cases {
		t.Run(fmt.Sprintf("%d", input), func(t *testing.T) {
			result := DigitalRootWithString(input)
			if!expect {
				t.Errorf("got %d, expected fail", result)
			}
		})
	}
}

func TestDigitalRootWithString_recursiveCall(t *testing.T) {
	const input = 123456789
	expected := 9

	t.Run("large input", func(t *testing.T) {
		result := DigitalRootWithString(input)
		if result!= expected {
			t.Errorf("got %d, expected %d", result, expected)
		}
	})
}