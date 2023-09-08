package main

import (
	"fmt"
	"testing"
)

func TestDigitalRootWithString_3b1c30d484(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{123, 6},
		{493193, 2},
		{123456789, 4},
		{987654321, 9},
		{0, 0},
		{-123, 6}, // negative number
		{1234567890, 1}, // number with more than 9 digits
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input=%d", tc.input), func(t *testing.T) {
			actual := DigitalRootWithString(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}
