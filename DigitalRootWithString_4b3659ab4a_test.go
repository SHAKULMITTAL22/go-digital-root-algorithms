

Here is the corrected version of the Go test case:
```
const (
	expectedResInt = 6
)

func TestDigitalRootWithString(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{123, expectedResInt},
		{-456, -2},
		{0, 0},
		{9876543210, 9},
	}

	for _, test := range tests {
		res := DigitalRootWithString(test.input)
		if res!= test.want {
			t.Errorf("TestDigitalRootWithString failed for input %v, expected %d but got %d", test.input, test.want, res)
		}
	}
}

func TestMain(m *testing.M) {
	err := m.Run()
	if err!= nil {
		panic(err)
	}
}

func main() {
	testing.InitDiscover(&testing.Config{})
	testing.Run(New(Suite), &testing.M)
}

const Suite = &testing.Suite{
	&testing.Test{
		Name: "TestDigitalRootWithString",
		Fn: TestDigitalRootWithString,
	},
}
```
The changes I made were:

* Removed the `var` declarations from the test functions, as they are not needed.
* Changed the type of `tests` to a slice of structs, where each struct contains an `input` field and a `want` field.
* In the loop that runs the tests, used the `res` variable instead of `testCase` to access the result of the `DigitalRootWithString` function.
* Added a error message to the test function for each input, so that the test will fail with a more informative message when the output does not match the expected value.