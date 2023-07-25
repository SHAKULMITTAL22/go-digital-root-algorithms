f'''[INST] <<SYS>>Write code without any explanation, and write code in code block<</SYS>>[/INST]'''

package test_digitalroot

import (
	"testing"
)

func TestDigitalRootWithModulo(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{"single digit", 5, 5},
		{"two digits", 23, 7},
		{"three digits", 987, 9},
		{"negative", -4, -4},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DigitalRootWithModulo(tt.arg)
			if got!= tt.want {
				t.Errorf("DigitalRootWithModulo(%d): expected %d, got %d", tt.arg, tt.want, got)
			}
		})
	}
}

func TestDigitalRootWithModulo_cornerCases(t *testing.T) {
	cases := []struct {
		name string
		arg  int
		want int
	}{
		{"one digit", 5, 5},
		{"zero", 0, 0},
		{"negative", -4, -4},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			got := DigitalRootWithModulo(cc.arg)
			if got!= cc.want {
				t.Errorf("DigitalRootWithModulo(%d): expected %d, got %d", cc.arg, cc.want, got)
			}
		})
	}
}

func TestDigitalRootWithModulo_recursiveCall(t *testing.T) {
	const want = 9
	arg := 123

	got := DigitalRootWithModulo(arg)

	if got!= want {
		t.Errorf("DigitalRootWithModulo(%d): expected %d, got %d", arg, want, got)
	}
}