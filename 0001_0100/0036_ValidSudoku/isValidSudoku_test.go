package isValidSudoku

import (
	"reflect"
	"testing"
)

func Test_isValidSudoku(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]byte
		expect bool
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isValidSudoku(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input)
			}
		})
	}
}
