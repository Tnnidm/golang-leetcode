package groupAnagrams

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect [][]string
	}{
		{"case 1", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"case 2", []string{""}, [][]string{{""}}},
		{"case 3", []string{"a"}, [][]string{{"a"}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := groupAnagrams(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
