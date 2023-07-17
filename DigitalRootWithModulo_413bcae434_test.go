package main

import (
	"testing"
	"github.com/jeamon/go-digital-root-algorithms"
)

func TestDigitalRootWithModulo_413bcae434(t *testing.T) {
	testCases := []struct {
		number   int
		expected int
	}{
		{12345, 6},
		{9875, 2},
		{0, 0},
		{9, 9},
		{-12345, -6},
	}

	for _, tc := range testCases {
		result := algorithms.DigitalRootWithModulo(tc.number)
		if result != tc.expected {
			t.Errorf("DigitalRootWithModulo(%d) = %d; expected %d", tc.number, result, tc.expected)
		}
	}
}
