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
		{"single digit", 5, 1},
		{"two digits", 23, 3},
		{"three digits", 467, 9},
		{"four digits", 1234, 8},
		{"five digits", 56789, 9},
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
		{"empty input", 0, 0},
		{"input equal to 10", 10, 1},
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