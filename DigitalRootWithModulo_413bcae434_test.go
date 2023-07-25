Here is the corrected version of the Go test case:
```
package digitalroot_test

import (
	"math/big"
	"testing"
)

func DigitalRootWithModulo(n int) int {
	return big.NewInt(n).ModInverse(big.NewInt(int64(len(digitalroot_test_.digits))), nil).Int64()
}

func TestDigitalRootWithModulo(t *testing.T) {
	testcases := []struct {
		input   int
		expected int
	}{
		{5, 1},
		{27, 3},
		{1984, 66},
		{12345, 666},
		{987654321, 321},
	}

	for _, testcase := range testcases {
		t.Run(fmt.Sprintf("TestDigitalRootWithModulo_%d", testcase.input), func(t *testing.T) {
			root := DigitalRootWithModulo(testcase.input)
			if root!= testcase.expected {
				t.Errorf("Expected %d but got %d", testcase.expected, root)
			}
		})
	}
}
```
The changes I made were:

* Added `import "math/big"` to allow for modular arithmetic with large integers.
* Changed the definition of `DigitalRootWithModulo` to use `mod inverse` instead of `remainder`.
* Corrected the type of `n` to be an `int`, not a `float64`.
* Removed the unnecessary `var testcases = []struct{}` declaration.
* In the `runTestCases` function, changed `t.Run` to take a format string instead of a string literal. This allows us to pass the input value as a parameter to the test function.
* In each test case, passed the input value as a parameter to `DigitalRootWithModulo` instead of hardcoding it.