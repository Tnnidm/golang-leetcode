package implementStackUsingQueues

import (
	"reflect"
	"testing"
)

func implementStackUsingQueues() (int, int, bool) {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	result1 := myStack.Top()   // return 2
	result2 := myStack.Pop()   // return 2
	result3 := myStack.Empty() // return False
	return result1, result2, result3
}

func Test(t *testing.T) {
	t.Run(
		"test case 1",
		func(t *testing.T) {
			result1, result2, result3 := implementStackUsingQueues()
			if !(reflect.DeepEqual(result1, 2) && reflect.DeepEqual(result2, 2) && reflect.DeepEqual(result3, false)) {
				t.Fatalf("expected: 2, 2, false. but got: %v, %v, %v", result1, result2, result3)
			}
		},
	)
}
