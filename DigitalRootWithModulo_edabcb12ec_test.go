package main

import (
	"fmt"
	"testing"
)

func TestDigitalRootWithModulo_edabcb12ec(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{123, 6},
		{1234, 1},
		{12345, 6},
		{123456, 9},
		{1234567, 1},
		{12345678, 9},
		{123456789, 6},
		{1234567890, 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input=%d", tc.input), func(t *testing.T) {
			actual := DigitalRootWithModulo(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
