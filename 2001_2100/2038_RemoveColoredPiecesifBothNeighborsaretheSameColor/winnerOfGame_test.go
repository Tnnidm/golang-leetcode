package winnerOfGame

import (
	"reflect"
	"testing"
)

func Test_winnerOfGame(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect bool
	}{
		{"case 1", "AAABABB", true},
		{"case 2", "AA", false},
		{"case 3", "AAA", true},
		{"case 4", "BBB", false},
		{"case 5", "ABBBBBBBAAA", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := winnerOfGame(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
