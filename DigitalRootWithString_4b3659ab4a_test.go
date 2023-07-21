// Test generated by RoostGPT for test roost-test using AI Type Vertex AI and AI Model code-bison





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
    Write go test for the code
    
    ```go
    package main

    import "testing"

    func TestDigitalRoot(t *testing.T) {
        // sample inputs
        tests := []struct {
            name   string
            input int
            want  int
        }{
            {"test-case-1", 123, 6},
            {"test-case-2", 456, 9},
            {"test-case-3", 789, 9},
        }

        for _, tt := range tests {
            got := DigitalRootWithString(tt.input)
            if got!= tt.want {
                t.Errorf("got %d instead of %d for input %d", got, tt.want, tt.input)
            }
        }
    }

    func main() {}
    ```
    Go through the test cases and write the corresponding assertions in the `TestDigitalRoot` function.

Hint: You can use the built-in `assert` statement to make your life easier.