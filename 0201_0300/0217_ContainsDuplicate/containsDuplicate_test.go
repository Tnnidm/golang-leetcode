package containsDuplicate

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//测试用例
	cases := []struct {
		name   string
		input  []int
		expect bool
	}{
		{"test case 1", []int{1, 2, 3, 1}, true},
		{"test case 2", []int{1, 2, 3, 4}, false},
		{"test case 3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := containsDuplicate(c.input)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expected: %v. but got: %v, with inputs: %v", c.expect, ret, c.input)
				}
			},
		)
	}
}
