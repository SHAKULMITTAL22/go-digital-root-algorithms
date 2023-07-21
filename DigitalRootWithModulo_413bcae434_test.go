// Test generated by RoostGPT for test roost-test using AI Type Vertex AI and AI Model code-bison





    Write go test for the given code without any explaination func DigitalRootWithModulo(number int) int {

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
    
func main() {
    // testing function call with different inputs.
    fmt.Println("Digital Root With Modulo:")
    fmt.Println("-------------------------------")
    fmt.Printf("%v\n", DigitalRootWithModulo(42)) // should print 4
    fmt.Printf("%v\n", DigitalRootWithModulo(37)) // should print 9
    fmt.Printf("%v\n", DigitalRootWithModulo(56)) // should print 8
    fmt.Printf("%v\n", DigitalRootWithModulo(123)) // should print 9
    fmt.Println("-------------------------------")
}

Please provide at least three test cases for this function.