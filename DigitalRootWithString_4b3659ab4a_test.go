// Test generated by RoostGPT for test roost-test using AI Type Open AI and AI Model gpt-4

package main

import (
	"testing"
	"strconv"
)

func TestDigitalRootWithString_4b3659ab4a(t *testing.T) {
	// Test Case 1: Testing with a positive number
	number := 123456
	expectedResult := 3
	result := DigitalRootWithString(number)
	if result != expectedResult {
		t.Errorf("Expected %d but received %d", expectedResult, result)
	}

	// Test Case 2: Testing with a single digit number
	number = 9
	expectedResult = 9
	result = DigitalRootWithString(number)
	if result != expectedResult {
		t.Errorf("Expected %d but received %d", expectedResult, result)
	}

	// Test Case 3: Testing with a number that is already a digital root
	number = 7
	expectedResult = 7
	result = DigitalRootWithString(number)
	if result != expectedResult {
		t.Errorf("Expected %d but received %d", expectedResult, result)
	}

	// Test Case 4: Testing with a large number
	number = 9876543210
	expectedResult = 9
	result = DigitalRootWithString(number)
	if result != expectedResult {
		t.Errorf("Expected %d but received %d", expectedResult, result)
	}

	// Test Case 5: Testing with zero
	number = 0
	expectedResult = 0
	result = DigitalRootWithString(number)
	if result != expectedResult {
		t.Errorf("Expected %d but received %d", expectedResult, result)
	}
}

func DigitalRootWithString(number int) int {

	// convert number from integer to string.
	s := strconv.Itoa(number)

	// corner case - if number of one digit.
	if len(s) == 1 {
		return number
	}
	
	// once here means - number has more than 1 digit inside.
	// variables to store conversion and to avoid redeclaring
	// new one for each iteration. 
	var sum, digit int
	
	// loop over the string and add each character to the sum.
	for _, c := range s {
		// c is a rune type so convert into string type.
		// then convert from string to integer.
		digit, _ = strconv.Atoi(string(c))
		// add the digit to the sum.
		sum = sum + digit
	}

	// recursively perform on the sum value.
	return DigitalRootWithString(sum)
} 
