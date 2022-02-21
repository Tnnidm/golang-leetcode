package exist

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		input1 [][]byte
		input2 string
		expect bool
	}{
		{"case 1", [][]byte{{'a', 'b', 't', 'g'}, {'c', 'f', 'c', 's'}, {'j', 'd', 'e', 'h'}}, "bfce", true},
		{"case 2", [][]byte{{'a', 'b', 't', 'g'}, {'c', 'f', 'c', 's'}, {'j', 'd', 'e', 'h'}}, "abfb", false},
		{"case 3", [][]byte{{'a', 'b', 't', 'g'}}, "abt", true},
		{"case 4", [][]byte{{'a', 'b', 't', 'g'}}, "abf", false},
		{"case 5", [][]byte{{'a'}}, "a", true},
		{"case 6", [][]byte{{'a', 'b'}}, "ba", true},
	}
	for _, c := range cases {
		t.Run(
			c.name,
			func(t *testing.T) {
				ret := exist(c.input1, c.input2)
				if !reflect.DeepEqual(ret, c.expect) {
					t.Fatalf("expect: %v, but got: %v, with input: %v and %v", c.expect, ret, c.input1, c.input2)
				}
			},
		)
	}
}
