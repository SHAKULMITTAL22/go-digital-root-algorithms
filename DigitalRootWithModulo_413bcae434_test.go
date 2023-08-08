// Test generated by RoostGPT for test roost-test using AI Type Open AI and AI Model gpt-4

package main

import (
	"testing"
)

func TestDigitalRootWithModulo_413bcae434(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{"Single Digit", 5, 5},
		{"Double Digit", 45, 9},
		{"Triple Digit", 123, 6},
		{"Negative Number", -123, 0},
		{"Zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigitalRootWithModulo(tt.arg); got != tt.want {
				t.Errorf("DigitalRootWithModulo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func DigitalRootWithModulo(number int) int {

	//corner case - check if number with one digit.
	if (10 - number > 0) {
		return number
	}

	// once here - means number has more than 1 digit.
	// variables to store computation - default to 0 as integer.
	var sum, remainder int
	// perform euclidian division until quotient which
	// is the number equals 0.
	for number != 0 { 
		// get the rest and add it to the sum.
		remainder = number % 10  
		sum = sum + remainder
	  	// result of euclidean division by 10 will be used for next iteration.
	 	number = number / 10  
	}

	// recursively perform on the sum value.
	return DigitalRootWithModulo(sum)
}
