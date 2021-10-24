package summaryRanges

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	//测试用例
	cases := []struct {
		name   string
		input  []int
		expect []string
	}{
		{"test case 1", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"test case 2", []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{"test case 3", []int{}, []string{}},
		{"test case 3", []int{0}, []string{"0"}},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := summaryRanges(c.input)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expected: %v. but got: %v, with inputs: %v", c.expect, ret, c.input)
				}
			},
		)
	}
}
