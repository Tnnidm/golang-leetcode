package isPalindrome

import (
	"reflect"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect bool
	}{
		{"case 1", "A man, a plan, a canal: Panama", true},
		{"case 2", "race a car", false},
		{"case 3", "p", true},
		{"case 4", "pe", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isPalindrome(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v but got: %v with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
