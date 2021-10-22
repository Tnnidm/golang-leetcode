package containsNearbyDuplicate

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//测试用例
	cases := []struct {
		name   string
		input1 []int
		input2 int
		expect bool
	}{
		{"test case 1", []int{1, 2, 3, 1}, 3, true},
		{"test case 2", []int{1, 0, 1, 1}, 1, true},
		{"test case 3", []int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := containsNearbyDuplicate(c.input1, c.input2)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expected: %v. but got: %v, with inputs: %v, %v", c.expect, ret, c.input1, c.input2)
				}
			},
		)
	}
}
