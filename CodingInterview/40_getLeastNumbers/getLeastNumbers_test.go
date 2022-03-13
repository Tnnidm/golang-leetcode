package getLeastNumbers

import (
	"reflect"
	"testing"
)

func TestGetLeast(t *testing.T) {
	cases := []struct {
		name   string
		arr    []int
		k      int
		expect []int
	}{
		{"case 1", []int{3, 2, 1}, 2, []int{1, 2}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := getLeastNumbers(c.arr, c.k)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with: %v and %v", c.expect, ret, c.arr, c.k)
			}
		})
	}
}
