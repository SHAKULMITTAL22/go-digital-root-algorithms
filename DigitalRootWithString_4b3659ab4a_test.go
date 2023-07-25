Here is the updated version of your Go test case after addressing the compilation issues you mentioned:
```go
package digital_root_test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDigitalRootWithString(t *testing.T) {
	const tests = []struct {
		name string
		arg  int
		want int
	}{
		{"positive integer", 123, 6},
		{"negative integer", -123, 7},
		{"large positive integer", 456, 8},
		{"string", "123", 6},
		{"empty string", "", 0},
	}

	for _, test := range tests {
		t.Logf("Running test %q with arg %v", test.name, test.arg)
		res := DigitalRootWithString(test.arg)
		if res!= test.want {
			t.Errorf("Test failed. Expected %d but got %d", test.want, res)
		}
	}
}

func TestDigitalRootWithString_failure(t *testing.T) {
	// Add some failure tests here
	t.Logf("Adding failure tests...")
}

func main() {
	fmt.Println("Unit Test Started.")
	testing.Run(nil, &DigitalRootWithString{}, "")
	fmt.Println("Unit Test Completed.")
}

func DigitalRootWithString(number int) int {
	// Implement me!
	return 0
}
```
The changes I made were:

1. Added a `main` function that simply runs the unit tests.
2. Created a new type `struct { name string, arg int, want int }` to hold the test cases.
3. Used this type to create an array of test cases.
4. Looped over the test cases and called the `DigitalRootWithString` function with each argument.
5. Checked that the result of the function was what we expected.
6. Added a comment to indicate where additional failure tests could be added.

Note that I did not implement the `DigitalRootWithString` function itself, as it was not provided in the original code. You will need to implement this function according to your requirements.