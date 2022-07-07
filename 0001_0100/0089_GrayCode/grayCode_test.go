package grayCode

import (
    "testing"
    "reflect"
)

func Test_grayCode(t *testing.T) {
    cases := []struct{
        name string
        input int
        expect int
    }{}
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T){
            ret := grayCode(c.input)
            if !reflect.DeepEqual(ret, c.expect) { 
                t.Fatalf("expect: %v, but got: %v, with input %v", c.expect, ret, c.input) 
            } 
        })
    }
}
