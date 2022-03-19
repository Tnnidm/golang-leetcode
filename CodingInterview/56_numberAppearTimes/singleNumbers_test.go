package numberAppearTimes

import (
	"reflect"
	"testing"
)

func Test_singleNumbers(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"case 1", []int{4, 1, 4, 6}, []int{1, 6}},
		{"case 2", []int{0, 1}, []int{0, 1}},
		{"case 2", []int{1, 2, 10, 4, 1, 4, 3, 3}, []int{2, 10}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := singleNumbers(c.input)
			if !reflect.DeepEqual(ret, c.expect) && !reflect.DeepEqual(ret, exchange(c.expect)) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}

func exchange(arr []int) []int {
	arr[0], arr[1] = arr[1], arr[0]
	return arr
}
