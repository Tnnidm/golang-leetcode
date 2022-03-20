package dicesProbability

import (
	"reflect"
	"testing"
)

func Test_dicesProbability(t *testing.T) {
	cases := []struct {
		name   string
		input  int
		expect []float64
	}{
		{"case 1", 1, []float64{1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6, 1.0 / 6}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := dicesProbability(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
