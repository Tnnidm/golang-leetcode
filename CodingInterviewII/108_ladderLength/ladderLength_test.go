package ladderLength

import (
	"reflect"
	"testing"
)

func Test_ladderLength(t *testing.T) {
	cases := []struct {
		name   string
		input1 string
		input2 string
		input3 []string
		expect int
	}{
		{"case 1", "hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := ladderLength(c.input1, c.input2, c.input3)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v, %v and %v", c.expect, ret, c.input1, c.input2, c.input3)
			}
		})
	}
}
