package evalRPN

import (
	"reflect"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect int
	}{
		{"case 1", []string{"2", "1", "+", "3", "*"}, 9},
		{"case 2", []string{"4", "13", "5", "/", "+"}, 6},
		{"case 3", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := evalRPN(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("Expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
