package twoCitySchedCost

import (
	"reflect"
	"testing"
)

func Test_twoCitySchedCost(t *testing.T) {
	cases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{"case 1", [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
		{"case 2", [][]int{}, 0},
		{"case 3", [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, 1859},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := twoCitySchedCost(c.input)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expect: %v, but got: %v, with input: %v", c.expect, ret, c.input)
			}
		})
	}
}
