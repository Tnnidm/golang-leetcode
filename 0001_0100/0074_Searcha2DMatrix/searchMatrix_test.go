package searchMatrix

import (
	"reflect"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	cases := []struct {
		name   string
		input1 [][]int
		input2 int
		expect bool
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := searchMatrix(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v and %v", c.expect, ret, c.input1, c.input2)
			}
		})
	}
}
