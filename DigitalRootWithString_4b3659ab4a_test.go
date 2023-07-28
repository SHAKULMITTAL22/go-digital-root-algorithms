// Test generated by RoostGPT for test roost-test using AI Type Open Source AI and AI Model meta-llama/Llama-2-13b-chat

package test_digital_root

import (
	"fmt"
	"strconv"
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
			name:     "multi digit",
			input:    123,
			expected: 6,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := DigitalRootWithString(tc.input)
			if result!= tc.expected {
				t.Errorf("got %v, wanted %v", result, tc.expected)
			}
		})
	}
}

func TestDigitalRootWithString_edgeCases(t *testing.T) {
	testCases := []struct {
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
			expected: -7,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := DigitalRootWithString(tc.input)
			if result!= tc.expected {
				t.Errorf("got %v, wanted %v", result, tc.expected)
			}
		})
	}
}
