package solveSudoku

import (
	"reflect"
	"testing"
)

func Test_solveSudoku(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]byte
		expect [][]byte
	}{}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			solveSudoku(c.input)
			if !reflect.DeepEqual(c.input, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, c.input, c.input)
			}
		})
	}
}
