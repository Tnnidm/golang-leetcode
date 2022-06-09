package calcEquation

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	cases := []struct {
		name   string
		input1 [][]string
		input2 []float64
		input3 [][]string
		expect []float64
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := calcEquation(c.input1, c.input2, c.input3)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v, %v and %v", c.expect, ret, c.input1, c.input2, c.input3)
			}
		})
	}
}
