package setZeroes

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect [][]int
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setZeroes(c.input)
			if !reflect.DeepEqual(c.input, c.expect) {
				t.Fatalf("expect: %v, but got: %v", c.expect, c.input)
			}
		})
	}
}
